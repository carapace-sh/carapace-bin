package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zig",
	Short: "Programming language designed for robustness, optimality, and clarity",
	Long:  "https://ziglang.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddCommand(
		buildCmd(),
		simpleCmd("fetch", "Copy a package into global cache and print its hash"),
		simpleCmd("init", "Initialize a Zig package in the current directory"),
		compileCmd("build-exe", "Create executable from source or object files"),
		compileCmd("build-lib", "Create library from source or object files"),
		compileCmd("build-obj", "Create object from source or object files"),
		testCmd(),
		compileCmd("run", "Create executable and run immediately"),
		compileCmd("ast-check", "Look for simple compile errors in any set of files"),
		fmtCmd(),
		compileCmd("reduce", "Minimize a bug report"),
		compileCmd("translate-c", "Convert C code to Zig code"),
		simpleCmd("ar", "Use Zig as a drop-in archiver"),
		simpleCmd("cc", "Use Zig as a drop-in C compiler"),
		simpleCmd("c++", "Use Zig as a drop-in C++ compiler"),
		simpleCmd("dlltool", "Use Zig as a drop-in dlltool.exe"),
		simpleCmd("lib", "Use Zig as a drop-in lib.exe"),
		simpleCmd("ranlib", "Use Zig as a drop-in ranlib"),
		simpleCmd("objcopy", "Use Zig as a drop-in objcopy"),
		simpleCmd("rc", "Use Zig as a drop-in rc.exe"),
		simpleCmd("env", "Print lib path, std path, cache directory, and version"),
		simpleCmd("help", "Print this help and exit"),
		simpleCmd("std", "View standard library documentation in a browser"),
		simpleCmd("libc", "Display native libc paths file or validate one"),
		simpleCmd("targets", "List available compilation targets"),
		simpleCmd("version", "Print version number and exit"),
		simpleCmd("zen", "Print Zen of Zig and exit"),
	)

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"build", "Build project from build.zig",
			"fetch", "Copy a package into global cache and print its hash",
			"init", "Initialize a Zig package in the current directory",
			"build-exe", "Create executable from source or object files",
			"build-lib", "Create library from source or object files",
			"build-obj", "Create object from source or object files",
			"test", "Perform unit testing",
			"run", "Create executable and run immediately",
			"ast-check", "Look for simple compile errors in any set of files",
			"fmt", "Reformat Zig source into canonical form",
			"reduce", "Minimize a bug report",
			"translate-c", "Convert C code to Zig code",
			"ar", "Use Zig as a drop-in archiver",
			"cc", "Use Zig as a drop-in C compiler",
			"c++", "Use Zig as a drop-in C++ compiler",
			"dlltool", "Use Zig as a drop-in dlltool.exe",
			"lib", "Use Zig as a drop-in lib.exe",
			"ranlib", "Use Zig as a drop-in ranlib",
			"objcopy", "Use Zig as a drop-in objcopy",
			"rc", "Use Zig as a drop-in rc.exe",
			"env", "Print lib path, std path, cache directory, and version",
			"help", "Print this help and exit",
			"std", "View standard library documentation in a browser",
			"libc", "Display native libc paths file or validate one",
			"targets", "List available compilation targets",
			"version", "Print version number and exit",
			"zen", "Print Zen of Zig and exit",
		),
	)
}

func simpleCmd(use, short string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(cmd).Standalone()
	carapace.Gen(cmd).PositionalAnyCompletion(carapace.ActionFiles())
	return cmd
}

func compileCmd(use, short string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(cmd).Standalone()
	addCompileFlags(cmd)
	carapace.Gen(cmd).PositionalAnyCompletion(carapace.ActionFiles(".zig", ".c", ".cpp", ".h", ".o", ".a"))
	return cmd
}

func testCmd() *cobra.Command {
	cmd := compileCmd("test", "Perform unit testing")

	cmd.Flags().StringArray("test-cmd", nil, "Specify test execution command")
	cmd.Flags().Bool("test-cmd-bin", false, "Appends test binary path to test command args")
	cmd.Flags().String("test-filter", "", "Skip tests that do not match filter")
	cmd.Flags().String("test-name-prefix", "", "Add prefix to all tests")
	cmd.Flags().Bool("test-no-exec", false, "Compiles test binary without running it")

	return cmd
}

func buildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build project from build.zig",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(cmd).Standalone()
	cmd.Flags().String("build-file", "", "Override path to build.zig")
	cmd.Flags().String("build-runner", "", "Override path to build runner")
	cmd.Flags().String("debounce", "", "Delay before rebuilding after changed file detected")
	cmd.Flags().StringP("prefix", "p", "", "Where to install files")
	cmd.Flags().String("prefix-exe-dir", "", "Where to install executables")
	cmd.Flags().String("prefix-include-dir", "", "Where to install headers")
	cmd.Flags().String("prefix-lib-dir", "", "Where to install libraries")
	cmd.Flags().String("release", "", "Select optimizations")
	cmd.Flags().String("summary", "", "Control the printing of the build summary")
	cmd.Flags().Bool("watch", false, "Continuously rebuild when source files are modified")
	addGlobalFlags(cmd)

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"build-file":         carapace.ActionFiles(".zig"),
		"build-runner":       carapace.ActionFiles(),
		"cache-dir":          carapace.ActionDirectories(),
		"color":              actionColors(),
		"global-cache-dir":   carapace.ActionDirectories(),
		"prefix":             carapace.ActionDirectories(),
		"prefix-exe-dir":     carapace.ActionDirectories(),
		"prefix-include-dir": carapace.ActionDirectories(),
		"prefix-lib-dir":     carapace.ActionDirectories(),
		"release":            actionOptimizationModes(),
		"summary":            carapace.ActionValues("all", "new", "failures", "none"),
		"zig-lib-dir":        carapace.ActionDirectories(),
	})

	carapace.Gen(cmd).PositionalAnyCompletion(carapace.ActionFiles())
	return cmd
}

func fmtCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fmt",
		Short: "Reformat Zig source into canonical form",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(cmd).Standalone()
	cmd.Flags().Bool("ast-check", false, "Run zig ast-check on every file")
	cmd.Flags().Bool("check", false, "List non-conforming files and exit with an error if the list is non-empty")
	cmd.Flags().String("color", "", "Enable or disable colored error messages")
	cmd.Flags().StringArray("exclude", nil, "Exclude file or directory from formatting")
	cmd.Flags().Bool("stdin", false, "Format code from stdin and output to stdout")
	cmd.Flags().Bool("zon", false, "Treat all input files as ZON, regardless of file extension")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"color":   actionColors(),
		"exclude": carapace.ActionFiles(),
	})

	carapace.Gen(cmd).PositionalAnyCompletion(carapace.ActionFiles(".zig", ".zon"))
	return cmd
}

func addGlobalFlags(cmd *cobra.Command) {
	cmd.Flags().String("cache-dir", "", "Override the local cache directory")
	cmd.Flags().String("color", "", "Enable or disable colored error messages")
	cmd.Flags().StringS("femit-asm", "femit-asm", "", "Output assembly code")
	cmd.Flags().StringS("femit-bin", "femit-bin", "", "Output machine code")
	cmd.Flags().StringS("femit-docs", "femit-docs", "", "Create a docs directory with HTML documentation")
	cmd.Flags().StringS("femit-h", "femit-h", "", "Generate a C header file")
	cmd.Flags().StringS("femit-llvm-bc", "femit-llvm-bc", "", "Produce LLVM bitcode")
	cmd.Flags().StringS("femit-llvm-ir", "femit-llvm-ir", "", "Produce LLVM IR")
	cmd.Flags().BoolS("fincremental", "fincremental", false, "Enable incremental compilation")
	cmd.Flags().BoolS("fno-emit-asm", "fno-emit-asm", false, "Do not output assembly code")
	cmd.Flags().BoolS("fno-emit-bin", "fno-emit-bin", false, "Do not output machine code")
	cmd.Flags().BoolS("fno-emit-docs", "fno-emit-docs", false, "Do not produce docs")
	cmd.Flags().BoolS("fno-emit-h", "fno-emit-h", false, "Do not generate a C header file")
	cmd.Flags().BoolS("fno-emit-llvm-bc", "fno-emit-llvm-bc", false, "Do not produce LLVM bitcode")
	cmd.Flags().BoolS("fno-emit-llvm-ir", "fno-emit-llvm-ir", false, "Do not produce LLVM IR")
	cmd.Flags().BoolS("fno-incremental", "fno-incremental", false, "Disable incremental compilation")
	cmd.Flags().String("global-cache-dir", "", "Override the global cache directory")
	cmd.Flags().StringS("j", "j", "", "Limit concurrent jobs")
	cmd.Flags().String("zig-lib-dir", "", "Override path to Zig installation lib directory")
}

func addCompileFlags(cmd *cobra.Command) {
	addGlobalFlags(cmd)

	cmd.Flags().StringS("I", "I", "", "Add directory to include search path")
	cmd.Flags().StringS("L", "L", "", "Add directory to library search path")
	cmd.Flags().StringS("M", "M", "", "Create a module based on current per-module settings")
	cmd.Flags().StringS("O", "O", "", "Choose what to optimize for")
	cmd.Flags().StringS("T", "T", "", "Use linker script")
	cmd.Flags().String("dep", "", "Add an entry to the next module's import table")
	cmd.Flags().String("dynamic-linker", "", "Set dynamic linker path")
	cmd.Flags().String("error-limit", "", "Set the maximum amount of distinct error values")
	cmd.Flags().BoolS("fPIC", "fPIC", false, "Force-enable Position Independent Code")
	cmd.Flags().BoolS("fPIE", "fPIE", false, "Force-enable Position Independent Executable")
	cmd.Flags().BoolS("fclang", "fclang", false, "Force using Clang as the C/C++ compilation backend")
	cmd.Flags().BoolS("fllvm", "fllvm", false, "Force using LLVM as the codegen backend")
	cmd.Flags().BoolS("flto", "flto", false, "Force-enable Link Time Optimization")
	cmd.Flags().BoolS("fno-PIC", "fno-PIC", false, "Force-disable Position Independent Code")
	cmd.Flags().BoolS("fno-PIE", "fno-PIE", false, "Force-disable Position Independent Executable")
	cmd.Flags().BoolS("fno-clang", "fno-clang", false, "Prevent using Clang as the C/C++ compilation backend")
	cmd.Flags().BoolS("fno-llvm", "fno-llvm", false, "Prevent using LLVM as a codegen backend")
	cmd.Flags().BoolS("fno-lto", "fno-lto", false, "Force-disable Link Time Optimization")
	cmd.Flags().StringS("l", "l", "", "Link against system library")
	cmd.Flags().String("libc", "", "Provide a file which specifies libc paths")
	cmd.Flags().StringS("mcmodel", "mcmodel", "", "Limit range of code and data virtual addresses")
	cmd.Flags().StringS("mcpu", "mcpu", "", "Specify target CPU and feature set")
	cmd.Flags().String("name", "", "Override root name")
	cmd.Flags().StringS("ofmt", "ofmt", "", "Override target object format")
	cmd.Flags().BoolS("rdynamic", "rdynamic", false, "Add all symbols to the dynamic symbol table")
	cmd.Flags().BoolS("strip", "strip", false, "Omit debug symbols")
	cmd.Flags().String("sysroot", "", "Set the system root directory")
	cmd.Flags().StringS("target", "target", "", "Target triple")
	cmd.Flags().StringS("x", "x", "", "Treat subsequent input files as having type")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"cache-dir":        carapace.ActionDirectories(),
		"color":            actionColors(),
		"dynamic-linker":   carapace.ActionFiles(),
		"femit-asm":        carapace.ActionFiles(".s"),
		"femit-bin":        carapace.ActionFiles(),
		"femit-docs":       carapace.ActionDirectories(),
		"femit-h":          carapace.ActionFiles(".h"),
		"femit-llvm-bc":    carapace.ActionFiles(".bc"),
		"femit-llvm-ir":    carapace.ActionFiles(".ll"),
		"global-cache-dir": carapace.ActionDirectories(),
		"I":                carapace.ActionDirectories(),
		"L":                carapace.ActionDirectories(),
		"libc":             carapace.ActionFiles(),
		"mcmodel":          carapace.ActionValues("default", "extreme", "kernel", "large", "medany", "medium", "medlow", "medmid", "normal", "small", "tiny"),
		"O":                actionOptimizationModes(),
		"ofmt":             carapace.ActionValues("elf", "c", "wasm", "coff", "macho", "spirv", "plan9", "hex", "raw"),
		"sysroot":          carapace.ActionDirectories(),
		"T":                carapace.ActionFiles(),
		"x":                carapace.ActionValues("c", "c-header", "cpp", "assembly", "object", "zig"),
		"zig-lib-dir":      carapace.ActionDirectories(),
	})
}

func actionColors() carapace.Action {
	return carapace.ActionValues("on", "off", "auto")
}

func actionOptimizationModes() carapace.Action {
	return carapace.ActionValues("Debug", "ReleaseFast", "ReleaseSafe", "ReleaseSmall")
}
