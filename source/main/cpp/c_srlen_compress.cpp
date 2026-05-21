#include "ccore/c_target.h"
#include "ccore/c_debug.h"

#include "ccompress/c_compress.h"

namespace ncore
{
    namespace nsrlen
    {
        struct bit_writer_t
        {
            u8* m_dst;
            u64 m_accumulator;
            u32 m_byte_idx;
            i32 m_bits_buffered;

            void init(u8* out_ptr)
            {
                m_dst           = out_ptr;
                m_byte_idx      = 0;
                m_accumulator   = 0;
                m_bits_buffered = 0;
            }

            void write_bits(u64 value, i32 count)
            {
                // Mask value to safe bit count bounds
                value &= (count >= 64) ? ~0ULL : ((1ULL << count) - 1);
                m_accumulator = (m_accumulator << count) | value;
                m_bits_buffered += count;

                // Flush full bytes out of the 64-bit cache accumulator
                while (m_bits_buffered >= 8)
                {
                    m_bits_buffered -= 8;
                    m_dst[m_byte_idx++] = static_cast<u8>(m_accumulator >> m_bits_buffered);
                }
            }

            void write_bits_repeat(u64 value, i32 count, i32 n)
            {
                for (i32 i = 0; i < n; ++i)
                {
                    write_bits(value, count);
                }
            }

            void flush()
            {
                if (m_bits_buffered > 0)
                {
                    m_accumulator <<= (8 - m_bits_buffered);
                    m_dst[m_byte_idx++] = static_cast<u8>(m_accumulator);
                    m_accumulator       = 0;
                    m_bits_buffered     = 0;
                }
            }
        };

        struct bit_reader_t
        {
            const u8* m_src;
            u32       m_byte_idx;
            u32       m_max_bytes;
            u64       m_accumulator;
            i32       m_bits_buffered;

            void init(const u8* in_ptr, uint_t len)
            {
                m_src           = in_ptr;
                m_byte_idx      = 0;
                m_max_bytes     = len;
                m_accumulator   = 0;
                m_bits_buffered = 0;
            }

            // @count: Number of bits to read (maximum = 8)
            u64 read_bits(i32 count)
            {
                // Keep the accumulator filled with at least 32 bits, or until we run out of input bytes
                while (m_bits_buffered < 32 && m_byte_idx < m_max_bytes)
                {
                    m_accumulator = (m_accumulator << 8) | m_src[m_byte_idx++];
                    m_bits_buffered += 8;
                }

                // Extract the requested bits from the accumulator
                i32 shift = m_bits_buffered - count;
                u64 value = (m_accumulator >> shift) & ((1ULL << count) - 1);
                m_bits_buffered -= count;
                return value;
            }

            // @count: Number of bits to peek (maximum = 8)
            u64 peek_bits(i32 count)
            {
                // Similar to read_bits but does not consume the bits from the accumulator
                while (m_bits_buffered < 32 && m_byte_idx < m_max_bytes)
                {
                    m_accumulator = (m_accumulator << 8) | m_src[m_byte_idx++];
                    m_bits_buffered += 8;
                }

                i32 shift = m_bits_buffered - count;
                return (m_accumulator >> shift) & ((1ULL << count) - 1);
            }

            // @count: Number of bits to consume (maximum = 8)
            void consume_bits(i32 count)
            {
                m_bits_buffered -= count;
                if (m_bits_buffered < 0)
                {
                    // If we consume more bits than buffered, we need to read more bytes to fill the accumulator
                    while (m_bits_buffered < 0 && m_byte_idx < m_max_bytes)
                    {
                        m_accumulator = (m_accumulator << 8) | m_src[m_byte_idx++];
                        m_bits_buffered += 8;
                    }
                }
            }

            bool has_more(i32 count) const
            {
                // Check if there are at least 'count' bits available to read
                return (m_byte_idx * 8 + m_bits_buffered) >= count;
            }
        };

        struct histogram_t
        {
            inline histogram_t(uint_t* run_counts_ptr, u32 max_rb, u32 symbol_bits)
                : m_run_counts(run_counts_ptr)
                , m_max_run_bits(max_rb)
                , m_symbol_bits(symbol_bits)
                , m_symbol_info_size(max_rb)
                , m_current_run_sym(0)
                , m_current_run_len(0)
            {
            }

            uint_t* m_run_counts;  // Pointer to workspace array segment for run count histograms
            u16     m_max_run_bits;
            u16     m_symbol_bits;
            u16     m_symbol_info_size;  // Precomputed value = max_run_bits, used for indexing into the workspace
            u64     m_current_run_sym;
            uint_t  m_current_run_len;
        };

        static inline void initialize(histogram_t& h)
        {
            uint_t total_unique_symbols = 1ULL << h.m_symbol_bits;
            for (uint_t s = 0; s < (total_unique_symbols * h.m_symbol_info_size); ++s)
            {
                h.m_run_counts[s] = 0;
            }
        }

        static inline void handle_chain(histogram_t& h)
        {
            // The repeating chain broke. Evaluate chunk counts for the completed run.
            for (i32 rb = 1; rb <= h.m_max_run_bits; ++rb)
            {
                uint_t max_allowed_chunk = 1ULL << rb;
                uint_t remaining_run     = h.m_current_run_len;
                while (remaining_run > 0)
                {
                    h.m_run_counts[h.m_current_run_sym * h.m_symbol_info_size + rb]++;
                    if (remaining_run >= max_allowed_chunk)
                    {
                        remaining_run -= max_allowed_chunk;
                    }
                    else
                    {
                        remaining_run = 0;
                    }
                }
            }
        }

        static inline void handle_symbol(histogram_t& h, u64 sym)
        {
            // Increment count for symbol 'sym' (rb = 0 tier)
            h.m_run_counts[sym * h.m_symbol_info_size]++;

            if (h.m_current_run_len == 0)
            {
                h.m_current_run_sym = sym;
                h.m_current_run_len = 1;
            }
            else if (sym == h.m_current_run_sym)
            {
                h.m_current_run_len++;
            }
            else
            {
                handle_chain(h);

                // Reset state machine to track the new symbol
                h.m_current_run_sym = sym;
                h.m_current_run_len = 1;
            }
        }

        static inline void flush(histogram_t& h)
        {
            // Flush the final active run remaining in the pipeline at End-of-Stream
            if (h.m_current_run_len > 0)
            {
                handle_chain(h);
            }
        }

        static inline void calculate_optimal_rb(const histogram_t& h, u8* optimal_rb)
        {
            uint_t total_unique_symbols = 1ULL << h.m_symbol_info_size;

            for (uint_t s = 0; s < total_unique_symbols; ++s)
            {
                // Baseline cost: Option 0 (RAW) means every instance costs exactly N bits
                uint_t best_size = h.m_run_counts[s * h.m_symbol_info_size] * h.m_symbol_bits;
                u8     best_rb   = 0;

                if (best_size == 0)
                {
                    optimal_rb[s] = 0;
                    continue;
                }

                // Check if any RLE run-bit tier beats the RAW baseline size
                for (i32 rb = 1; rb <= h.m_max_run_bits; ++rb)
                {
                    uint_t total_chunks = h.m_run_counts[s * h.m_symbol_info_size + rb];
                    if (total_chunks == 0)
                        continue;

                    // RLE Cost calculation formula:
                    // Total chunks * (Symbol Identity Size + Run Payload Bit Size)
                    uint_t rle_cost = total_chunks * (h.m_symbol_bits + rb);

                    if (rle_cost < best_size)
                    {
                        best_size = rle_cost;
                        best_rb   = rb;
                    }
                }
                optimal_rb[s] = best_rb;
            }
        }

        /**
         * SRLEN Multi-Phase Optimized Encoder
         *
         * @param src               Input packed bits array buffer.
         * @param total_symbols     Total count of N-bit elements contained inside src.
         * @param dst               Destination compressed bitstream payload block buffer.
         * @param symbol_bits       Size of a single symbol in bits (N). Supported: 1 to 8.
         * @param max_run_bits      The configurable max R-bits ceiling parameter. Supported: 1 to 8.
         * @param workspace         External scratch memory array for symbol_info_t, uint_t[(2^symbol_bits) * max_run_bits].
         * @param optimal_rb        Output array of size 2^symbol_bits to receive the selected optimal R-bit configuration for each symbol.
         * @return                  Total bytes written out to dst.
         */
        uint_t encode(const u8* src, uint_t total_symbols, u8* dst, i32 symbol_bits, i32 max_run_bits, uint_t* workspace, u8* optimal_rb)
        {
            if (total_symbols == 0)
                return 0;

            histogram_t h(workspace, max_run_bits, symbol_bits);
            initialize(h);

            // --- PHASE 1: SINGLE-PASS HISTOGRAM SWEEP ---
            bit_reader_t reader;
            reader.init(src, (total_symbols * symbol_bits + 7) / 8);

            uint_t symbols_read = 0;
            while (symbols_read < total_symbols)
            {
                const u64 sym = reader.read_bits(symbol_bits);
                handle_symbol(h, sym);
                symbols_read++;
            }
            flush(h);

            // --- PHASE 2: MATHEMATICAL STRATEGY SELECTION ---
            calculate_optimal_rb(h, optimal_rb);

            // --- PHASE 3: COMPRESSED BITSTREAM GENERATION ---
            bit_writer_t writer;
            writer.init(dst);

            // Emit the Compressed Stream Payloads
            reader.init(src, (total_symbols * symbol_bits + 7) / 8);
            uint_t sym_idx = 0;

            while (sym_idx < total_symbols)
            {
                const u64 sym = reader.read_bits(symbol_bits);
                const u8  rb  = optimal_rb[sym];

                if (rb == 0)
                {
                    writer.write_bits(sym, symbol_bits);
                    sym_idx++;
                }
                else
                {
                    uint_t max_run_len = 1ULL << rb;
                    uint_t run_len     = 1;

                    // Peek & Consume to pack matching items within the allowed run limit
                    while ((sym_idx + run_len < total_symbols) && (run_len < max_run_len))
                    {
                        if (reader.peek_bits(symbol_bits) == sym)
                        {
                            reader.consume_bits(symbol_bits);
                            run_len++;
                        }
                        else
                        {
                            break;
                        }
                    }

                    writer.write_bits(sym, symbol_bits);
                    writer.write_bits(run_len - 1, rb);

                    sym_idx += run_len;
                }
            }

            writer.flush();
            return static_cast<i32>(writer.m_byte_idx);
        }

        // SRLEN Parameterized Decoder
        i32 decode(const u8* src, uint_t src_len, u16 symbol_bits, u16 max_run_bits, u8* dst, uint_t max_dst_len, const u8* optimal_rb)
        {
            if (src_len < 10)
                return -1;

            bit_reader_t reader;
            reader.init(src, src_len);

            bit_writer_t writer;
            writer.init(dst);

            const uint_t total_symbols        = (src_len * 8) / symbol_bits;
            const uint_t total_unique_symbols = 1ULL << symbol_bits;
            uint_t       decoded_symbols      = 0;

            // Step 3: Decode and reconstruct arbitrary bits stream sequence payload blocks
            while (decoded_symbols < total_symbols && reader.has_more(symbol_bits))
            {
                const u16 sym = reader.read_bits(symbol_bits);
                const u8  rb  = optimal_rb[sym];

                if (rb == 0)
                {
                    writer.write_bits(sym, symbol_bits);
                    decoded_symbols++;
                }
                else
                {
                    const u16    run_minus_1 = reader.read_bits(rb);
                    const uint_t actual_run  = static_cast<uint_t>(run_minus_1) + 1;
                    writer.write_bits_repeat(sym, symbol_bits, actual_run);
                    decoded_symbols += actual_run;
                }
            }

            writer.flush();

            if (writer.m_byte_idx > max_dst_len)
                return -2;  // Security boundary guard

            return static_cast<i32>(writer.m_byte_idx);
        }
    }  // namespace nsrlen
}  // namespace ncore
