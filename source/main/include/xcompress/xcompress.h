// x_compressor.h - Core Compression API - 
#ifndef __XCOMPRESS_COMPRESSOR_INTERFACE_H__
#define __XCOMPRESS_COMPRESSOR_INTERFACE_H__
#include "xbase/x_target.h"
#ifdef USE_PRAGMA_ONCE
#pragma once
#endif


namespace xcore
{
	class xcompressor
	{
	public:
		virtual void		reset() = 0;
		virtual void		compress(u8 const* src, u8 const* src_end, u8 * dst, u8 *& dst_end, u8 * dst_eos) = 0;
		virtual void		decompress(u8 const* src, u8 const* src_end, u8 * dst, u8 *& dst_end, u8 * dst_eos) = 0;
	};
}

#endif	///< __XCOMPRESS_COMPRESSOR_INTERFACE_H__