#include "ccore/c_target.h"

#include "ccompress/c_compress.h"

#include "cunittest/cunittest.h"

using namespace ncore;

UNITTEST_SUITE_BEGIN(lz4)
{
    UNITTEST_FIXTURE(base)
    {
        UNITTEST_FIXTURE_SETUP() {}
        UNITTEST_FIXTURE_TEARDOWN() {}

        UNITTEST_TEST(test_compress)
        {
            // Highly repetitive mock text stream designed to benchmark dictionary tree depth matching
            const u8 payload[] = "BANANA_ANANAS_BANANA_ANANAS_PARAM_PANAMA_BANANA_ANANAS_BANANA_ANANAS_END";
            uint_t   len       = sizeof(payload);

            u8           output_buffer[512];
            const uint_t HASH_SIZE = 4096;

            // 1. Run Level 1 (Fast Single-Hash Mode)
            i32    fast_workspace[HASH_SIZE];
            uint_t fast_size = nlz4::compress(payload, len, output_buffer, fast_workspace, HASH_SIZE, 1);

            // 2. Run Level 12 (Ultra High-Compression HC Chain-Backtracking Mode)
            // Dynamic buffer footprint expands: Hash slots + Match Index positions
            i32    hc_workspace[HASH_SIZE + sizeof(payload)];
            uint_t hc_size = nlz4::compress(payload, len, output_buffer, hc_workspace, HASH_SIZE, 12);

            // std::cout << "Original Size:       " << len << " bytes\n";
            // std::cout << "LZ4 Fast (Level 1):  " << fast_size << " bytes\n";
            // std::cout << "LZ4 HC   (Level 12): " << hc_size << " bytes\n";
        }

        UNITTEST_TEST(test_roundtrip)
        {
            // Repeated pattern string payload
            const u8 original_data[] = "MISSISSIPPI_RIVER_DENSE_DICTIONARY_MATCH_MISSISSIPPI_RIVER_DENSE_DICTIONARY_MATCH";
            uint_t   orig_len        = sizeof(original_data);

            u8 compressed_buffer[512];
            u8 decompressed_buffer[512];

            const uint_t TABLE_SIZE = 1024;
            i32          fast_workspace[TABLE_SIZE];

            // 1. Compress using Fast Engine (Level 1)
            uint_t comp_len = nlz4::compress(original_data, orig_len, compressed_buffer, fast_workspace, TABLE_SIZE, 1);

            // 2. Decompress using our bounds-checked decoder
            int_t decomp_status = nlz4::decompress(compressed_buffer, comp_len, decompressed_buffer, 512);

            // 3. Print verification diagnostic reports
            // std::cout << "Original Size:       " << orig_len << " bytes\n";
            // std::cout << "Compressed Size:     " << comp_len << " bytes\n";
            // std::cout << "Decompress Status:   " << decomp_status << " bytes output\n";

            if (decomp_status > 0 && static_cast<uint_t>(decomp_status) == orig_len)
            {
                bool match = true;
                for (uint_t i = 0; i < orig_len; ++i)
                {
                    if (original_data[i] != decompressed_buffer[i])
                        match = false;
                }
                CHECK_TRUE_T(match, "Integrity Validation FAILED");
            }
            else
            {
                CHECK_GE_T(decomp_status, 0, "Decompression Error");
            }
        }
    }
}
UNITTEST_SUITE_END
