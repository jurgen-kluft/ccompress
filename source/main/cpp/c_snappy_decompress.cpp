#include "ccore/c_target.h"
#include "ccore/c_debug.h"

#include "ccompress/c_compress.h"

namespace ncore
{
    namespace nsnappy
    {
        // Parses variable length integer header
        static uint_t decode_varint(const u8* src, uint_t* src_idx)
        {
            uint_t value = 0;
            uint_t shift = 0;
            uint_t idx   = *src_idx;

            while (true)
            {
                u8 b = src[idx++];
                value |= static_cast<uint_t>(b & 0x7F) << shift;
                if (!(b & 0x80))
                    break;
                shift += 7;
            }
            *src_idx = idx;
            return value;
        }

        // Global Core Decoder Function
        int_t decompress(const u8* src, uint_t src_len, u8* dst, uint_t max_dst_len)
        {
            if (src_len == 0)
                return 0;

            uint_t src_idx = 0;
            uint_t dst_idx = 0;

            uint_t uncompressed_len = decode_varint(src, &src_idx);
            if (uncompressed_len > max_dst_len)
            {
                return -1;  // Avoid output overflow corruption
            }

            while (src_idx < src_len)
            {
                u8 tag_byte = src[src_idx++];
                u8 tag_type = tag_byte & 0x03;

                // Tag 00: Literal Parsing
                if (tag_type == 0x00)
                {
                    uint_t length = tag_byte >> 2;
                    if (length < 60)
                    {
                        length += 1;
                    }
                    else
                    {
                        uint_t num_bytes = length - 59;
                        length           = 0;
                        for (uint_t i = 0; i < num_bytes; ++i)
                        {
                            length |= static_cast<uint_t>(src[src_idx++]) << (8 * i);
                        }
                        length += 1;
                    }

                    for (uint_t i = 0; i < length; ++i)
                    {
                        dst[dst_idx++] = src[src_idx++];
                    }
                }
                // Tag 01: Copy 1-byte explicit offset layout
                else if (tag_type == 0x01)
                {
                    uint_t length = ((tag_byte >> 2) & 0x07) + 4;
                    uint_t offset = ((tag_byte >> 5) & 0x07) << 8;
                    offset |= src[src_idx++];

                    for (uint_t i = 0; i < length; ++i)
                    {
                        uint_t back_idx = dst_idx - offset;
                        dst[dst_idx++]  = dst[back_idx];
                    }
                }
                // Tag 02: Copy 2-byte explicit offset layout
                else if (tag_type == 0x02)
                {
                    uint_t length = (tag_byte >> 2) + 1;
                    uint_t offset = src[src_idx] | (src[src_idx + 1] << 8);
                    src_idx += 2;

                    for (uint_t i = 0; i < length; ++i)
                    {
                        uint_t back_idx = dst_idx - offset;
                        dst[dst_idx++]  = dst[back_idx];
                    }
                }
            }

            return (dst_idx == uncompressed_len) ? static_cast<i32>(dst_idx) : -2;
        }

    }  // namespace nsnappy
}  // namespace ncore
