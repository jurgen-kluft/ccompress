#ifndef __CCOMPRESS_INTERFACE_H__
#define __CCOMPRESS_INTERFACE_H__
#include "ccore/c_target.h"
#ifdef USE_PRAGMA_ONCE
#pragma once
#endif

namespace ncore
{
    namespace nsnappy
    {
        // Compresses data using an externally-provided hash table workspace.
        //
        // @param src               Pointer to the input data.
        // @param src_len           Length of the input data.
        // @param dst               Pointer to the output destination buffer.
        // @param hash_table        Pointer to an external workspace array of i32.
        // @param hash_table_size   The number of elements in hash_table. MUST be a power of 2 (e.g., 512, 1024, 2048, 4096).
        // @return                  Total number of compressed bytes written to dst.
        //
        uint_t snappy_compress(const u8* src, uint_t src_len, u8* dst, i32* hash_table, uint_t hash_table_size);
        uint_t snappy_decompress(const u8* src, uint_t src_len, u8* dst, uint_t max_dst_len);
    }
}

#endif // __CCOMPRESS_INTERFACE_H__
