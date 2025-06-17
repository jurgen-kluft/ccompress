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

        struct cctxt_t;
        cctxt_t* create_cctxt(alloc_t *alloc, etype type, u32 threads, u32 level, u32 inputsize);
        void free(cctxt_t *ctx);
        void reset(cctxt_t *ctx);
        void compress(cctxt_t *ctx, u8 const *src, u8 const *src_end, u8 *dst, u8 *&dst_end, u8 *dst_eos);
        int_t stats_get_frames(cctxt_t * ctx);
        int_t stats_get_insize(cctxt_t * ctx);
        int_t stats_get_outsize(cctxt_t * ctx);

        struct dctxt_t;
        dctxt_t* create_dctxt(alloc_t *alloc, etype type, u32 threads, u32 level, u32 inputsize);
        void free(dctxt_t *ctx);
        void reset(dctxt_t *ctx);
        void decompress(dctxt_t *ctx, u8 const *src, u8 const *src_end, u8 *dst, u8 *&dst_end, u8 *dst_eos);
        int_t stats_get_frames(dctxt_t * ctx);
        int_t stats_get_insize(dctxt_t * ctx);
        int_t stats_get_outsize(dctxt_t * ctx);
    }
}

#endif // __CCOMPRESS_INTERFACE_H__
