#include "ccore/c_target.h"
#include "ccore/c_debug.h"

#include "ccompress/c_compress.h"

namespace ncore
{
    namespace nlz4
    {
        /**
         * Decompresses an LZ4 compressed block using strict C-style memory bounds.
         *
         * @param src          Pointer to the compressed input stream.
         * @param src_len      Total length of the compressed input data.
         * @param dst          Pointer to the pre-allocated output buffer.
         * @param max_dst_len  The maximum safety capacity of the dst buffer.
         * @return             Total decompressed bytes written on success,
         *                     or a negative integer code on structural corruption.
         */
        int_t decompress(const u8* src, uint_t src_len, u8* dst, uint_t max_dst_len)
        {
            if (src_len == 0)
                return 0;

            const u8* const src_end = src + src_len;
            u8* const       dst_end = dst + max_dst_len;

            const u8* src_ptr = src;
            u8*       dst_ptr = dst;

            while (src_ptr < src_end)
            {
                // 1. Read Token Byte
                u8 token   = *src_ptr++;
                uint_t  lit_len = token >> 4;
                uint_t  mat_len = token & 0x0F;

                // 2. Parse Variable-Length Literal Extensions
                if (lit_len == 15)
                {
                    u8 ext_byte;
                    do
                    {
                        if (src_ptr >= src_end)
                            return -1;  // End of Stream Error
                        ext_byte = *src_ptr++;
                        lit_len += ext_byte;
                    } while (ext_byte == 255);
                }

                // 3. Copy Literal Payload to Destination
                if (lit_len > 0)
                {
                    if (dst_ptr + lit_len > dst_end || src_ptr + lit_len > src_end)
                    {
                        return -2;  // Destination or Source Buffer Overflow
                    }
                    // Byte-by-byte direct memory copy loop
                    for (uint_t i = 0; i < lit_len; ++i)
                    {
                        *dst_ptr++ = *src_ptr++;
                    }
                }

                // If we hit the exact end of the source buffer, parsing stops successfully
                if (src_ptr >= src_end)
                {
                    break;
                }

                // 4. Read Match Offset (2 Bytes, Little-Endian)
                if (src_ptr + 2 > src_end)
                    return -3;  // Truncated offset field
                uint_t offset = src_ptr[0] | (static_cast<uint_t>(src_ptr[1]) << 8);
                src_ptr += 2;

                if (offset == 0 || dst_ptr - offset < dst)
                {
                    return -4;  // Invalid Back-Reference Lookback Window
                }

                // 5. Parse Variable-Length Match Extensions
                if (mat_len == 15)
                {
                    u8 ext_byte;
                    do
                    {
                        if (src_ptr >= src_end)
                            return -1;
                        ext_byte = *src_ptr++;
                        mat_len += ext_byte;
                    } while (ext_byte == 255);
                }
                mat_len += 4;  // Add the minimum implied match layout constraint

                if (dst_ptr + mat_len > dst_end)
                {
                    return -2;  // Destination buffer overflow protection
                }

                // 6. Copy Overlapping Match Dictionary Windows
                u8* match_ptr = dst_ptr - offset;
                for (uint_t i = 0; i < mat_len; ++i)
                {
                    *dst_ptr++ = *match_ptr++;
                }
            }

            // Return total decompressed output payload size metric
            return (int_t)(dst_ptr - dst);
        }

    }  // namespace nlz4
}  // namespace ncore
