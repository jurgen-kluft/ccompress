#ifndef SNAPPYMT_H
#define SNAPPYMT_H
#include "ccore/c_target.h"
#ifdef USE_PRAGMA_ONCE
#pragma once
#endif

namespace ncore
{
#define SNAPPY_OK 0
#define SNAPPYMT_THREAD_MAX 128
#define SNAPPYMT_MAGICNUMBER 0x5053          // SP
#define SNAPPYMT_MAGIC_SKIPPABLE 0x184D2A50U // MT magic number

    typedef enum
    {
        SNAPPYMT_error_no_error,
        SNAPPYMT_error_memory_allocation,
        SNAPPYMT_error_read_fail,
        SNAPPYMT_error_write_fail,
        SNAPPYMT_error_data_error,
        SNAPPYMT_error_frame_compress,
        SNAPPYMT_error_frame_decompress,
        SNAPPYMT_error_compressionParameter_unsupported,
        SNAPPYMT_error_compression_library,
        SNAPPYMT_error_canceled,
        SNAPPYMT_error_maxCode
    } SNAPPYMT_ErrorCode;

    typedef struct
    {
        void *buf;
        int_t size;
        int_t allocated;
    } buffer_t;

    typedef int(fnRead)(void *args, buffer_t *in);
    typedef int(fnWrite)(void *args, buffer_t *out);

    typedef struct
    {
        fnRead *fn_read;
        void *arg_read;
        fnWrite *fn_write;
        void *arg_write;
    } SNAPPYMT_RdWr_t;

}
#endif
