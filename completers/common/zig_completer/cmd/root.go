package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zig",
	Short: "General-purpose programming language and toolchain",
	Long:  "https://ziglang.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("cache-dir", "", "Override the local cache directory")
	rootCmd.PersistentFlags().String("color", "", "Enable or disable colored error messages")
	rootCmd.PersistentFlags().StringS("femit-asm", "femit-asm", "", "Output assembly code")
	rootCmd.PersistentFlags().StringS("femit-bin", "femit-bin", "", "Output machine code")
	rootCmd.PersistentFlags().StringS("femit-docs", "femit-docs", "", "Create a docs directory with HTML documentation")
	rootCmd.PersistentFlags().StringS("femit-h", "femit-h", "", "Generate a C header file")
	rootCmd.PersistentFlags().StringS("femit-implib", "femit-implib", "", "Produce a Windows import library")
	rootCmd.PersistentFlags().StringS("femit-llvm-bc", "femit-llvm-bc", "", "Produce an LLVM bitcode file")
	rootCmd.PersistentFlags().StringS("femit-llvm-ir", "femit-llvm-ir", "", "Produce an LLVM IR file")
	rootCmd.PersistentFlags().BoolS("fincremental", "fincremental", false, "Enable incremental compilation")
	rootCmd.PersistentFlags().BoolS("fno-emit-asm", "fno-emit-asm", false, "Do not output assembly code")
	rootCmd.PersistentFlags().BoolS("fno-emit-bin", "fno-emit-bin", false, "Do not output machine code")
	rootCmd.PersistentFlags().BoolS("fno-emit-docs", "fno-emit-docs", false, "Do not produce documentation")
	rootCmd.PersistentFlags().BoolS("fno-emit-h", "fno-emit-h", false, "Do not generate a C header file")
	rootCmd.PersistentFlags().BoolS("fno-emit-implib", "fno-emit-implib", false, "Do not produce a Windows import library")
	rootCmd.PersistentFlags().BoolS("fno-emit-llvm-bc", "fno-emit-llvm-bc", false, "Do not produce an LLVM bitcode file")
	rootCmd.PersistentFlags().BoolS("fno-emit-llvm-ir", "fno-emit-llvm-ir", false, "Do not produce an LLVM IR file")
	rootCmd.PersistentFlags().BoolS("fno-incremental", "fno-incremental", false, "Disable incremental compilation")
	rootCmd.PersistentFlags().String("global-cache-dir", "", "Override the global cache directory")
	rootCmd.PersistentFlags().StringS("j", "j", "", "Limit concurrent jobs")
	rootCmd.PersistentFlags().Bool("show-builtin", false, "Output the source of @import(\"builtin\") then exit")
	rootCmd.PersistentFlags().String("zig-lib-dir", "", "Override path to Zig installation lib directory")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cache-dir":        carapace.ActionDirectories(),
		"color":            actionColorModes(),
		"femit-asm":        carapace.ActionFiles(),
		"femit-bin":        carapace.ActionFiles(),
		"femit-docs":       carapace.ActionDirectories(),
		"femit-h":          carapace.ActionFiles(".h"),
		"femit-implib":     carapace.ActionFiles(".lib"),
		"femit-llvm-bc":    carapace.ActionFiles(".bc"),
		"femit-llvm-ir":    carapace.ActionFiles(".ll"),
		"global-cache-dir": carapace.ActionDirectories(),
		"zig-lib-dir":      carapace.ActionDirectories(),
	})
}

func actionColorModes() carapace.Action {
	return carapace.ActionValues("on", "off", "auto").StyleF(style.ForKeyword)
}

func actionOptimizationModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"Debug", "optimizations off, safety on",
		"ReleaseFast", "optimizations on, safety off",
		"ReleaseSafe", "optimizations on, safety on",
		"ReleaseSmall", "optimize for small binary, safety off",
	).StyleF(style.ForKeyword)
}

func actionObjectFormats() carapace.Action {
	return carapace.ActionValuesDescribed(
		"elf", "Executable and Linking Format",
		"c", "C source code",
		"wasm", "WebAssembly",
		"coff", "Common Object File Format",
		"macho", "macOS relocatables",
		"spirv", "Standard Portable Intermediate Representation V",
		"plan9", "Plan 9 object format",
		"hex", "Intel IHEX",
		"raw", "raw machine code",
	).StyleF(style.ForKeyword)
}

func actionBuildIDs() carapace.Action {
	return carapace.ActionValuesDescribed(
		"fast", "8-byte non-cryptographic hash",
		"sha1", "20-byte cryptographic hash",
		"tree", "20-byte cryptographic hash",
		"md5", "16-byte cryptographic hash",
		"uuid", "16-byte random UUID",
		"none", "no build ID",
	).StyleF(style.ForKeyword)
}

func actionZigSourceFiles() carapace.Action {
	return carapace.ActionFiles(".zig", ".zon")
}

func actionZigInputFiles() carapace.Action {
	return carapace.ActionFiles(
		".zig", ".zon", ".o", ".obj", ".lib", ".a", ".so", ".dll", ".dylib",
		".tbd", ".s", ".S", ".c", ".cpp", ".C", ".cc", ".cxx", ".c++",
		".bc", ".def", ".rc", ".res",
	)
}

func actionBuildSteps() carapace.Action {
	return carapace.ActionExecCommand("zig", "build", "--help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		values := make([]string, 0)
		inSteps := false

		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				if inSteps {
					break
				}
				continue
			}

			if strings.HasPrefix(line, "Usage:") || strings.HasPrefix(line, "Options:") {
				continue
			}

			fields := strings.Fields(line)
			if len(fields) == 0 || strings.HasPrefix(fields[0], "-") {
				if inSteps {
					break
				}
				continue
			}

			inSteps = true
			name := fields[0]
			description := strings.TrimSpace(strings.TrimPrefix(line, name))
			values = append(values, name, description)
		}

		return carapace.ActionValuesDescribed(values...)
	})
}

func addCompileFlags(cmd *cobra.Command) {
	cmd.Flags().StringS("cflags", "cflags", "", "Set extra flags for the next positional C source files")
	cmd.Flags().StringP("define", "D", "", "Define C macro")
	cmd.Flags().String("dep", "", "Add an entry to the next module's import table")
	cmd.Flags().String("error-limit", "", "Set the maximum amount of distinct error values")
	cmd.Flags().BoolS("fPIC", "fPIC", false, "Force-enable Position Independent Code")
	cmd.Flags().BoolS("fPIE", "fPIE", false, "Force-enable Position Independent Executable")
	cmd.Flags().BoolS("fasync-unwind-tables", "fasync-unwind-tables", false, "Always produce asynchronous unwind table entries")
	cmd.Flags().BoolS("fbuiltin", "fbuiltin", false, "Enable implicit builtin knowledge of functions")
	cmd.Flags().BoolS("fclang", "fclang", false, "Force using Clang as the C/C++ compilation backend")
	cmd.Flags().BoolS("fdata-sections", "fdata-sections", false, "Place each data in a separate section")
	cmd.Flags().BoolS("fdll-export-fns", "fdll-export-fns", false, "Mark exported functions as DLL exports")
	cmd.Flags().BoolS("ferror-tracing", "ferror-tracing", false, "Enable error tracing in ReleaseFast mode")
	cmd.Flags().BoolS("ffunction-sections", "ffunction-sections", false, "Place each function in a separate section")
	cmd.Flags().BoolS("ffuzz", "ffuzz", false, "Enable fuzz testing instrumentation")
	cmd.Flags().BoolS("flibllvm", "flibllvm", false, "Force using the LLVM API in the codegen backend")
	cmd.Flags().BoolS("fllvm", "fllvm", false, "Force using LLVM as the codegen backend")
	cmd.Flags().BoolS("flto", "flto", false, "Force-enable Link Time Optimization")
	cmd.Flags().BoolS("fno-PIC", "fno-PIC", false, "Force-disable Position Independent Code")
	cmd.Flags().BoolS("fno-PIE", "fno-PIE", false, "Force-disable Position Independent Executable")
	cmd.Flags().BoolS("fno-builtin", "fno-builtin", false, "Disable implicit builtin knowledge of functions")
	cmd.Flags().BoolS("fno-clang", "fno-clang", false, "Prevent using Clang as the C/C++ compilation backend")
	cmd.Flags().BoolS("fno-data-sections", "fno-data-sections", false, "All data go into same section")
	cmd.Flags().BoolS("fno-dll-export-fns", "fno-dll-export-fns", false, "Force-disable marking exported functions as DLL exports")
	cmd.Flags().BoolS("fno-error-tracing", "fno-error-tracing", false, "Disable error tracing")
	cmd.Flags().BoolS("fno-function-sections", "fno-function-sections", false, "All functions go into same section")
	cmd.Flags().BoolS("fno-fuzz", "fno-fuzz", false, "Disable fuzz testing instrumentation")
	cmd.Flags().BoolS("fno-libllvm", "fno-libllvm", false, "Prevent using the LLVM API in the codegen backend")
	cmd.Flags().BoolS("fno-llvm", "fno-llvm", false, "Prevent using LLVM as a codegen backend")
	cmd.Flags().BoolS("fno-lto", "fno-lto", false, "Force-disable Link Time Optimization")
	cmd.Flags().BoolS("fno-omit-frame-pointer", "fno-omit-frame-pointer", false, "Store the stack frame pointer")
	cmd.Flags().BoolS("fno-reference-trace", "fno-reference-trace", false, "Disable reference trace")
	cmd.Flags().BoolS("fno-sanitize-c", "fno-sanitize-c", false, "Disable C undefined behavior detection")
	cmd.Flags().BoolS("fno-sanitize-thread", "fno-sanitize-thread", false, "Disable Thread Sanitizer")
	cmd.Flags().BoolS("fno-single-threaded", "fno-single-threaded", false, "Code may not assume there is only one thread")
	cmd.Flags().BoolS("fno-stack-check", "fno-stack-check", false, "Disable stack probing in safe builds")
	cmd.Flags().BoolS("fno-stack-protector", "fno-stack-protector", false, "Disable stack probing in safe builds")
	cmd.Flags().BoolS("fno-strip", "fno-strip", false, "Keep debug symbols")
	cmd.Flags().BoolS("fno-unwind-tables", "fno-unwind-tables", false, "Never produce unwind table entries")
	cmd.Flags().BoolS("fno-valgrind", "fno-valgrind", false, "Omit valgrind client requests in debug builds")
	cmd.Flags().BoolS("fomit-frame-pointer", "fomit-frame-pointer", false, "Omit the stack frame pointer")
	cmd.Flags().StringS("freference-trace", "freference-trace", "", "How many lines of reference trace should be shown per compile error")
	cmd.Flags().StringS("fsanitize-c", "fsanitize-c", "", "Enable C undefined behavior detection")
	cmd.Flags().BoolS("fsanitize-thread", "fsanitize-thread", false, "Enable Thread Sanitizer")
	cmd.Flags().BoolS("fsingle-threaded", "fsingle-threaded", false, "Code assumes there is only one thread")
	cmd.Flags().BoolS("fstack-check", "fstack-check", false, "Enable stack probing in unsafe builds")
	cmd.Flags().BoolS("fstack-protector", "fstack-protector", false, "Enable stack protection in unsafe builds")
	cmd.Flags().BoolS("fstrip", "fstrip", false, "Omit debug symbols")
	cmd.Flags().BoolS("funwind-tables", "funwind-tables", false, "Always produce unwind table entries")
	cmd.Flags().BoolS("fvalgrind", "fvalgrind", false, "Include valgrind client requests in release builds")
	cmd.Flags().StringS("idirafter", "idirafter", "", "Add directory to AFTER include search path")
	cmd.Flags().StringP("include-dir", "I", "", "Add directory to include search path")
	cmd.Flags().StringS("isystem", "isystem", "", "Add directory to SYSTEM include search path")
	cmd.Flags().StringP("language", "x", "", "Treat subsequent input files as having type")
	cmd.Flags().String("libc", "", "Provide a file which specifies libc paths")
	cmd.Flags().StringS("mcmodel", "mcmodel", "", "Limit range of code and data virtual addresses")
	cmd.Flags().StringS("mcpu", "mcpu", "", "Specify target CPU and feature set")
	cmd.Flags().StringS("mexec-model", "mexec-model", "", "WASI execution model")
	cmd.Flags().BoolS("mno-red-zone", "mno-red-zone", false, "Force-disable the red-zone")
	cmd.Flags().StringP("module", "M", "", "Create a module based on current per-module settings")
	cmd.Flags().BoolS("mred-zone", "mred-zone", false, "Force-enable the red-zone")
	cmd.Flags().BoolS("municode", "municode", false, "Use wmain/wWinMain as entry point")
	cmd.Flags().String("name", "", "Override root name")
	cmd.Flags().StringS("ofmt", "ofmt", "", "Override target object format")
	cmd.Flags().StringP("optimize", "O", "", "Choose what to optimize for")
	cmd.Flags().StringS("rcincludes", "rcincludes", "", "Set the type of includes to use when compiling rc source files")
	cmd.Flags().StringS("target", "target", "", "Target triple")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"idirafter":   carapace.ActionDirectories(),
		"include-dir": carapace.ActionDirectories(),
		"isystem":     carapace.ActionDirectories(),
		"language":    carapace.ActionValues("c", "c-header", "cpp", "cpp-header", "zig", "zig-object", "object", "assembly", "assembly-with-cpp", "ll", "bc"),
		"libc":        carapace.ActionFiles(),
		"mcmodel":     carapace.ActionValues("default", "extreme", "kernel", "large", "medany", "medium", "medlow", "medmid", "normal", "small", "tiny").StyleF(style.ForKeyword),
		"mexec-model": carapace.ActionValues("command", "reactor").StyleF(style.ForKeyword),
		"ofmt":        actionObjectFormats(),
		"optimize":    actionOptimizationModes(),
		"rcincludes":  carapace.ActionValues("any", "msvc", "gnu", "none").StyleF(style.ForKeyword),
		"fsanitize-c": carapace.ActionValues("trap", "full").StyleF(style.ForKeyword),
	})
}

func addLinkFlags(cmd *cobra.Command) {
	cmd.Flags().BoolS("Bsymbolic", "Bsymbolic", false, "Bind global references locally")
	cmd.Flags().BoolS("ObjC", "ObjC", false, "Darwin force load Objective-C members")
	cmd.Flags().String("build-id", "", "Embed a build ID in binaries")
	cmd.Flags().String("compress-debug-sections", "", "Debug section compression setting")
	cmd.Flags().BoolS("dead_strip", "dead_strip", false, "Darwin remove unreachable functions and data")
	cmd.Flags().BoolS("dead_strip_dylibs", "dead_strip_dylibs", false, "Darwin remove unreachable dylibs")
	cmd.Flags().Bool("disable-new-dtags", false, "Use RPATH behavior for dynamic tags")
	cmd.Flags().BoolS("dynamic", "dynamic", false, "Force output to be dynamically linked")
	cmd.Flags().String("dynamic-linker", "", "Set the dynamic interpreter path")
	cmd.Flags().Bool("eh-frame-hdr", false, "Enable C++ exception handling")
	cmd.Flags().Bool("emit-relocs", false, "Enable output of relocation sections")
	cmd.Flags().Bool("enable-new-dtags", false, "Use RUNPATH behavior for dynamic tags")
	cmd.Flags().String("entitlements", "", "Darwin entitlements file")
	cmd.Flags().String("export", "", "WebAssembly force a symbol to be exported")
	cmd.Flags().Bool("export-memory", false, "WebAssembly export memory to host")
	cmd.Flags().Bool("export-table", false, "WebAssembly export function table")
	cmd.Flags().BoolS("fallow-shlib-undefined", "fallow-shlib-undefined", false, "Allow undefined symbols in shared libraries")
	cmd.Flags().BoolS("fallow-so-scripts", "fallow-so-scripts", false, "Allow .so files to be GNU ld scripts")
	cmd.Flags().BoolS("fcompiler-rt", "fcompiler-rt", false, "Always include compiler-rt symbols in output")
	cmd.Flags().BoolS("feach-lib-rpath", "feach-lib-rpath", false, "Add rpath for each used dynamic library")
	cmd.Flags().StringS("fentry", "fentry", "", "Set the entrypoint symbol name")
	cmd.Flags().BoolS("flld", "flld", false, "Force using LLD as the linker")
	cmd.Flags().BoolS("fno-allow-shlib-undefined", "fno-allow-shlib-undefined", false, "Disallow undefined symbols in shared libraries")
	cmd.Flags().BoolS("fno-allow-so-scripts", "fno-allow-so-scripts", false, ".so files must be ELF files")
	cmd.Flags().BoolS("fno-compiler-rt", "fno-compiler-rt", false, "Prevent including compiler-rt symbols in output")
	cmd.Flags().BoolS("fno-each-lib-rpath", "fno-each-lib-rpath", false, "Prevent adding rpath for each used dynamic library")
	cmd.Flags().BoolS("fno-entry", "fno-entry", false, "Do not output any entry point")
	cmd.Flags().BoolS("fno-lld", "fno-lld", false, "Prevent using LLD as the linker")
	cmd.Flags().BoolS("fno-soname", "fno-soname", false, "Disable emitting a SONAME")
	cmd.Flags().BoolS("fno-ubsan-rt", "fno-ubsan-rt", false, "Prevent including ubsan-rt symbols in output")
	cmd.Flags().String("force_undefined", "", "Specify the symbol that must be defined")
	cmd.Flags().StringS("framework", "framework", "", "Darwin link against framework")
	cmd.Flags().StringP("framework-directory", "F", "", "Darwin add search path for frameworks")
	cmd.Flags().StringS("fsoname", "fsoname", "", "Override the default SONAME value")
	cmd.Flags().BoolS("fubsan-rt", "fubsan-rt", false, "Always include ubsan-rt symbols in output")
	cmd.Flags().Bool("gc-sections", false, "Remove unreachable functions and data")
	cmd.Flags().String("global-base", "", "WebAssembly global data base")
	cmd.Flags().StringS("headerpad", "headerpad", "", "Darwin load command expansion space")
	cmd.Flags().BoolS("headerpad_max_install_names", "headerpad_max_install_names", false, "Darwin MAXPATHLEN header padding")
	cmd.Flags().String("image-base", "", "Set base address for executable image")
	cmd.Flags().Bool("import-memory", false, "WebAssembly import memory from environment")
	cmd.Flags().Bool("import-symbols", false, "WebAssembly import missing symbols")
	cmd.Flags().Bool("import-table", false, "WebAssembly import function table")
	cmd.Flags().String("initial-memory", "", "WebAssembly initial memory size")
	cmd.Flags().StringS("install_name", "install_name", "", "Darwin dylib install name")
	cmd.Flags().StringP("library", "l", "", "Link against system library")
	cmd.Flags().StringP("library-directory", "L", "", "Add a directory to the library search path")
	cmd.Flags().StringP("linker-extension", "z", "", "Set linker extension flags")
	cmd.Flags().String("max-memory", "", "WebAssembly maximum memory size")
	cmd.Flags().String("needed-library", "", "Link against system library even if unused")
	cmd.Flags().StringS("needed_framework", "needed_framework", "", "Darwin link against framework even if unused")
	cmd.Flags().Bool("no-eh-frame-hdr", false, "Disable C++ exception handling")
	cmd.Flags().Bool("no-gc-sections", false, "Do not force removal of unreachable functions and data")
	cmd.Flags().Bool("no-undefined-version", false, "Disallow version scripts from referring to undefined symbols")
	cmd.Flags().StringS("pagezero_size", "pagezero_size", "", "Darwin __PAGEZERO segment size")
	cmd.Flags().BoolS("rdynamic", "rdynamic", false, "Add all symbols to the dynamic symbol table")
	cmd.Flags().StringS("rpath", "rpath", "", "Add directory to the runtime library search path")
	cmd.Flags().StringP("script", "T", "", "Use a custom linker script")
	cmd.Flags().Bool("shared-memory", false, "WebAssembly use shared linear memory")
	cmd.Flags().String("sort-section", "", "Sort wildcard section patterns")
	cmd.Flags().String("stack", "", "Override default stack size")
	cmd.Flags().BoolS("static", "static", false, "Force output to be statically linked")
	cmd.Flags().String("subsystem", "", "Windows subsystem")
	cmd.Flags().String("sysroot", "", "Set the system root directory")
	cmd.Flags().Bool("undefined-version", false, "Allow version scripts to refer to undefined symbols")
	cmd.Flags().String("version", "", "Dynamic library semver")
	cmd.Flags().String("version-script", "", "Provide a version map file")
	cmd.Flags().StringS("weak_framework", "weak_framework", "", "Darwin link against framework and mark weak")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"build-id":                actionBuildIDs(),
		"compress-debug-sections": carapace.ActionValues("none", "zlib", "zstd").StyleF(style.ForKeyword),
		"dynamic-linker":          carapace.ActionFiles(),
		"entitlements":            carapace.ActionFiles(),
		"framework-directory":     carapace.ActionDirectories(),
		"library-directory":       carapace.ActionDirectories(),
		"rpath":                   carapace.ActionDirectories(),
		"script":                  carapace.ActionFiles(),
		"sysroot":                 carapace.ActionDirectories(),
		"version-script":          carapace.ActionFiles(),
	})
}

func addDebugFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("debug-compile-errors", false, "Crash with diagnostics at the first compile error")
	cmd.Flags().Bool("debug-link-snapshot", false, "Dump linker state in JSON format")
	cmd.Flags().String("debug-log", "", "Enable printing debug/info log messages for scope")
	cmd.Flags().Bool("debug-rt", false, "Debug compiler runtime libraries")
	cmd.Flags().StringS("fopt-bisect-limit", "fopt-bisect-limit", "", "Only run first LLVM optimization passes")
	cmd.Flags().BoolS("fstack-report", "fstack-report", false, "Print stack size diagnostics")
	cmd.Flags().BoolS("ftime-report", "ftime-report", false, "Print timing diagnostics")
	cmd.Flags().Bool("verbose-air", false, "Enable compiler debug output for Zig AIR")
	cmd.Flags().Bool("verbose-cc", false, "Display C compiler invocations")
	cmd.Flags().Bool("verbose-cimport", false, "Enable compiler debug output for C imports")
	cmd.Flags().Bool("verbose-generic-instances", false, "Enable compiler debug output for generic instance generation")
	cmd.Flags().Bool("verbose-intern-pool", false, "Enable compiler debug output for InternPool")
	cmd.Flags().Bool("verbose-link", false, "Display linker invocations")
	cmd.Flags().String("verbose-llvm-bc", "", "Enable compiler debug output for unoptimized LLVM BC")
	cmd.Flags().Bool("verbose-llvm-cpu-features", false, "Enable compiler debug output for LLVM CPU features")
	cmd.Flags().Bool("verbose-llvm-ir", false, "Enable compiler debug output for LLVM IR")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"verbose-llvm-bc": carapace.ActionFiles(".bc"),
	})
}

func addTestFlags(cmd *cobra.Command) {
	cmd.Flags().String("test-cmd", "", "Specify test execution command one arg at a time")
	cmd.Flags().Bool("test-cmd-bin", false, "Append test binary path to test command arguments")
	cmd.Flags().Bool("test-evented-io", false, "Run the test in evented I/O mode")
	cmd.Flags().String("test-filter", "", "Skip tests that do not match filter")
	cmd.Flags().String("test-name-prefix", "", "Add prefix to all tests")
	cmd.Flags().Bool("test-no-exec", false, "Compile test binary without running it")
	cmd.Flags().String("test-runner", "", "Specify a custom test runner")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"test-runner": carapace.ActionFiles(),
	})
}

func addCompileCompletions(cmd *cobra.Command) {
	carapace.Gen(cmd).PositionalAnyCompletion(
		actionZigInputFiles(),
	)
}
