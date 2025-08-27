// --------------------------------------------------------------------
// ---------------------- GENERATED -----------------------------------
// --------------------------------------------------------------------
package main

import (
    clay "github.com/jurgen-kluft/ccode/clay"
)

func main() {
    clay.ClayAppCreateProjectsFunc = CreateProjects
    clay.ClayAppMainDesktop()
}

// Setup Project Identifiers
const (
    library_ccore_debug_dev_id int = 0
    library_ccore_release_dev_id int = 1
    library_cbase_debug_dev_id int = 2
    library_cbase_release_dev_id int = 3
    library_ccompress_debug_dev_id int = 4
    library_ccompress_release_dev_id int = 5
    unittest_library_cunittest_debug_dev_test_id int = 6
    unittest_library_cunittest_release_dev_test_id int = 7
    unittest_library_ccore_debug_dev_test_id int = 8
    unittest_library_ccore_release_dev_test_id int = 9
    unittest_library_cbase_debug_dev_test_id int = 10
    unittest_library_cbase_release_dev_test_id int = 11
    unittest_library_ccompress_debug_dev_test_id int = 12
    unittest_library_ccompress_release_dev_test_id int = 13
    unittest_ccompress_debug_dev_test_id int = 14
    unittest_ccompress_release_dev_test_id int = 15
)

func CreateProjects() []*clay.Project {
    projects := make([]*clay.Project, 16)
    {
        configName := "debug-dev"
        projectName := "library_ccore"
        projectConfig := clay.NewConfig("darwin", "arm64", configName)
        project := clay.NewLibraryProject(projectName, projectConfig)

        // Project Include directories
        project.IncludeDirs = clay.NewIncludeMap(1)
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/include")

        // Project Define macros
        configDefines := projectConfig.GetCppDefines()
        project.Defines = clay.NewDefineMap(5 + len(configDefines))
        project.Defines.Add("TARGET_MAC")
        project.Defines.Add("UNICODE")
        project.Defines.Add("_UNICODE")
        project.Defines.Add("TARGET_DEBUG")
        project.Defines.Add("_DEBUG")
        project.Defines.AddMany(configDefines...)

        // Project Source files
        project.SourceFiles = make([]clay.SourceFile, 0, 10)
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_allocator.cpp", "c_allocator.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_binary_search.cpp", "c_binary_search.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_binmap1.cpp", "c_binmap1.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_debug.cpp", "c_debug.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_error.cpp", "c_error.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_memcpy_neon.cpp", "c_memcpy_neon.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_memcpy_sse.cpp", "c_memcpy_sse.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_memory.cpp", "c_memory.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_qsort.cpp", "c_qsort.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_vmem.cpp", "c_vmem.cpp")

        projects[library_ccore_debug_dev_id] = project
    }
    {
        configName := "release-dev"
        projectName := "library_ccore"
        projectConfig := clay.NewConfig("darwin", "arm64", configName)
        project := clay.NewLibraryProject(projectName, projectConfig)

        // Project Include directories
        project.IncludeDirs = clay.NewIncludeMap(1)
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/include")

        // Project Define macros
        configDefines := projectConfig.GetCppDefines()
        project.Defines = clay.NewDefineMap(5 + len(configDefines))
        project.Defines.Add("TARGET_MAC")
        project.Defines.Add("UNICODE")
        project.Defines.Add("_UNICODE")
        project.Defines.Add("TARGET_RELEASE")
        project.Defines.Add("NDEBUG")
        project.Defines.AddMany(configDefines...)

        // Project Source files
        project.SourceFiles = make([]clay.SourceFile, 0, 10)
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_allocator.cpp", "c_allocator.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_binary_search.cpp", "c_binary_search.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_binmap1.cpp", "c_binmap1.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_debug.cpp", "c_debug.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_error.cpp", "c_error.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_memcpy_neon.cpp", "c_memcpy_neon.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_memcpy_sse.cpp", "c_memcpy_sse.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_memory.cpp", "c_memory.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_qsort.cpp", "c_qsort.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_vmem.cpp", "c_vmem.cpp")

        projects[library_ccore_release_dev_id] = project
    }
    {
        configName := "debug-dev"
        projectName := "library_cbase"
        projectConfig := clay.NewConfig("darwin", "arm64", configName)
        project := clay.NewLibraryProject(projectName, projectConfig)

        // Project Include directories
        project.IncludeDirs = clay.NewIncludeMap(2)
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/include")

        // Project Define macros
        configDefines := projectConfig.GetCppDefines()
        project.Defines = clay.NewDefineMap(5 + len(configDefines))
        project.Defines.Add("TARGET_MAC")
        project.Defines.Add("UNICODE")
        project.Defines.Add("_UNICODE")
        project.Defines.Add("TARGET_DEBUG")
        project.Defines.Add("_DEBUG")
        project.Defines.AddMany(configDefines...)

        // Project Source files
        project.SourceFiles = make([]clay.SourceFile, 0, 32)
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_allocator.cpp", "c_allocator.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_allocator_system.cpp", "c_allocator_system.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_base.cpp", "c_base.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_binary_search.cpp", "c_binary_search.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_binmap.cpp", "c_binmap.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_buffer.cpp", "c_buffer.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_console.cpp", "c_console.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_console_mac.cpp", "c_console_mac.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_context.cpp", "c_context.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_dconv.cpp", "c_dconv.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_debug.cpp", "c_debug.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_duomap.cpp", "c_duomap.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_guid.cpp", "c_guid.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_hash.cpp", "c_hash.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_log.cpp", "c_log.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_log_to_console.cpp", "c_log_to_console.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_map.cpp", "c_map.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_memory.cpp", "c_memory.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_pdqsort.cpp", "c_pdqsort.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_printf.cpp", "c_printf.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_random.cpp", "c_random.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_runes.cpp", "c_runes.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_slice.cpp", "c_slice.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_sscanf.cpp", "c_sscanf.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_strfmt.cpp", "c_strfmt.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_tblfmt.cpp", "c_tblfmt.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_tree.cpp", "c_tree.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_tree32.cpp", "c_tree32.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_tree_obsolete.cpp", "c_tree_obsolete.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_va_list.cpp", "c_va_list.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_vector.cpp", "c_vector.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_wyhash.cpp", "c_wyhash.cpp")

        projects[library_cbase_debug_dev_id] = project
    }
    {
        configName := "release-dev"
        projectName := "library_cbase"
        projectConfig := clay.NewConfig("darwin", "arm64", configName)
        project := clay.NewLibraryProject(projectName, projectConfig)

        // Project Include directories
        project.IncludeDirs = clay.NewIncludeMap(2)
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/include")

        // Project Define macros
        configDefines := projectConfig.GetCppDefines()
        project.Defines = clay.NewDefineMap(5 + len(configDefines))
        project.Defines.Add("TARGET_MAC")
        project.Defines.Add("UNICODE")
        project.Defines.Add("_UNICODE")
        project.Defines.Add("TARGET_RELEASE")
        project.Defines.Add("NDEBUG")
        project.Defines.AddMany(configDefines...)

        // Project Source files
        project.SourceFiles = make([]clay.SourceFile, 0, 32)
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_allocator.cpp", "c_allocator.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_allocator_system.cpp", "c_allocator_system.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_base.cpp", "c_base.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_binary_search.cpp", "c_binary_search.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_binmap.cpp", "c_binmap.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_buffer.cpp", "c_buffer.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_console.cpp", "c_console.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_console_mac.cpp", "c_console_mac.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_context.cpp", "c_context.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_dconv.cpp", "c_dconv.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_debug.cpp", "c_debug.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_duomap.cpp", "c_duomap.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_guid.cpp", "c_guid.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_hash.cpp", "c_hash.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_log.cpp", "c_log.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_log_to_console.cpp", "c_log_to_console.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_map.cpp", "c_map.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_memory.cpp", "c_memory.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_pdqsort.cpp", "c_pdqsort.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_printf.cpp", "c_printf.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_random.cpp", "c_random.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_runes.cpp", "c_runes.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_slice.cpp", "c_slice.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_sscanf.cpp", "c_sscanf.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_strfmt.cpp", "c_strfmt.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_tblfmt.cpp", "c_tblfmt.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_tree.cpp", "c_tree.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_tree32.cpp", "c_tree32.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_tree_obsolete.cpp", "c_tree_obsolete.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_va_list.cpp", "c_va_list.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_vector.cpp", "c_vector.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_wyhash.cpp", "c_wyhash.cpp")

        projects[library_cbase_release_dev_id] = project
    }
    {
        configName := "debug-dev"
        projectName := "library_ccompress"
        projectConfig := clay.NewConfig("darwin", "arm64", configName)
        project := clay.NewLibraryProject(projectName, projectConfig)

        // Project Include directories
        project.IncludeDirs = clay.NewIncludeMap(3)
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccompress/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/include")

        // Project Define macros
        configDefines := projectConfig.GetCppDefines()
        project.Defines = clay.NewDefineMap(5 + len(configDefines))
        project.Defines.Add("TARGET_MAC")
        project.Defines.Add("UNICODE")
        project.Defines.Add("_UNICODE")
        project.Defines.Add("TARGET_DEBUG")
        project.Defines.Add("_DEBUG")
        project.Defines.AddMany(configDefines...)

        // Project Source files
        project.SourceFiles = make([]clay.SourceFile, 0, 1)
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccompress/source/main/cpp/c_snappy.cpp", "c_snappy.cpp")

        projects[library_ccompress_debug_dev_id] = project
    }
    {
        configName := "release-dev"
        projectName := "library_ccompress"
        projectConfig := clay.NewConfig("darwin", "arm64", configName)
        project := clay.NewLibraryProject(projectName, projectConfig)

        // Project Include directories
        project.IncludeDirs = clay.NewIncludeMap(3)
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccompress/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/include")

        // Project Define macros
        configDefines := projectConfig.GetCppDefines()
        project.Defines = clay.NewDefineMap(5 + len(configDefines))
        project.Defines.Add("TARGET_MAC")
        project.Defines.Add("UNICODE")
        project.Defines.Add("_UNICODE")
        project.Defines.Add("TARGET_RELEASE")
        project.Defines.Add("NDEBUG")
        project.Defines.AddMany(configDefines...)

        // Project Source files
        project.SourceFiles = make([]clay.SourceFile, 0, 1)
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccompress/source/main/cpp/c_snappy.cpp", "c_snappy.cpp")

        projects[library_ccompress_release_dev_id] = project
    }
    {
        configName := "debug-dev-test"
        projectName := "unittest_library_cunittest"
        projectConfig := clay.NewConfig("darwin", "arm64", configName)
        project := clay.NewLibraryProject(projectName, projectConfig)

        // Project Include directories
        project.IncludeDirs = clay.NewIncludeMap(1)
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/include")

        // Project Define macros
        configDefines := projectConfig.GetCppDefines()
        project.Defines = clay.NewDefineMap(6 + len(configDefines))
        project.Defines.Add("TARGET_MAC")
        project.Defines.Add("UNICODE")
        project.Defines.Add("_UNICODE")
        project.Defines.Add("TARGET_DEBUG")
        project.Defines.Add("_DEBUG")
        project.Defines.Add("TARGET_TEST")
        project.Defines.AddMany(configDefines...)

        // Project Source files
        project.SourceFiles = make([]clay.SourceFile, 0, 16)
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/entry/ut_Entry_Mac.cpp", "ut_Entry_Mac.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_Checks.cpp", "ut_Checks.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_Config.cpp", "ut_Config.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_File_Mac.cpp", "ut_File_Mac.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_Stdout_Mac.cpp", "ut_Stdout_Mac.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_StringBuilder.cpp", "ut_StringBuilder.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_Test.cpp", "ut_Test.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_TestReporter.cpp", "ut_TestReporter.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_TestReporterStdout.cpp", "ut_TestReporterStdout.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_TestReporterTeamCity.cpp", "ut_TestReporterTeamCity.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_TestResults.cpp", "ut_TestResults.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_Test_Mac.cpp", "ut_Test_Mac.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_TimeConstraint.cpp", "ut_TimeConstraint.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_TimeHelpers_Mac.cpp", "ut_TimeHelpers_Mac.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_Utils.cpp", "ut_Utils.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_Utils_Mac.cpp", "ut_Utils_Mac.cpp")

        projects[unittest_library_cunittest_debug_dev_test_id] = project
    }
    {
        configName := "release-dev-test"
        projectName := "unittest_library_cunittest"
        projectConfig := clay.NewConfig("darwin", "arm64", configName)
        project := clay.NewLibraryProject(projectName, projectConfig)

        // Project Include directories
        project.IncludeDirs = clay.NewIncludeMap(1)
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/include")

        // Project Define macros
        configDefines := projectConfig.GetCppDefines()
        project.Defines = clay.NewDefineMap(6 + len(configDefines))
        project.Defines.Add("TARGET_MAC")
        project.Defines.Add("UNICODE")
        project.Defines.Add("_UNICODE")
        project.Defines.Add("TARGET_RELEASE")
        project.Defines.Add("NDEBUG")
        project.Defines.Add("TARGET_TEST")
        project.Defines.AddMany(configDefines...)

        // Project Source files
        project.SourceFiles = make([]clay.SourceFile, 0, 16)
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/entry/ut_Entry_Mac.cpp", "ut_Entry_Mac.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_Checks.cpp", "ut_Checks.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_Config.cpp", "ut_Config.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_File_Mac.cpp", "ut_File_Mac.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_Stdout_Mac.cpp", "ut_Stdout_Mac.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_StringBuilder.cpp", "ut_StringBuilder.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_Test.cpp", "ut_Test.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_TestReporter.cpp", "ut_TestReporter.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_TestReporterStdout.cpp", "ut_TestReporterStdout.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_TestReporterTeamCity.cpp", "ut_TestReporterTeamCity.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_TestResults.cpp", "ut_TestResults.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_Test_Mac.cpp", "ut_Test_Mac.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_TimeConstraint.cpp", "ut_TimeConstraint.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_TimeHelpers_Mac.cpp", "ut_TimeHelpers_Mac.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_Utils.cpp", "ut_Utils.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/cpp/ut_Utils_Mac.cpp", "ut_Utils_Mac.cpp")

        projects[unittest_library_cunittest_release_dev_test_id] = project
    }
    {
        configName := "debug-dev-test"
        projectName := "unittest_library_ccore"
        projectConfig := clay.NewConfig("darwin", "arm64", configName)
        project := clay.NewLibraryProject(projectName, projectConfig)

        // Project Include directories
        project.IncludeDirs = clay.NewIncludeMap(1)
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/include")

        // Project Define macros
        configDefines := projectConfig.GetCppDefines()
        project.Defines = clay.NewDefineMap(6 + len(configDefines))
        project.Defines.Add("TARGET_MAC")
        project.Defines.Add("UNICODE")
        project.Defines.Add("_UNICODE")
        project.Defines.Add("TARGET_DEBUG")
        project.Defines.Add("_DEBUG")
        project.Defines.Add("TARGET_TEST")
        project.Defines.AddMany(configDefines...)

        // Project Source files
        project.SourceFiles = make([]clay.SourceFile, 0, 10)
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_allocator.cpp", "c_allocator.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_binary_search.cpp", "c_binary_search.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_binmap1.cpp", "c_binmap1.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_debug.cpp", "c_debug.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_error.cpp", "c_error.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_memcpy_neon.cpp", "c_memcpy_neon.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_memcpy_sse.cpp", "c_memcpy_sse.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_memory.cpp", "c_memory.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_qsort.cpp", "c_qsort.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_vmem.cpp", "c_vmem.cpp")

        projects[unittest_library_ccore_debug_dev_test_id] = project
    }
    {
        configName := "release-dev-test"
        projectName := "unittest_library_ccore"
        projectConfig := clay.NewConfig("darwin", "arm64", configName)
        project := clay.NewLibraryProject(projectName, projectConfig)

        // Project Include directories
        project.IncludeDirs = clay.NewIncludeMap(1)
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/include")

        // Project Define macros
        configDefines := projectConfig.GetCppDefines()
        project.Defines = clay.NewDefineMap(6 + len(configDefines))
        project.Defines.Add("TARGET_MAC")
        project.Defines.Add("UNICODE")
        project.Defines.Add("_UNICODE")
        project.Defines.Add("TARGET_RELEASE")
        project.Defines.Add("NDEBUG")
        project.Defines.Add("TARGET_TEST")
        project.Defines.AddMany(configDefines...)

        // Project Source files
        project.SourceFiles = make([]clay.SourceFile, 0, 10)
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_allocator.cpp", "c_allocator.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_binary_search.cpp", "c_binary_search.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_binmap1.cpp", "c_binmap1.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_debug.cpp", "c_debug.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_error.cpp", "c_error.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_memcpy_neon.cpp", "c_memcpy_neon.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_memcpy_sse.cpp", "c_memcpy_sse.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_memory.cpp", "c_memory.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_qsort.cpp", "c_qsort.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/cpp/c_vmem.cpp", "c_vmem.cpp")

        projects[unittest_library_ccore_release_dev_test_id] = project
    }
    {
        configName := "debug-dev-test"
        projectName := "unittest_library_cbase"
        projectConfig := clay.NewConfig("darwin", "arm64", configName)
        project := clay.NewLibraryProject(projectName, projectConfig)

        // Project Include directories
        project.IncludeDirs = clay.NewIncludeMap(2)
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/include")

        // Project Define macros
        configDefines := projectConfig.GetCppDefines()
        project.Defines = clay.NewDefineMap(6 + len(configDefines))
        project.Defines.Add("TARGET_MAC")
        project.Defines.Add("UNICODE")
        project.Defines.Add("_UNICODE")
        project.Defines.Add("TARGET_DEBUG")
        project.Defines.Add("_DEBUG")
        project.Defines.Add("TARGET_TEST")
        project.Defines.AddMany(configDefines...)

        // Project Source files
        project.SourceFiles = make([]clay.SourceFile, 0, 32)
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_allocator.cpp", "c_allocator.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_allocator_system.cpp", "c_allocator_system.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_base.cpp", "c_base.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_binary_search.cpp", "c_binary_search.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_binmap.cpp", "c_binmap.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_buffer.cpp", "c_buffer.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_console.cpp", "c_console.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_console_mac.cpp", "c_console_mac.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_context.cpp", "c_context.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_dconv.cpp", "c_dconv.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_debug.cpp", "c_debug.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_duomap.cpp", "c_duomap.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_guid.cpp", "c_guid.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_hash.cpp", "c_hash.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_log.cpp", "c_log.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_log_to_console.cpp", "c_log_to_console.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_map.cpp", "c_map.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_memory.cpp", "c_memory.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_pdqsort.cpp", "c_pdqsort.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_printf.cpp", "c_printf.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_random.cpp", "c_random.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_runes.cpp", "c_runes.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_slice.cpp", "c_slice.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_sscanf.cpp", "c_sscanf.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_strfmt.cpp", "c_strfmt.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_tblfmt.cpp", "c_tblfmt.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_tree.cpp", "c_tree.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_tree32.cpp", "c_tree32.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_tree_obsolete.cpp", "c_tree_obsolete.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_va_list.cpp", "c_va_list.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_vector.cpp", "c_vector.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_wyhash.cpp", "c_wyhash.cpp")

        projects[unittest_library_cbase_debug_dev_test_id] = project
    }
    {
        configName := "release-dev-test"
        projectName := "unittest_library_cbase"
        projectConfig := clay.NewConfig("darwin", "arm64", configName)
        project := clay.NewLibraryProject(projectName, projectConfig)

        // Project Include directories
        project.IncludeDirs = clay.NewIncludeMap(2)
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/include")

        // Project Define macros
        configDefines := projectConfig.GetCppDefines()
        project.Defines = clay.NewDefineMap(6 + len(configDefines))
        project.Defines.Add("TARGET_MAC")
        project.Defines.Add("UNICODE")
        project.Defines.Add("_UNICODE")
        project.Defines.Add("TARGET_RELEASE")
        project.Defines.Add("NDEBUG")
        project.Defines.Add("TARGET_TEST")
        project.Defines.AddMany(configDefines...)

        // Project Source files
        project.SourceFiles = make([]clay.SourceFile, 0, 32)
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_allocator.cpp", "c_allocator.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_allocator_system.cpp", "c_allocator_system.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_base.cpp", "c_base.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_binary_search.cpp", "c_binary_search.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_binmap.cpp", "c_binmap.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_buffer.cpp", "c_buffer.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_console.cpp", "c_console.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_console_mac.cpp", "c_console_mac.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_context.cpp", "c_context.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_dconv.cpp", "c_dconv.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_debug.cpp", "c_debug.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_duomap.cpp", "c_duomap.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_guid.cpp", "c_guid.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_hash.cpp", "c_hash.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_log.cpp", "c_log.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_log_to_console.cpp", "c_log_to_console.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_map.cpp", "c_map.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_memory.cpp", "c_memory.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_pdqsort.cpp", "c_pdqsort.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_printf.cpp", "c_printf.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_random.cpp", "c_random.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_runes.cpp", "c_runes.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_slice.cpp", "c_slice.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_sscanf.cpp", "c_sscanf.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_strfmt.cpp", "c_strfmt.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_tblfmt.cpp", "c_tblfmt.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_tree.cpp", "c_tree.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_tree32.cpp", "c_tree32.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_tree_obsolete.cpp", "c_tree_obsolete.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_va_list.cpp", "c_va_list.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_vector.cpp", "c_vector.cpp")
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/cpp/c_wyhash.cpp", "c_wyhash.cpp")

        projects[unittest_library_cbase_release_dev_test_id] = project
    }
    {
        configName := "debug-dev-test"
        projectName := "unittest_library_ccompress"
        projectConfig := clay.NewConfig("darwin", "arm64", configName)
        project := clay.NewLibraryProject(projectName, projectConfig)

        // Project Include directories
        project.IncludeDirs = clay.NewIncludeMap(3)
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccompress/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/include")

        // Project Define macros
        configDefines := projectConfig.GetCppDefines()
        project.Defines = clay.NewDefineMap(6 + len(configDefines))
        project.Defines.Add("TARGET_MAC")
        project.Defines.Add("UNICODE")
        project.Defines.Add("_UNICODE")
        project.Defines.Add("TARGET_DEBUG")
        project.Defines.Add("_DEBUG")
        project.Defines.Add("TARGET_TEST")
        project.Defines.AddMany(configDefines...)

        // Project Source files
        project.SourceFiles = make([]clay.SourceFile, 0, 1)
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccompress/source/main/cpp/c_snappy.cpp", "c_snappy.cpp")

        projects[unittest_library_ccompress_debug_dev_test_id] = project
    }
    {
        configName := "release-dev-test"
        projectName := "unittest_library_ccompress"
        projectConfig := clay.NewConfig("darwin", "arm64", configName)
        project := clay.NewLibraryProject(projectName, projectConfig)

        // Project Include directories
        project.IncludeDirs = clay.NewIncludeMap(3)
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccompress/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/include")

        // Project Define macros
        configDefines := projectConfig.GetCppDefines()
        project.Defines = clay.NewDefineMap(6 + len(configDefines))
        project.Defines.Add("TARGET_MAC")
        project.Defines.Add("UNICODE")
        project.Defines.Add("_UNICODE")
        project.Defines.Add("TARGET_RELEASE")
        project.Defines.Add("NDEBUG")
        project.Defines.Add("TARGET_TEST")
        project.Defines.AddMany(configDefines...)

        // Project Source files
        project.SourceFiles = make([]clay.SourceFile, 0, 1)
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccompress/source/main/cpp/c_snappy.cpp", "c_snappy.cpp")

        projects[unittest_library_ccompress_release_dev_test_id] = project
    }
    {
        configName := "debug-dev-test"
        projectName := "unittest_ccompress"
        projectConfig := clay.NewConfig("darwin", "arm64", configName)
        project := clay.NewExecutableProject(projectName, projectConfig)

        // Project Include directories
        project.IncludeDirs = clay.NewIncludeMap(5)
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccompress/source/test/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccompress/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/include")

        // Project Define macros
        configDefines := projectConfig.GetCppDefines()
        project.Defines = clay.NewDefineMap(6 + len(configDefines))
        project.Defines.Add("TARGET_MAC")
        project.Defines.Add("UNICODE")
        project.Defines.Add("_UNICODE")
        project.Defines.Add("TARGET_DEBUG")
        project.Defines.Add("_DEBUG")
        project.Defines.Add("TARGET_TEST")
        project.Defines.AddMany(configDefines...)

        // Project Source files
        project.SourceFiles = make([]clay.SourceFile, 0, 1)
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccompress/source/test/cpp/test_main.cpp", "test_main.cpp")

        projects[unittest_ccompress_debug_dev_test_id] = project
    }
    {
        configName := "release-dev-test"
        projectName := "unittest_ccompress"
        projectConfig := clay.NewConfig("darwin", "arm64", configName)
        project := clay.NewExecutableProject(projectName, projectConfig)

        // Project Include directories
        project.IncludeDirs = clay.NewIncludeMap(5)
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccompress/source/test/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccompress/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cunittest/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccore/source/main/include")
        project.IncludeDirs.Add("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/cbase/source/main/include")

        // Project Define macros
        configDefines := projectConfig.GetCppDefines()
        project.Defines = clay.NewDefineMap(6 + len(configDefines))
        project.Defines.Add("TARGET_MAC")
        project.Defines.Add("UNICODE")
        project.Defines.Add("_UNICODE")
        project.Defines.Add("TARGET_RELEASE")
        project.Defines.Add("NDEBUG")
        project.Defines.Add("TARGET_TEST")
        project.Defines.AddMany(configDefines...)

        // Project Source files
        project.SourceFiles = make([]clay.SourceFile, 0, 1)
        project.AddSourceFile("/Users/obnosis5/dev.go/src/github.com/jurgen-kluft/ccompress/source/test/cpp/test_main.cpp", "test_main.cpp")

        projects[unittest_ccompress_release_dev_test_id] = project
    }
    // Setup Project Dependencies
    {
        project := projects[library_cbase_debug_dev_id]
        project.Dependencies = []*clay.Project{
            projects[library_ccore_debug_dev_id],
        }
    }
    {
        project := projects[library_cbase_release_dev_id]
        project.Dependencies = []*clay.Project{
            projects[library_ccore_release_dev_id],
        }
    }
    {
        project := projects[library_ccompress_debug_dev_id]
        project.Dependencies = []*clay.Project{
            projects[library_ccore_debug_dev_id],
            projects[library_cbase_debug_dev_id],
        }
    }
    {
        project := projects[library_ccompress_release_dev_id]
        project.Dependencies = []*clay.Project{
            projects[library_ccore_release_dev_id],
            projects[library_cbase_release_dev_id],
        }
    }
    {
        project := projects[unittest_library_cbase_debug_dev_test_id]
        project.Dependencies = []*clay.Project{
            projects[unittest_library_ccore_debug_dev_test_id],
        }
    }
    {
        project := projects[unittest_library_cbase_release_dev_test_id]
        project.Dependencies = []*clay.Project{
            projects[unittest_library_ccore_release_dev_test_id],
        }
    }
    {
        project := projects[unittest_library_ccompress_debug_dev_test_id]
        project.Dependencies = []*clay.Project{
            projects[unittest_library_ccore_debug_dev_test_id],
            projects[unittest_library_cbase_debug_dev_test_id],
        }
    }
    {
        project := projects[unittest_library_ccompress_release_dev_test_id]
        project.Dependencies = []*clay.Project{
            projects[unittest_library_ccore_release_dev_test_id],
            projects[unittest_library_cbase_release_dev_test_id],
        }
    }
    {
        project := projects[unittest_ccompress_debug_dev_test_id]
        project.Dependencies = []*clay.Project{
            projects[unittest_library_cunittest_debug_dev_test_id],
            projects[unittest_library_ccore_debug_dev_test_id],
            projects[unittest_library_cbase_debug_dev_test_id],
            projects[unittest_library_ccompress_debug_dev_test_id],
        }
    }
    {
        project := projects[unittest_ccompress_release_dev_test_id]
        project.Dependencies = []*clay.Project{
            projects[unittest_library_cunittest_release_dev_test_id],
            projects[unittest_library_ccore_release_dev_test_id],
            projects[unittest_library_cbase_release_dev_test_id],
            projects[unittest_library_ccompress_release_dev_test_id],
        }
    }
    return projects
}

