#include "cbase/c_base.h"
#include "cbase/c_allocator.h"
#include "cbase/c_console.h"

#include "ncore/x_core.h"

#include "cunittest\cunittest.h"

UNITTEST_SUITE_LIST(xCoreUnitTest);

UNITTEST_SUITE_DECLARE(xCoreUnitTest, xlist);
UNITTEST_SUITE_DECLARE(xCoreUnitTest, xhashmap);
UNITTEST_SUITE_DECLARE(xCoreUnitTest, xqueue);
UNITTEST_SUITE_DECLARE(xCoreUnitTest, xmap);
UNITTEST_SUITE_DECLARE(xCoreUnitTest, xpqueue);
UNITTEST_SUITE_DECLARE(xCoreUnitTest, xstack);
UNITTEST_SUITE_DECLARE(xCoreUnitTest, xilist);
UNITTEST_SUITE_DECLARE(xCoreUnitTest, xset);
UNITTEST_SUITE_DECLARE(xCoreUnitTest, xtree);
UNITTEST_SUITE_DECLARE(xCoreUnitTest, xvector);

namespace ncore
{
	class UnitTestAllocator : public UnitTest::Allocator
	{
		x_iallocator*	mAllocator;
	public:
						UnitTestAllocator(x_iallocator* allocator)			{ mAllocator = allocator; }
		virtual void*	Allocate(size_t size)								{ return mAllocator->allocate(size, 4); }
		virtual void	Deallocate(void* ptr)								{ mAllocator->deallocate(ptr); }
	};

	class TestAllocator : public x_iallocator
	{
		x_iallocator*		mAllocator;
	public:
							TestAllocator(x_iallocator* allocator) : mAllocator(allocator) { }

		virtual const char*	name() const										{ return "ncore unittest test heap allocator"; }

		virtual void*		allocate(xsize_t size, u32 alignment)
		{
			UnitTest::IncNumAllocations();
			return mAllocator->allocate(size, alignment);
		}

		virtual void*		reallocate(void* mem, xsize_t size, u32 alignment)
		{
			if (mem==NULL)
				return allocate(size, alignment);
			else 
				return mAllocator->reallocate(mem, size, alignment);
		}

		virtual void		deallocate(void* mem)
		{
			UnitTest::DecNumAllocations();
			mAllocator->deallocate(mem);
		}

		virtual void		release()
		{
			mAllocator->release();
			mAllocator = NULL;
		}
	};
}

ncore::x_iallocator* gTestAllocator = NULL;

bool gRunUnitTest(UnitTest::TestReporter& reporter)
{
	cbase::init();

	ncore::UnitTestAllocator unittestAllocator( ncore::x_iallocator::get_default() );
	UnitTest::SetAllocator(&unittestAllocator);

	ncore::xconsole::add_default_console();

	ncore::TestAllocator testAllocator(ncore::x_iallocator::get_default());
	gTestAllocator = &testAllocator;

	ncore::x_Init(gTestAllocator);
	int r = UNITTEST_SUITE_RUN(reporter, xCoreUnitTest);
	ncore::x_Exit();

	gTestAllocator->release();

	UnitTest::SetAllocator(NULL);

	cbase::exit();
	return r==0;
}

