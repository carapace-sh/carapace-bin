package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// AddBuildOptions adds the shared build options for build-exe, build-lib, build-obj, test, test-obj, run, translate-c
func AddBuildOptions(cmd *cobra.Command) {
	// General Options
	cmd.Flags().String("cache-dir", "", "Override the local cache directory")
	cmd.Flags().String("color", "", "Enable or disable colored error messages")
	cmd.Flags().String("emit-asm", "", "Output .s (assembly code)")
	cmd.Flags().String("emit-bin", "", "Output machine code")
	cmd.Flags().String("emit-docs", "", "Create a docs/ dir with html documentation")
	cmd.Flags().String("emit-h", "", "Generate a C header file (.h)")
	cmd.Flags().String("emit-implib", "", "Produce an import .lib when building a Windows DLL")
	cmd.Flags().String("emit-llvm-bc", "", "Produce an optimized LLVM module as a .bc file")
	cmd.Flags().String("emit-llvm-ir", "", "Produce a .ll file with optimized LLVM IR")
	cmd.Flags().String("global-cache-dir", "", "Override the global cache directory")
	cmd.Flags().BoolP("help", "h", false, "Print this help and exit")
	cmd.Flags().Bool("incremental", false, "Enable incremental compilation")
	cmd.Flags().String("job-count", "", "Limit concurrent jobs (default is to use all CPU cores)")
	cmd.Flags().Bool("no-emit-asm", false, "Do not output .s (assembly code)")
	cmd.Flags().Bool("no-emit-bin", false, "Do not output machine code")
	cmd.Flags().Bool("no-emit-docs", false, "Do not produce docs/ dir with html documentation")
	cmd.Flags().Bool("no-emit-h", false, "Do not generate a C header file (.h)")
	cmd.Flags().Bool("no-emit-implib", false, "Do not produce an import .lib when building a Windows DLL")
	cmd.Flags().Bool("no-emit-llvm-bc", false, "Do not produce an optimized LLVM module as a .bc file")
	cmd.Flags().Bool("no-emit-llvm-ir", false, "Do not produce a .ll file with optimized LLVM IR")
	cmd.Flags().Bool("no-incremental", false, "Disable incremental compilation")
	cmd.Flags().Bool("show-builtin", false, "Output the source of @import(\"builtin\") then exit")
	cmd.Flags().String("zig-lib-dir", "", "Override path to Zig installation lib directory")

	cmd.Flags().Lookup("emit-asm").NoOptDefVal = " "
	cmd.Flags().Lookup("emit-bin").NoOptDefVal = " "
	cmd.Flags().Lookup("emit-docs").NoOptDefVal = " "
	cmd.Flags().Lookup("emit-h").NoOptDefVal = " "
	cmd.Flags().Lookup("emit-implib").NoOptDefVal = " "
	cmd.Flags().Lookup("emit-llvm-bc").NoOptDefVal = " "
	cmd.Flags().Lookup("emit-llvm-ir").NoOptDefVal = " "

	// Global Compile Options
	cmd.Flags().Bool("PIE", false, "Force-enable Position Independent Executable")
	cmd.Flags().Bool("clang", false, "Force using Clang as the C/C++ compilation backend")
	cmd.Flags().Bool("data-sections", false, "Places each data in a separate section")
	cmd.Flags().String("dep", "", "Add an entry to the next module's import table")
	cmd.Flags().Bool("dll-export-fns", false, "Mark exported functions as DLL exports (Windows)")
	cmd.Flags().String("error-limit", "", "Set the maximum amount of distinct error values")
	cmd.Flags().Bool("function-sections", false, "Places each function in a separate section")
	cmd.Flags().String("libc", "", "Provide a file which specifies libc paths")
	cmd.Flags().Bool("libllvm", false, "Force using the LLVM API in the codegen backend")
	cmd.Flags().Bool("llvm", false, "Force using LLVM as the codegen backend")
	cmd.Flags().Bool("lto", false, "Force-enable Link Time Optimization")
	cmd.Flags().String("lto-mode", "", "Set LTO mode")
	cmd.Flags().String("mexec-model", "", "Execution model (WASI)")
	cmd.Flags().Bool("municode", false, "Use wmain/wWinMain as entry point (Windows)")
	cmd.Flags().String("name", "", "Compilation unit name (not a file path)")
	cmd.Flags().Bool("no-PIE", false, "Force-disable Position Independent Executable")
	cmd.Flags().Bool("no-clang", false, "Prevent using Clang as the C/C++ compilation backend")
	cmd.Flags().Bool("no-data-sections", false, "All data go into same section")
	cmd.Flags().Bool("no-dll-export-fns", false, "Force-disable marking exported functions as DLL exports")
	cmd.Flags().Bool("no-function-sections", false, "All functions go into same section")
	cmd.Flags().Bool("no-libllvm", false, "Prevent using the LLVM API in the codegen backend")
	cmd.Flags().Bool("no-llvm", false, "Prevent using LLVM as the codegen backend")
	cmd.Flags().Bool("no-lto", false, "Force-disable Link Time Optimization")
	cmd.Flags().Bool("no-reference-trace", false, "Disable reference trace")
	cmd.Flags().Bool("no-structured-cfg", false, "Force SPIR-V kernels to not use structured control flow")
	cmd.Flags().String("reference-trace", "", "Show num lines of reference trace per compile error")
	cmd.Flags().Bool("structured-cfg", false, "Force SPIR-V kernels to use structured control flow")
	cmd.Flags().Bool("time-report", false, "Send timing diagnostics to '--listen' clients")

	cmd.Flags().Lookup("reference-trace").NoOptDefVal = " "

	// Per-Module Compile Options
	cmd.Flags().Bool("PIC", false, "Force-enable Position Independent Code")
	cmd.Flags().Bool("async-unwind-tables", false, "Always produce asynchronous unwind table entries for all functions")
	cmd.Flags().Bool("builtin", false, "Enable implicit builtin knowledge of functions")
	cmd.Flags().String("cflags", "", "Set extra flags for the next positional C source files")
	cmd.Flags().StringSlice("define", nil, "Define C macro to value")
	cmd.Flags().String("embed-dir", "", "Add directory to embed search path")
	cmd.Flags().Bool("error-tracing", false, "Enable error tracing in release builds")
	cmd.Flags().Bool("fuzz", false, "Enable fuzz testing instrumentation")
	cmd.Flags().String("idirafter", "", "Add directory to AFTER include search path")
	cmd.Flags().StringSlice("include", nil, "Add directory to include search path")
	cmd.Flags().String("isystem", "", "Add directory to SYSTEM include search path")
	cmd.Flags().String("mcmodel", "", "Limit range of code and data virtual addresses")
	cmd.Flags().String("mcpu", "", "Specify target CPU and feature set")
	cmd.Flags().Bool("mno-red-zone", false, "Force-disable the \"red-zone\"")
	cmd.Flags().Bool("mred-zone", false, "Force-enable the \"red-zone\"")
	cmd.Flags().Bool("no-PIC", false, "Force-disable Position Independent Code")
	cmd.Flags().Bool("no-builtin", false, "Disable implicit builtin knowledge of functions")
	cmd.Flags().Bool("no-error-tracing", false, "Disable error tracing in debug builds")
	cmd.Flags().Bool("no-fuzz", false, "Disable fuzz testing instrumentation")
	cmd.Flags().Bool("no-omit-frame-pointer", false, "Store the stack frame pointer")
	cmd.Flags().Bool("no-sanitize-c", false, "Disable C undefined behavior detection in safe builds")
	cmd.Flags().Bool("no-sanitize-thread", false, "Disable Thread Sanitizer")
	cmd.Flags().Bool("no-single-threaded", false, "Code may not assume there is only one thread")
	cmd.Flags().Bool("no-stack-check", false, "Disable stack probing in safe builds")
	cmd.Flags().Bool("no-stack-protector", false, "Disable stack protection in safe builds")
	cmd.Flags().Bool("no-strip", false, "Keep debug symbols")
	cmd.Flags().Bool("no-unwind-tables", false, "Never produce unwind table entries")
	cmd.Flags().Bool("no-valgrind", false, "Omit valgrind client requests in debug builds")
	cmd.Flags().String("ofmt", "", "Override target object format")
	cmd.Flags().Bool("omit-frame-pointer", false, "Omit the stack frame pointer")
	cmd.Flags().String("optimize", "", "Choose what to optimize for")
	cmd.Flags().String("rcflags", "", "Set extra flags for the next positional .rc source files")
	cmd.Flags().String("rcincludes", "", "Set the type of includes for .rc files")
	cmd.Flags().String("sanitize-c", "", "Enable C undefined behavior detection")
	cmd.Flags().Bool("sanitize-thread", false, "Enable Thread Sanitizer")
	cmd.Flags().Bool("single-threaded", false, "Code assumes there is only one thread")
	cmd.Flags().Bool("stack-check", false, "Enable stack probing in unsafe builds")
	cmd.Flags().Bool("stack-protector", false, "Enable stack protection in unsafe builds")
	cmd.Flags().Bool("strip", false, "Omit debug symbols")
	cmd.Flags().String("target", "", "Target triple (e.g. x86_64-linux-gnu)")
	cmd.Flags().Bool("unwind-tables", false, "Always produce unwind table entries for all functions")
	cmd.Flags().Bool("valgrind", false, "Include valgrind client requests in release builds")

	cmd.Flags().Lookup("sanitize-c").NoOptDefVal = " "

	// Global Link Options
	cmd.Flags().Bool("Bsymbolic", false, "Bind global references locally")
	cmd.Flags().Bool("ObjC", false, "Darwin: force load all members of static archives that implement ObjC class/category")
	cmd.Flags().Bool("allow-shlib-undefined", false, "Allows undefined symbols in shared libraries")
	cmd.Flags().Bool("allow-so-scripts", false, "Allows .so files to be GNU ld scripts")
	cmd.Flags().String("build-id", "", "Embed a build ID in binaries")
	cmd.Flags().Bool("compiler-rt", false, "Always include compiler-rt symbols in output")
	cmd.Flags().String("compress-debug-sections", "", "Debug section compression")
	cmd.Flags().Bool("dead-strip", false, "Darwin: remove functions and data unreachable by entry point or exported symbols")
	cmd.Flags().Bool("dead-strip-dylibs", false, "Darwin: remove dylibs unreachable by entry point or exported symbols")
	cmd.Flags().Bool("disable-new-dtags", false, "Use the old behavior for dynamic tags (RPATH)")
	cmd.Flags().Bool("dynamic", false, "Force output to be dynamically linked")
	cmd.Flags().String("dynamic-linker", "", "Set the dynamic interpreter path (usually ld.so)")
	cmd.Flags().Bool("each-lib-rpath", false, "Ensure adding rpath for each used dynamic library")
	cmd.Flags().Bool("eh-frame-hdr", false, "Enable C++ exception handling by passing --eh-frame-hdr to linker")
	cmd.Flags().Bool("emit-relocs", false, "Enable output of relocation sections for post build tools")
	cmd.Flags().Bool("enable-new-dtags", false, "Use the new behavior for dynamic tags (RUNPATH)")
	cmd.Flags().String("entitlements", "", "Darwin: add path to entitlements file for embedding in code signature")
	cmd.Flags().String("entry", "", "Enable entry point with default symbol name")
	cmd.Flags().String("export", "", "WebAssembly: Force a symbol to be exported")
	cmd.Flags().Bool("export-memory", false, "WebAssembly: export memory to the host")
	cmd.Flags().Bool("export-table", false, "WebAssembly: export function table to the host environment")
	cmd.Flags().String("force-undefined", "", "Specify the symbol must be defined for the link to succeed")
	cmd.Flags().Bool("gc-sections", false, "Force removal of unreachable functions and data")
	cmd.Flags().String("global-base", "", "WebAssembly: where to start placing global data")
	cmd.Flags().String("headerpad", "", "Darwin: set minimum space for future expansion of load commands")
	cmd.Flags().Bool("headerpad-max-install-names", false, "Darwin: set enough space as if all paths were MAXPATHLEN")
	cmd.Flags().String("image-base", "", "Set base address for executable image")
	cmd.Flags().Bool("import-memory", false, "WebAssembly: import memory from the environment")
	cmd.Flags().Bool("import-symbols", false, "WebAssembly: import missing symbols from the host environment")
	cmd.Flags().Bool("import-table", false, "WebAssembly: import function table from the host environment")
	cmd.Flags().String("initial-memory", "", "WebAssembly: initial size of linear memory")
	cmd.Flags().String("install-name", "", "Darwin: add dylib's install name")
	cmd.Flags().Bool("lld", false, "Force using LLD as the linker")
	cmd.Flags().String("max-memory", "", "WebAssembly: maximum size of linear memory")
	cmd.Flags().Bool("new-linker", false, "Force using the new linker (Zig linker)")
	cmd.Flags().Bool("no-allow-shlib-undefined", false, "Disallows undefined symbols in shared libraries")
	cmd.Flags().Bool("no-allow-so-scripts", false, ".so files must be ELF files")
	cmd.Flags().Bool("no-compiler-rt", false, "Prevent including compiler-rt symbols in output")
	cmd.Flags().Bool("no-dynamic-linker", false, "Do not set any dynamic interpreter path")
	cmd.Flags().Bool("no-each-lib-rpath", false, "Prevent adding rpath for each used dynamic library")
	cmd.Flags().Bool("no-eh-frame-hdr", false, "Disable C++ exception handling by passing --no-eh-frame-hdr to linker")
	cmd.Flags().Bool("no-entry", false, "Do not output any entry point")
	cmd.Flags().Bool("no-gc-sections", false, "Don't force removal of unreachable functions and data")
	cmd.Flags().Bool("no-lld", false, "Prevent using LLD as the linker")
	cmd.Flags().Bool("no-new-linker", false, "Prevent using the new linker")
	cmd.Flags().Bool("no-soname", false, "Disable emitting a SONAME")
	cmd.Flags().Bool("no-ubsan-rt", false, "Prevent including ubsan-rt symbols in the output")
	cmd.Flags().Bool("no-undefined-version", false, "Disallow version scripts from referring to undefined symbols")
	cmd.Flags().String("pagezero-size", "", "Darwin: size of __PAGEZERO segment in hexadecimal")
	cmd.Flags().Bool("rdynamic", false, "Add all symbols to the dynamic symbol table")
	cmd.Flags().String("script", "", "Use a custom linker script")
	cmd.Flags().Bool("shared-memory", false, "WebAssembly: use shared linear memory")
	cmd.Flags().String("soname", "", "Override the default SONAME value")
	cmd.Flags().String("sort-section", "", "Sort wildcard section patterns by 'name' or 'alignment'")
	cmd.Flags().String("stack", "", "Override default stack size")
	cmd.Flags().Bool("static", false, "Force output to be statically linked")
	cmd.Flags().String("subsystem", "", "Windows /SUBSYSTEM:<subsystem> to the linker")
	cmd.Flags().String("sysroot", "", "Set the system root directory (usually /)")
	cmd.Flags().Bool("ubsan-rt", false, "Always include ubsan-rt symbols in the output")
	cmd.Flags().Bool("undefined-version", false, "Allow version scripts to refer to undefined symbols")
	cmd.Flags().String("version", "", "Dynamic library semver")
	cmd.Flags().String("version-script", "", "Provide a version .map file")

	cmd.Flags().Lookup("build-id").NoOptDefVal = " "
	cmd.Flags().Lookup("entry").NoOptDefVal = " "
	cmd.Flags().Lookup("soname").NoOptDefVal = " "

	// Per-Module Link Options
	cmd.Flags().String("framework", "", "Darwin: link against framework")
	cmd.Flags().StringSlice("library", nil, "Link against system library (only if actually used)")
	cmd.Flags().String("library-directory", "", "Add a directory to the library search path")
	cmd.Flags().String("needed-framework", "", "Darwin: link against framework (even if unused)")
	cmd.Flags().StringSlice("needed-library", nil, "Link against system library (even if unused)")
	cmd.Flags().StringSlice("rpath", nil, "Add directory to the runtime library search path")
	cmd.Flags().String("weak-framework", "", "Darwin: link against framework and mark it and all referenced symbols as weak")
	cmd.Flags().StringSlice("weak-library", nil, "Link against system library marking it and all referenced symbols as weak")

	// Debug Options
	cmd.Flags().Bool("debug-compile-errors", false, "Crash with helpful diagnostics at the first compile error")
	cmd.Flags().Bool("debug-incremental", false, "Enable incremental compilation debug features")
	cmd.Flags().Bool("debug-link-snapshot", false, "Enable dumping of the linker's state in JSON format")
	cmd.Flags().String("debug-log", "", "Enable printing debug/info log messages for scope")
	cmd.Flags().String("debug-rt", "", "Build compiler runtime libraries with optimization")
	cmd.Flags().String("listen", "", "Listen on address for IDE integration")
	cmd.Flags().String("opt-bisect-limit", "", "Only run first N LLVM optimization passes")
	cmd.Flags().Bool("stack-report", false, "Print stack size diagnostics")
	cmd.Flags().Bool("verbose-air", false, "Enable compiler debug output for Zig AIR")
	cmd.Flags().Bool("verbose-cc", false, "Display C compiler invocations")
	cmd.Flags().Bool("verbose-cimport", false, "Enable compiler debug output for C imports")
	cmd.Flags().Bool("verbose-generic-instances", false, "Enable compiler debug output for generic instance generation")
	cmd.Flags().Bool("verbose-intern-pool", false, "Enable compiler debug output for InternPool")
	cmd.Flags().Bool("verbose-link", false, "Display linker invocations")
	cmd.Flags().String("verbose-llvm-bc", "", "Enable compiler debug output for unoptimized LLVM BC")
	cmd.Flags().Bool("verbose-llvm-cpu-features", false, "Enable compiler debug output for LLVM CPU features")
	cmd.Flags().String("verbose-llvm-ir", "", "Enable compiler debug output for unoptimized LLVM IR")

	cmd.Flags().Lookup("debug-rt").NoOptDefVal = " "
	cmd.Flags().Lookup("opt-bisect-limit").NoOptDefVal = " "
	cmd.Flags().Lookup("verbose-llvm-bc").NoOptDefVal = " "
	cmd.Flags().Lookup("verbose-llvm-ir").NoOptDefVal = " "

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"build-id":                carapace.ActionValues("fast", "sha1", "tree", "md5", "uuid", "none"),
		"cache-dir":               carapace.ActionDirectories(),
		"color":                   carapace.ActionValues("auto", "off", "on"),
		"compress-debug-sections": carapace.ActionValues("none", "zlib", "zstd"),
		"debug-rt":                carapace.ActionValues("Debug", "ReleaseFast", "ReleaseSafe", "ReleaseSmall"),
		"dynamic-linker":          carapace.ActionFiles(),
		"embed-dir":               carapace.ActionDirectories(),
		"emit-asm":                carapace.ActionFiles(),
		"emit-bin":                carapace.ActionFiles(),
		"emit-docs":               carapace.ActionDirectories(),
		"emit-h":                  carapace.ActionFiles(),
		"emit-implib":             carapace.ActionFiles(),
		"emit-llvm-bc":            carapace.ActionFiles(),
		"emit-llvm-ir":            carapace.ActionFiles(),
		"entitlements":            carapace.ActionFiles(),
		"entry":                   carapace.ActionValues(),
		"global-cache-dir":        carapace.ActionDirectories(),
		"idirafter":               carapace.ActionDirectories(),
		"include":                 carapace.ActionDirectories(),
		"isystem":                 carapace.ActionDirectories(),
		"libc":                    carapace.ActionFiles(),
		"library-directory":       carapace.ActionDirectories(),
		"listen":                  carapace.ActionValues("-"),
		"lto-mode":                carapace.ActionValues("full", "thin"),
		"mcmodel":                 carapace.ActionValues("default", "extreme", "kernel", "large", "medany", "medium", "medlow", "medmid", "normal", "small", "tiny"),
		"ofmt":                    carapace.ActionValues("elf", "c", "wasm", "coff", "macho", "spirv", "plan9"),
		"optimize":                carapace.ActionValues("Debug", "ReleaseFast", "ReleaseSafe", "ReleaseSmall"),
		"rcincludes":              carapace.ActionValues("any", "msvc", "gnu", "none"),
		"reference-trace":         carapace.ActionValues(),
		"rpath":                   carapace.ActionDirectories(),
		"sanitize-c":              carapace.ActionValues("trap", "full"),
		"script":                  carapace.ActionFiles(),
		"soname":                  carapace.ActionValues(),
		"sort-section":            carapace.ActionValues("name", "alignment"),
		"sysroot":                 carapace.ActionDirectories(),
		"verbose-llvm-bc":         carapace.ActionFiles(),
		"verbose-llvm-ir":         carapace.ActionFiles(),
		"version-script":          carapace.ActionFiles(),
		"zig-lib-dir":             carapace.ActionDirectories(),
	})
}

// AddTestOptions adds test-specific options
func AddTestOptions(cmd *cobra.Command) {
	cmd.Flags().StringSlice("test-cmd", nil, "Specify test execution command one arg at a time")
	cmd.Flags().Bool("test-cmd-bin", false, "Appends test binary path to test cmd args")
	cmd.Flags().String("test-filter", "", "Skip tests that do not match any filter")
	cmd.Flags().Bool("test-no-exec", false, "Compiles test binary without running it")
	cmd.Flags().String("test-runner", "", "Specify a custom test runner")
}

// AddCompilerOptions adds compiler flags for cc/c++ commands
func AddCompilerOptions(cmd *cobra.Command) {
	cmd.Flags().String("abi", "", "Target ABI")
	cmd.Flags().String("arch", "", "Target architecture")
	cmd.Flags().String("color", "", "Enable or disable colored error messages")
	cmd.Flags().BoolP("compile", "c", false, "Only compile, do not link")
	cmd.Flags().StringSliceP("define", "D", nil, "Define macro")
	cmd.Flags().String("gcc-toolchain", "", "GCC toolchain path")
	cmd.Flags().BoolP("help", "h", false, "Print help")
	cmd.Flags().String("idirafter", "", "Add after include path")
	cmd.Flags().StringSliceP("include", "I", nil, "Add include path")
	cmd.Flags().String("isystem", "", "Add system include path")
	cmd.Flags().StringSliceP("library", "l", nil, "Link with library")
	cmd.Flags().StringP("library-directory", "L", "", "Add library search path")
	cmd.Flags().String("msvc-toolchain", "", "MSVC toolchain path")
	cmd.Flags().Bool("nostdinc", false, "Remove standard include paths")
	cmd.Flags().Bool("nostdlibinc", false, "Remove standard library include paths")
	cmd.Flags().String("os", "", "Target operating system")
	cmd.Flags().StringP("output", "o", "", "Output file")
	cmd.Flags().BoolP("preprocess", "E", false, "Preprocess only")
	cmd.Flags().String("sysroot", "", "Set system root")
	cmd.Flags().StringP("target", "t", "", "Target triple (e.g. x86_64-linux-gnu)")
	cmd.Flags().BoolP("verbose", "v", false, "Print commands to stderr")
	cmd.Flags().Bool("verbose-cc", false, "Display C compiler invocations")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"color":             carapace.ActionValues("auto", "off", "on"),
		"target":            carapace.ActionValues(),
		"library-directory": carapace.ActionDirectories(),
		"sysroot":           carapace.ActionDirectories(),
	})
}
