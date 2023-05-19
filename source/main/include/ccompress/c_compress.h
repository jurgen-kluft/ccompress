#ifndef __CCOMPRESS_COMPRESSOR_INTERFACE_H__
#define __CCOMPRESS_COMPRESSOR_INTERFACE_H__
#include "ccore/c_target.h"
#ifdef USE_PRAGMA_ONCE
#pragma once
#endif

namespace ncore
{
	class compressor_t
	{
	public:
		virtual void		reset() = 0;
		virtual void		compress(u8 const* src, u8 const* src_end, u8 * dst, u8 *& dst_end, u8 * dst_eos) = 0;
		virtual void		decompress(u8 const* src, u8 const* src_end, u8 * dst, u8 *& dst_end, u8 * dst_eos) = 0;
	};
}

#endif	///< __CCOMPRESS_COMPRESSOR_INTERFACE_H__