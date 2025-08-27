#ifndef __CCOMPRESS_INTERFACE_H__
#define __CCOMPRESS_INTERFACE_H__
#include "ccore/c_target.h"
#ifdef USE_PRAGMA_ONCE
#pragma once
#endif

namespace ncore
{
    class alloc_t;

    namespace ncompress
    {
        enum etype
        {
            type_none = 0,
            type_lz4,
            type_snappy,
            type_max
        };

        struct compressor_t;
        compressor_t* create_compressor(alloc_t *alloc, etype type, u32 threads, u32 level, u32 inputsize);
        void free(compressor_t *ctx);
        void reset(compressor_t *ctx);
        void compress(compressor_t *ctx, u8 const *src, u8 const *src_end, u8 *dst, u8 *&dst_end, u8 *dst_eos);
        int_t stats_get_frames(compressor_t * ctx);
        int_t stats_get_insize(compressor_t * ctx);
        int_t stats_get_outsize(compressor_t * ctx);

        struct decompressor_t;
        decompressor_t* create_decompressor(alloc_t *alloc, etype type, u32 threads, u32 level, u32 inputsize);
        void free(decompressor_t *ctx);
        void reset(decompressor_t *ctx);
        void decompress(decompressor_t *ctx, u8 const *src, u8 const *src_end, u8 *dst, u8 *&dst_end, u8 *dst_eos);
        int_t stats_get_frames(decompressor_t * ctx);
        int_t stats_get_insize(decompressor_t * ctx);
        int_t stats_get_outsize(decompressor_t * ctx);
    }
}

#endif // __CCOMPRESS_INTERFACE_H__
