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
        uint_t compress(const u8* src, uint_t src_len, u8* dst, i32* hash_table, uint_t hash_table_size);
        int_t decompress(const u8* src, uint_t src_len, u8* dst, uint_t max_dst_len);
    }

    namespace nlz4
    {
        // Compresses data using LZ4 format with variable compression level capabilities.
        //
        // @param src               Input data pointer.
        // @param src_len           Input data length.
        // @param dst               Output destination payload buffer.
        // @param workspace         External scratch memory array of i32.
        //                          Size MUST be:
        //                          - For levels 0-2: table_size elements.
        //                          - For levels 3-12 (HC): (table_size + src_len) elements.
        // @param table_size        Hash table size. Must be a power of two (e.g., 4096, 16384).
        // @param compression_level 0-2 for Fast LZ4, 3-12 for High Compression (HC mode).
        // @return                  Total compressed bytes written to dst.
        //
        uint_t compress(const u8* src, uint_t src_len, u8* dst, i32* workspace, uint_t table_size, i32 compression_level);
        int_t decompress(const u8* src, uint_t src_len, u8* dst, uint_t max_dst_len);
    }
}

#endif // __CCOMPRESS_INTERFACE_H__
