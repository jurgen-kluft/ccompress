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

        struct compress_t;
        compress_t* create_cctxt(alloc_t *alloc, etype type, u32 threads, u32 level, u32 inputsize);
        void free(compress_t *ctx);
        void reset(compress_t *ctx);
        void compress(compress_t *ctx, u8 const *src, u8 const *src_end, u8 *dst, u8 *&dst_end, u8 *dst_eos);
        int_t stats_get_frames(compress_t * ctx);
        int_t stats_get_insize(compress_t * ctx);
        int_t stats_get_outsize(compress_t * ctx);

        struct decompress_t;
        decompress_t* create_dctxt(alloc_t *alloc, etype type, u32 threads, u32 level, u32 inputsize);
        void free(decompress_t *ctx);
        void reset(decompress_t *ctx);
        void decompress(decompress_t *ctx, u8 const *src, u8 const *src_end, u8 *dst, u8 *&dst_end, u8 *dst_eos);
        int_t stats_get_frames(decompress_t * ctx);
        int_t stats_get_insize(decompress_t * ctx);
        int_t stats_get_outsize(decompress_t * ctx);
    }
}

#endif // __CCOMPRESS_INTERFACE_H__
