#include "ccore/c_target.h"
#include "ccore/c_debug.h"

#include "ccompress/c_compress.h"

namespace ncore
{
    namespace nlz4
    {
        // Read 4 bytes natively
        static u32 read32(const u8* ptr) { return ptr[0] | (ptr[1] << 8) | (ptr[2] << 16) | (ptr[3] << 24); }

        // Emits variable-length additive extension bytes for LZ4 large counts
        static uint_t write_length_extension(u8* dst, uint_t length)
        {
            uint_t idx = 0;
            while (length >= 255)
            {
                dst[idx++] = 255;
                length -= 255;
            }
            dst[idx++] = static_cast<u8>(length);
            return idx;
        }

        /**
         * Compresses data using LZ4 format with variable compression level capabilities.
         *
         * @param src               Input data pointer.
         * @param src_len           Input data length.
         * @param dst               Output destination payload buffer.
         * @param workspace         External scratch memory array of i32.
         *                          Size MUST be:
         *                          - For levels 0-2: table_size elements.
         *                          - For levels 3-12 (HC): (table_size + src_len) elements.
         * @param table_size        Hash table size. Must be a power of two (e.g., 4096, 16384).
         * @param compression_level 0-2 for Fast LZ4, 3-12 for High Compression (HC mode).
         * @return                  Total compressed bytes written to dst.
         */
        uint_t compress(const u8* src, uint_t src_len, u8* dst, i32* workspace, uint_t table_size, i32 compression_level)
        {
            if (src_len == 0)
                return 0;

            uint_t hash_mask = table_size - 1;
            uint_t dst_idx   = 0;
            uint_t src_idx   = 0;
            uint_t anchor    = 0;

            // Split workspace pointers based on strategy level choice
            i32* hash_table  = workspace;
            i32* chain_table = (compression_level >= 3) ? (workspace + table_size) : nullptr;

            // Clear hash table workspace
            for (uint_t i = 0; i < table_size; ++i)
            {
                hash_table[i] = -1;
            }

            // Configure backtracking depth limit based on HC tier selection
            uint_t max_attempts = 1;
            if (compression_level >= 3)
            {
                if (compression_level < 6)
                    max_attempts = 4;
                else if (compression_level < 9)
                    max_attempts = 16;
                else if (compression_level < 12)
                    max_attempts = 64;
                else
                    max_attempts = 256;  // Level 12 ultra deep match
            }

            // Leave trailing safety margin for 4-byte matches and tail encoding
            while (src_idx + 12 < src_len)
            {
                u32    h_val    = read32(&src[src_idx]);
                uint_t hash_key = ((h_val * 2654435761U) >> 16) & hash_mask;

                i32 match_idx = hash_table[hash_key];

                // Update historical records
                if (chain_table)
                {
                    // HC Mode: Record the previous owner of this hash bucket into the chain
                    chain_table[src_idx] = match_idx;
                }
                hash_table[hash_key] = static_cast<i32>(src_idx);

                uint_t best_match_len = 0;
                uint_t best_offset    = 0;
                i32    curr_candidate = match_idx;
                uint_t attempts       = 0;

                // --- MATCH SEARCH ENGINE ---
                while (curr_candidate != -1 && (src_idx - curr_candidate) < 65536 && attempts < max_attempts)
                {
                    if (read32(&src[curr_candidate]) == h_val)
                    {
                        uint_t current_len = 4;
                        while (src_idx + current_len < src_len - 5 && src[curr_candidate + current_len] == src[src_idx + current_len])
                        {
                            current_len++;
                        }

                        if (current_len > best_match_len)
                        {
                            best_match_len = current_len;
                            best_offset    = src_idx - curr_candidate;
                        }
                    }

                    if (!chain_table)
                        break;                                     // Fast Mode (Level 0-2) only evaluates the absolute newest slot
                    curr_candidate = chain_table[curr_candidate];  // HC Mode moves down the history chain
                    attempts++;
                }

                // --- EMIT SEQUENCE ---
                if (best_match_len >= 4)
                {
                    uint_t lit_len         = src_idx - anchor;
                    u8     token_lit       = (lit_len >= 15) ? 15 : static_cast<u8>(lit_len);
                    uint_t mat_len_minus_4 = best_match_len - 4;
                    u8     token_mat       = (mat_len_minus_4 >= 15) ? 15 : static_cast<u8>(mat_len_minus_4);

                    // 1. Token Byte
                    dst[dst_idx++] = static_cast<u8>((token_lit << 4) | token_mat);

                    // 2. Extended Literal Lengths
                    if (lit_len >= 15)
                    {
                        dst_idx += write_length_extension(&dst[dst_idx], lit_len - 15);
                    }

                    // 3. Literal Payload bytes Copy
                    for (uint_t i = 0; i < lit_len; ++i)
                    {
                        dst[dst_idx++] = src[anchor + i];
                    }

                    // 4. Back Reference Offset (2 Bytes, Little-Endian)
                    dst[dst_idx++] = static_cast<u8>(best_offset & 0xFF);
                    dst[dst_idx++] = static_cast<u8>((best_offset >> 8) & 0xFF);

                    // 5. Extended Match Lengths
                    if (mat_len_minus_4 >= 15)
                    {
                        dst_idx += write_length_extension(&dst[dst_idx], mat_len_minus_4 - 15);
                    }

                    // Move past the encoded sequence window block
                    if (chain_table)
                    {
                        // Populate chain values for internal elements skipped during macro jump
                        for (uint_t j = 1; j < best_match_len; ++j)
                        {
                            uint_t skipped_idx = src_idx + j;
                            if (skipped_idx + 4 <= src_len)
                            {
                                u32    s_val             = read32(&src[skipped_idx]);
                                uint_t s_key             = ((s_val * 2654435761U) >> 16) & hash_mask;
                                chain_table[skipped_idx] = hash_table[s_key];
                                hash_table[s_key]        = static_cast<i32>(skipped_idx);
                            }
                        }
                    }

                    src_idx += best_match_len;
                    anchor = src_idx;
                }
                else
                {
                    src_idx++;
                }
            }

            // --- EMIT FINAL FRAGMENT (LAST TRAILING LITERALS) ---
            uint_t final_lit_len = src_len - anchor;
            u8     final_token   = (final_lit_len >= 15) ? 15 : static_cast<u8>(final_lit_len);
            dst[dst_idx++]       = static_cast<u8>(final_token << 4);  // Match bits are explicitly zeroed out at EOF

            if (final_lit_len >= 15)
            {
                dst_idx += write_length_extension(&dst[dst_idx], final_lit_len - 15);
            }
            for (uint_t i = 0; i < final_lit_len; ++i)
            {
                dst[dst_idx++] = src[anchor + i];
            }

            return dst_idx;
        }

    }  // namespace nlz4
}  // namespace ncore
