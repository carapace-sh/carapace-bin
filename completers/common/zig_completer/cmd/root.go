package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
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
	rootCmd.Flags().BoolP("help", "h", false, "Print command-specific usage")

	addCommands()
}

func addCommands() {
	buildCmd := addCommand(rootCmd, "build [steps]", "Build project from build.zig")
	buildCmd.Flags().String("prefix", "", "Installation prefix")
	buildCmd.Flags().String("prefix-lib-dir", "", "Path to installation lib directory")
	buildCmd.Flags().String("prefix-exe-dir", "", "Path to installation executable directory")
	buildCmd.Flags().String("prefix-include-dir", "", "Path to installation include directory")
	buildCmd.Flags().String("release", "", "Request release mode")
	buildCmd.Flags().String("color", "", "Enable or disable colored error messages")
	buildCmd.Flags().StringS("j", "j", "", "Limit concurrent jobs")
	buildCmd.Flags().StringArrayS("D", "D", nil, "Set project option")
	buildCmd.Flags().String("build-file", "", "Override path to build.zig")
	buildCmd.Flags().String("cache-dir", "", "Override the local cache directory")
	buildCmd.Flags().String("global-cache-dir", "", "Override the global cache directory")
	buildCmd.Flags().String("zig-lib-dir", "", "Override path to Zig installation lib directory")
	buildCmd.Flags().String("summary", "", "Control build summary output")
	buildCmd.Flags().Bool("watch", false, "Rebuild when inputs change")
	buildCmd.Flags().Bool("fuzz", false, "Run fuzz tests")
	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"build-file":         carapace.ActionFiles(".zig"),
		"cache-dir":          carapace.ActionDirectories(),
		"color":              actionColor(),
		"global-cache-dir":   carapace.ActionDirectories(),
		"prefix":             carapace.ActionDirectories(),
		"prefix-exe-dir":     carapace.ActionDirectories(),
		"prefix-include-dir": carapace.ActionDirectories(),
		"prefix-lib-dir":     carapace.ActionDirectories(),
		"release":            carapace.ActionValues("fast", "safe", "small"),
		"summary":            carapace.ActionValues("all", "new", "failures", "none"),
		"zig-lib-dir":        carapace.ActionDirectories(),
	})

	addCommand(rootCmd, "fetch [url]", "Copy a package into global cache and print its hash")
	addCommand(rootCmd, "init", "Initialize a Zig package in the current directory")

	for _, spec := range []struct {
		use         string
		short       string
		sourceTypes []string
	}{
		{"build-exe [files]", "Create executable from source or object files", sourceFileTypes()},
		{"build-lib [files]", "Create library from source or object files", sourceFileTypes()},
		{"build-obj [files]", "Create object from source or object files", sourceFileTypes()},
		{"test [files]", "Perform unit testing", sourceFileTypes()},
		{"test-obj [files]", "Create object for unit testing", sourceFileTypes()},
		{"run [files]", "Create executable and run immediately", sourceFileTypes()},
		{"translate-c [file]", "Convert C code to Zig code", []string{".c", ".h"}},
		{"ast-check [files]", "Look for simple compile errors in any set of files", []string{".zig"}},
	} {
		addCompileCommand(spec.use, spec.short, spec.sourceTypes)
	}

	fmtCmd := addCommand(rootCmd, "fmt [files]", "Reformat Zig source into canonical form")
	fmtCmd.Flags().Bool("check", false, "List non-conforming files and exit with an error if any are found")
	fmtCmd.Flags().Bool("ast-check", false, "Run zig ast-check on every file")
	fmtCmd.Flags().StringArray("exclude", nil, "Exclude file or directory from formatting")
	fmtCmd.Flags().String("color", "", "Enable or disable colored error messages")
	carapace.Gen(fmtCmd).FlagCompletion(carapace.ActionMap{
		"color":   actionColor(),
		"exclude": carapace.ActionFiles(),
	})
	carapace.Gen(fmtCmd).PositionalAnyCompletion(carapace.ActionFiles(".zig").FilterArgs())

	reduceCmd := addCommand(rootCmd, "reduce [files]", "Minimize a bug report")
	reduceCmd.Flags().String("main-mod-path", "", "Set the directory of the root module")
	reduceCmd.Flags().String("mod", "", "Set a module dependency")
	carapace.Gen(reduceCmd).FlagCompletion(carapace.ActionMap{
		"main-mod-path": carapace.ActionDirectories(),
	})
	carapace.Gen(reduceCmd).PositionalAnyCompletion(carapace.ActionFiles(".zig").FilterArgs())

	translateDropIn(rootCmd, "ar", "Use Zig as a drop-in archiver", "ar")
	translateDropIn(rootCmd, "cc", "Use Zig as a drop-in C compiler", "cc")
	translateDropIn(rootCmd, "c++", "Use Zig as a drop-in C++ compiler", "c++")
	translateDropIn(rootCmd, "dlltool", "Use Zig as a drop-in dlltool.exe", "dlltool")
	translateDropIn(rootCmd, "lib", "Use Zig as a drop-in lib.exe", "lib")
	translateDropIn(rootCmd, "ranlib", "Use Zig as a drop-in ranlib", "ranlib")
	translateDropIn(rootCmd, "objcopy", "Use Zig as a drop-in objcopy", "objcopy")
	translateDropIn(rootCmd, "rc", "Use Zig as a drop-in rc.exe", "rc")

	addCommand(rootCmd, "env", "Print lib path, std path, cache directory, and version")

	helpCmd := addCommand(rootCmd, "help [command]", "Print help and exit")
	carapace.Gen(helpCmd).PositionalCompletion(actionCommands())

	addCommand(rootCmd, "std", "View standard library documentation in a browser")

	libcCmd := addCommand(rootCmd, "libc [file]", "Display native libc paths file or validate one")
	carapace.Gen(libcCmd).PositionalCompletion(carapace.ActionFiles())

	addCommand(rootCmd, "targets", "List available compilation targets")
	addCommand(rootCmd, "version", "Print version number and exit")
	addCommand(rootCmd, "zen", "Print Zen of Zig and exit")
}

func addCommand(parent *cobra.Command, use string, short string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	carapace.Gen(cmd).Standalone()
	parent.AddCommand(cmd)
	return cmd
}

func addCompileCommand(use string, short string, sourceTypes []string) {
	cmd := addCommand(rootCmd, use, short)
	addGeneralCompileFlags(cmd)
	addModuleCompileFlags(cmd)
	addLinkFlags(cmd)

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"cache-dir":               carapace.ActionDirectories(),
		"color":                   actionColor(),
		"femit-asm":               carapace.ActionFiles(".s"),
		"femit-bin":               carapace.ActionFiles(),
		"femit-docs":              carapace.ActionDirectories(),
		"femit-h":                 carapace.ActionFiles(".h"),
		"femit-implib":            carapace.ActionFiles(".lib"),
		"femit-llvm-bc":           carapace.ActionFiles(".bc"),
		"femit-llvm-ir":           carapace.ActionFiles(".ll"),
		"global-cache-dir":        carapace.ActionDirectories(),
		"libc":                    carapace.ActionFiles(),
		"mexec-model":             carapace.ActionValues("command", "reactor"),
		"mcmodel":                 carapace.ActionValues("default", "extreme", "kernel", "large", "medany", "medium", "medlow", "medmid", "normal", "small", "tiny"),
		"ofmt":                    actionObjectFormats(),
		"O":                       actionOptimizationModes(),
		"target":                  actionTargets(),
		"version-script":          carapace.ActionFiles(),
		"zig-lib-dir":             carapace.ActionDirectories(),
		"dynamic-linker":          carapace.ActionFiles(),
		"sysroot":                 carapace.ActionDirectories(),
		"subsystem":               carapace.ActionValues("console", "windows", "posix", "native", "efi_application", "efi_boot_service_driver", "efi_rom", "efi_runtime_driver"),
		"compress-debug-sections": carapace.ActionValues("none", "zlib", "zstd"),
	})

	carapace.Gen(cmd).PositionalAnyCompletion(carapace.ActionFiles(sourceTypes...).FilterArgs())
}

func addGeneralCompileFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("help", "h", false, "Print help and exit")
	cmd.Flags().String("color", "", "Enable or disable colored error messages")
	cmd.Flags().StringS("j", "j", "", "Limit concurrent jobs")
	cmd.Flags().BoolS("fincremental", "fincremental", false, "Enable incremental compilation")
	cmd.Flags().BoolS("fno-incremental", "fno-incremental", false, "Disable incremental compilation")
	cmd.Flags().StringS("femit-bin", "femit-bin", "", "Output machine code")
	cmd.Flags().BoolS("fno-emit-bin", "fno-emit-bin", false, "Do not output machine code")
	cmd.Flags().StringS("femit-asm", "femit-asm", "", "Output assembly code")
	cmd.Flags().BoolS("fno-emit-asm", "fno-emit-asm", false, "Do not output assembly code")
	cmd.Flags().StringS("femit-llvm-ir", "femit-llvm-ir", "", "Produce an LLVM IR file")
	cmd.Flags().BoolS("fno-emit-llvm-ir", "fno-emit-llvm-ir", false, "Do not produce an LLVM IR file")
	cmd.Flags().StringS("femit-llvm-bc", "femit-llvm-bc", "", "Produce an LLVM bitcode file")
	cmd.Flags().BoolS("fno-emit-llvm-bc", "fno-emit-llvm-bc", false, "Do not produce an LLVM bitcode file")
	cmd.Flags().StringS("femit-h", "femit-h", "", "Generate a C header file")
	cmd.Flags().BoolS("fno-emit-h", "fno-emit-h", false, "Do not generate a C header file")
	cmd.Flags().StringS("femit-docs", "femit-docs", "", "Create a docs directory with HTML documentation")
	cmd.Flags().BoolS("fno-emit-docs", "fno-emit-docs", false, "Do not produce docs")
	cmd.Flags().StringS("femit-implib", "femit-implib", "", "Produce an import library")
	cmd.Flags().BoolS("fno-emit-implib", "fno-emit-implib", false, "Do not produce an import library")
	cmd.Flags().Bool("show-builtin", false, "Output the source of @import(\"builtin\") then exit")
	cmd.Flags().String("cache-dir", "", "Override the local cache directory")
	cmd.Flags().String("global-cache-dir", "", "Override the global cache directory")
	cmd.Flags().String("zig-lib-dir", "", "Override path to Zig installation lib directory")
}

func addModuleCompileFlags(cmd *cobra.Command) {
	cmd.Flags().String("name", "", "Compilation unit name")
	cmd.Flags().String("libc", "", "Provide a file which specifies libc paths")
	cmd.Flags().StringS("x", "x", "", "Treat subsequent input files as having this language")
	cmd.Flags().StringArray("dep", nil, "Add an entry to the next module's import table")
	cmd.Flags().StringArrayS("M", "M", nil, "Create a module based on current per-module settings")
	cmd.Flags().String("error-limit", "", "Set the maximum amount of distinct error values")
	cmd.Flags().BoolS("fllvm", "fllvm", false, "Force using LLVM as the codegen backend")
	cmd.Flags().BoolS("fno-llvm", "fno-llvm", false, "Prevent using LLVM as the codegen backend")
	cmd.Flags().BoolS("fclang", "fclang", false, "Force using Clang as the C/C++ compilation backend")
	cmd.Flags().BoolS("fno-clang", "fno-clang", false, "Prevent using Clang as the C/C++ compilation backend")
	cmd.Flags().BoolS("flto", "flto", false, "Force-enable Link Time Optimization")
	cmd.Flags().BoolS("fno-lto", "fno-lto", false, "Force-disable Link Time Optimization")
	cmd.Flags().StringS("target", "target", "", "Target triple")
	cmd.Flags().StringS("O", "O", "", "Choose what to optimize for")
	cmd.Flags().StringS("ofmt", "ofmt", "", "Override target object format")
	cmd.Flags().StringS("mcpu", "mcpu", "", "Specify target CPU and feature set")
	cmd.Flags().StringS("mcmodel", "mcmodel", "", "Limit range of code and data virtual addresses")
	cmd.Flags().StringS("mexec-model", "mexec-model", "", "WASI execution model")
	cmd.Flags().BoolS("municode", "municode", false, "Use wmain/wWinMain as entry point")
	cmd.Flags().StringArrayS("I", "I", nil, "Add directory to include search path")
	cmd.Flags().StringArrayS("D", "D", nil, "Define C macro")
	cmd.Flags().StringArrayS("cflags", "cflags", nil, "Set extra flags for the next positional C source files")
}

func addLinkFlags(cmd *cobra.Command) {
	cmd.Flags().StringS("T", "T", "", "Use a custom linker script")
	cmd.Flags().String("version-script", "", "Provide a version map file")
	cmd.Flags().String("dynamic-linker", "", "Set the dynamic interpreter path")
	cmd.Flags().String("sysroot", "", "Set the system root directory")
	cmd.Flags().String("version", "", "Dynamic library semver")
	cmd.Flags().StringS("fentry", "fentry", "", "Set the entrypoint symbol name")
	cmd.Flags().BoolS("fno-entry", "fno-entry", false, "Do not output any entry point")
	cmd.Flags().BoolS("flld", "flld", false, "Force using LLD as the linker")
	cmd.Flags().BoolS("fno-lld", "fno-lld", false, "Prevent using LLD as the linker")
	cmd.Flags().BoolS("rdynamic", "rdynamic", false, "Add all symbols to the dynamic symbol table")
	cmd.Flags().StringS("z", "z", "", "Set linker extension flags")
	cmd.Flags().BoolS("dynamic", "dynamic", false, "Force output to be dynamically linked")
	cmd.Flags().BoolS("static", "static", false, "Force output to be statically linked")
	cmd.Flags().String("compress-debug-sections", "", "Debug section compression settings")
	cmd.Flags().Bool("gc-sections", false, "Force removal of unreachable functions and data")
	cmd.Flags().Bool("no-gc-sections", false, "Do not remove unreachable functions and data")
	cmd.Flags().String("subsystem", "", "Windows subsystem")
	cmd.Flags().StringArrayS("l", "l", nil, "Link against system library")
	cmd.Flags().StringArray("library", nil, "Link against system library")
	cmd.Flags().StringArrayS("needed-l", "needed-l", nil, "Link against needed system library")
	cmd.Flags().StringArray("needed-library", nil, "Link against needed system library")
	cmd.Flags().StringArrayS("weak-l", "weak-l", nil, "Weak-link against system library")
	cmd.Flags().StringArray("weak-library", nil, "Weak-link against system library")
}

func translateDropIn(parent *cobra.Command, use string, short string, bridgeName string) {
	cmd := addCommand(parent, use, short)
	carapace.Gen(cmd).PositionalAnyCompletion(bridge.ActionCarapaceBin(bridgeName).Split())
}

func actionCommands() carapace.Action {
	return carapace.ActionValuesDescribed(
		"build", "Build project from build.zig",
		"fetch", "Copy a package into global cache and print its hash",
		"init", "Initialize a Zig package in the current directory",
		"build-exe", "Create executable from source or object files",
		"build-lib", "Create library from source or object files",
		"build-obj", "Create object from source or object files",
		"test", "Perform unit testing",
		"test-obj", "Create object for unit testing",
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
		"help", "Print help and exit",
		"std", "View standard library documentation in a browser",
		"libc", "Display native libc paths file or validate one",
		"targets", "List available compilation targets",
		"version", "Print version number and exit",
		"zen", "Print Zen of Zig and exit",
	)
}

func actionColor() carapace.Action {
	return carapace.ActionValues("auto", "off", "on")
}

func actionObjectFormats() carapace.Action {
	return carapace.ActionValues("elf", "c", "wasm", "coff", "macho", "spirv", "plan9", "hex", "raw")
}

func actionOptimizationModes() carapace.Action {
	return carapace.ActionValues("Debug", "ReleaseFast", "ReleaseSafe", "ReleaseSmall")
}

func actionTargets() carapace.Action {
	return carapace.ActionValues(
		"aarch64-linux",
		"aarch64-macos",
		"aarch64-windows",
		"arm-linux",
		"riscv64-linux",
		"wasm32-freestanding",
		"wasm32-wasi",
		"x86-linux",
		"x86_64-freebsd",
		"x86_64-linux",
		"x86_64-macos",
		"x86_64-windows",
	)
}

func sourceFileTypes() []string {
	return []string{".zig", ".o", ".obj", ".lib", ".a", ".so", ".dll", ".dylib", ".tbd", ".s", ".S", ".c", ".cxx", ".cc", ".C", ".cpp", ".c++", ".m", ".mm", ".bc"}
}
