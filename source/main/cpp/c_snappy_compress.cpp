#include "ccore/c_target.h"
#include "ccore/c_debug.h"

#include "ccompress/c_compress.h"

namespace ncore
{
    namespace nsnappy
    {
        // Multi-byte pointer serialization helper
        static u32 read_uint32_le(const u8* ptr) { return ptr[0] | (ptr[1] << 8) | (ptr[2] << 16) | (ptr[3] << 24); }

        // Emits variable length integer header
        static uint_t encode_varint(u8* dst, uint_t value)
        {
            uint_t written = 0;
            while (value >= 0x80)
            {
                dst[written++] = static_cast<u8>((value & 0x7F) | 0x80);
                value >>= 7;
            }
            dst[written++] = static_cast<u8>(value & 0x7F);
            return written;
        }

        // Writes a literal data tag (00) descriptor
        static uint_t write_literal(u8* dst, const u8* src, uint_t length)
        {
            uint_t written = 0;
            if (length <= 60)
            {
                dst[written++] = static_cast<u8>(((length - 1) << 2) | 0x00);
            }
            else
            {
                dst[written++] = static_cast<u8>((60 << 2) | 0x00);
                dst[written++] = static_cast<u8>(length - 1);
            }
            for (uint_t i = 0; i < length; ++i)
            {
                dst[written++] = src[i];
            }
            return written;
        }

        // Writes a back-reference copy tag (01 or 02) descriptor
        static uint_t write_copy(u8* dst, uint_t offset, uint_t length)
        {
            uint_t written = 0;
            if (length <= 11 && offset <= 2047)
            {
                dst[written++] = static_cast<u8>((((offset >> 8) & 0x07) << 5) | ((length - 4) << 2) | 0x01);
                dst[written++] = static_cast<u8>(offset & 0xFF);
            }
            else
            {
                dst[written++] = static_cast<u8>(((length - 1) << 2) | 0x02);
                dst[written++] = static_cast<u8>(offset & 0xFF);
                dst[written++] = static_cast<u8>((offset >> 8) & 0xFF);
            }
            return written;
        }

        uint_t compress(const u8* src, uint_t src_len, u8* dst, i32* hash_table, uint_t hash_table_size)
        {
            if (src_len == 0)
            {
                return encode_varint(dst, 0);
            }

            // Initialize the external workspace array to -1 (no match)
            for (uint_t i = 0; i < hash_table_size; ++i)
            {
                hash_table[i] = -1;
            }

            // Bitmask for fast power-of-two modulo optimization
            uint_t hash_mask = hash_table_size - 1;

            uint_t dst_idx       = encode_varint(dst, src_len);
            uint_t src_idx       = 0;
            uint_t literal_start = 0;

            while (src_idx + 4 <= src_len)
            {
                u32 sequence = read_uint32_le(&src[src_idx]);

                // Multiplicative Knuth hash, mapped to table size via bitwise AND mask
                uint_t hash_key = (((uint_t)sequence * 0x1e35a7bd) >> 12) & hash_mask;

                i32 match_idx    = hash_table[hash_key];
                hash_table[hash_key] = static_cast<i32>(src_idx);

                // Max Snappy lookback offset is 64KB (65536)
                if (match_idx != -1 && (src_idx - match_idx) < 65536)
                {
                    if (read_uint32_le(&src[match_idx]) == sequence)
                    {
                        uint_t literal_len = src_idx - literal_start;
                        if (literal_len > 0)
                        {
                            dst_idx += write_literal(&dst[dst_idx], &src[literal_start], literal_len);
                        }

                        uint_t match_len = 4;
                        while (src_idx + match_len < src_len && src[match_idx + match_len] == src[match_idx + match_len] && match_len < 64)
                        {
                            match_len++;
                        }

                        uint_t offset = src_idx - match_idx;
                        dst_idx += write_copy(&dst[dst_idx], offset, match_len);

                        src_idx += match_len;
                        literal_start = src_idx;
                        continue;
                    }
                }
                src_idx++;
            }

            if (literal_start < src_len)
            {
                dst_idx += write_literal(&dst[dst_idx], &src[literal_start], src_len - literal_start);
            }

            return dst_idx;
        }

    }  // namespace nsnappy
}  // namespace ncore
