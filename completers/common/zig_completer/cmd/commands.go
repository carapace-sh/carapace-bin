package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(
		buildCmd(),
		fetchCmd(),
		initCmd(),
		compileCmd("build-exe", "Create executable from source or object files", true),
		compileCmd("build-lib", "Create library from source or object files", true),
		compileCmd("build-obj", "Create object from source or object files", true),
		testCmd(),
		compileCmd("run", "Create executable and run immediately", true),
		astCheckCmd(),
		fmtCmd(),
		reduceCmd(),
		compileCmd("translate-c", "Convert C code to Zig code", false),
		passthroughCmd("ar", "Use Zig as a drop-in archiver"),
		passthroughCmd("cc", "Use Zig as a drop-in C compiler"),
		passthroughCmd("c++", "Use Zig as a drop-in C++ compiler"),
		passthroughCmd("dlltool", "Use Zig as a drop-in dlltool.exe"),
		passthroughCmd("lib", "Use Zig as a drop-in lib.exe"),
		passthroughCmd("ranlib", "Use Zig as a drop-in ranlib"),
		passthroughCmd("objcopy", "Use Zig as a drop-in objcopy"),
		passthroughCmd("rc", "Use Zig as a drop-in rc.exe"),
		simpleCmd("env", "Print lib path, std path, cache directory, and version"),
		simpleCmd("help", "Print this help and exit"),
		simpleCmd("std", "View standard library documentation in a browser"),
		libcCmd(),
		simpleCmd("targets", "List available compilation targets"),
		simpleCmd("version", "Print version number and exit"),
		simpleCmd("zen", "Print Zen of Zig and exit"),
	)
}

func buildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build project from build.zig",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	cmd.Flags().String("build-file", "", "Override path to build.zig")
	cmd.Flags().String("build-id", "", "Embed a build ID in binaries")
	cmd.Flags().String("build-runner", "", "Override path to build runner")
	cmd.Flags().String("cache-dir", "", "Override path to local Zig cache directory")
	cmd.Flags().String("color", "", "Enable or disable colored error messages")
	cmd.Flags().String("debounce", "", "Delay before rebuilding after file changes")
	cmd.Flags().String("debug-log", "", "Enable debugging the compiler")
	cmd.Flags().Bool("debug-pkg-config", false, "Fail if unknown pkg-config flags are encountered")
	cmd.Flags().Bool("debug-rt", false, "Debug compiler runtime libraries")
	cmd.Flags().BoolS("fallow-so-scripts", "fallow-so-scripts", false, "Allow .so files to be GNU ld scripts")
	cmd.Flags().BoolS("fdarling", "fdarling", false, "Enable Darling integration")
	cmd.Flags().String("fetch", "", "Fetch dependency tree and exit")
	cmd.Flags().BoolS("fincremental", "fincremental", false, "Enable incremental compilation")
	cmd.Flags().BoolS("fno-allow-so-scripts", "fno-allow-so-scripts", false, ".so files must be ELF files")
	cmd.Flags().BoolS("fno-darling", "fno-darling", false, "Disable Darling integration")
	cmd.Flags().BoolS("fno-incremental", "fno-incremental", false, "Disable incremental compilation")
	cmd.Flags().BoolS("fno-qemu", "fno-qemu", false, "Disable QEMU integration")
	cmd.Flags().BoolS("fno-reference-trace", "fno-reference-trace", false, "Disable reference trace")
	cmd.Flags().BoolS("fno-rosetta", "fno-rosetta", false, "Do not rely on Rosetta")
	cmd.Flags().StringS("fno-sys", "fno-sys", "", "Disable a system integration")
	cmd.Flags().BoolS("fno-wasmtime", "fno-wasmtime", false, "Disable wasmtime integration")
	cmd.Flags().BoolS("fno-wine", "fno-wine", false, "Disable Wine integration")
	cmd.Flags().BoolS("fqemu", "fqemu", false, "Enable QEMU integration")
	cmd.Flags().StringS("freference-trace", "freference-trace", "", "How many lines of reference trace should be shown per compile error")
	cmd.Flags().BoolS("frosetta", "frosetta", false, "Rely on Rosetta on ARM64 macOS hosts")
	cmd.Flags().StringS("fsys", "fsys", "", "Enable a system integration")
	cmd.Flags().Bool("fuzz", false, "Continuously search for unit test failures")
	cmd.Flags().BoolS("fwasmtime", "fwasmtime", false, "Enable wasmtime integration")
	cmd.Flags().BoolS("fwine", "fwine", false, "Enable Wine integration")
	cmd.Flags().String("global-cache-dir", "", "Override path to global Zig cache directory")
	cmd.Flags().StringS("j", "j", "", "Limit concurrent jobs")
	cmd.Flags().String("libc", "", "Provide a file which specifies libc paths")
	cmd.Flags().String("libc-runtimes", "", "Directory of libc runtimes for QEMU integration")
	cmd.Flags().String("maxrss", "", "Limit memory usage")
	cmd.Flags().StringS("p", "p", "", "Override default install prefix")
	cmd.Flags().String("prefix", "", "Override default install prefix")
	cmd.Flags().String("prefix-exe-dir", "", "Override default executable directory path")
	cmd.Flags().String("prefix-include-dir", "", "Override default include directory path")
	cmd.Flags().String("prefix-lib-dir", "", "Override default library directory path")
	cmd.Flags().Bool("prominent-compile-errors", false, "Output compile errors formatted for a human to read")
	cmd.Flags().String("release", "", "Request release mode")
	cmd.Flags().String("search-prefix", "", "Add a path to look for binaries, libraries, headers")
	cmd.Flags().String("seed", "", "Shuffle dependency traversal order")
	cmd.Flags().Bool("skip-oom-steps", false, "Skip steps that would exceed maxrss")
	cmd.Flags().String("summary", "", "Control the printing of the build summary")
	cmd.Flags().String("sysroot", "", "Set the system root directory")
	cmd.Flags().String("system", "", "Disable package fetching; enable all integrations")
	cmd.Flags().Bool("verbose", false, "Print commands before executing them")
	cmd.Flags().Bool("verbose-air", false, "Enable compiler debug output for Zig AIR")
	cmd.Flags().Bool("verbose-cc", false, "Enable compiler debug output for C compilation")
	cmd.Flags().Bool("verbose-cimport", false, "Enable compiler debug output for C imports")
	cmd.Flags().Bool("verbose-link", false, "Enable compiler debug output for linking")
	cmd.Flags().String("verbose-llvm-bc", "", "Enable compiler debug output for LLVM BC")
	cmd.Flags().Bool("verbose-llvm-cpu-features", false, "Enable compiler debug output for LLVM CPU features")
	cmd.Flags().String("verbose-llvm-ir", "", "Enable compiler debug output for LLVM IR")
	cmd.Flags().Bool("watch", false, "Continuously rebuild when source files are modified")
	cmd.Flags().String("zig-lib-dir", "", "Override path to Zig lib directory")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"build-file":         carapace.ActionFiles(".zig"),
		"build-id":           actionBuildIDs(),
		"build-runner":       carapace.ActionFiles(),
		"cache-dir":          carapace.ActionDirectories(),
		"color":              actionColorModes(),
		"fetch":              carapace.ActionValues("needed", "all").StyleF(style.ForKeyword),
		"global-cache-dir":   carapace.ActionDirectories(),
		"libc":               carapace.ActionFiles(),
		"libc-runtimes":      carapace.ActionDirectories(),
		"p":                  carapace.ActionDirectories(),
		"prefix":             carapace.ActionDirectories(),
		"prefix-exe-dir":     carapace.ActionDirectories(),
		"prefix-include-dir": carapace.ActionDirectories(),
		"prefix-lib-dir":     carapace.ActionDirectories(),
		"release":            carapace.ActionValues("fast", "safe", "small").StyleF(style.ForKeyword),
		"search-prefix":      carapace.ActionDirectories(),
		"summary":            carapace.ActionValues("all", "new", "failures", "none").StyleF(style.ForKeyword),
		"sysroot":            carapace.ActionDirectories(),
		"verbose-llvm-bc":    carapace.ActionFiles(".bc"),
		"verbose-llvm-ir":    carapace.ActionFiles(".ll"),
		"zig-lib-dir":        carapace.ActionDirectories(),
	})

	carapace.Gen(cmd).PositionalAnyCompletion(
		actionBuildSteps(),
	)

	return cmd
}

func fetchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "Copy a package into global cache and print its hash",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	cmd.Flags().Bool("debug-hash", false, "Print verbose hash information")
	cmd.Flags().String("global-cache-dir", "", "Override the global cache directory")
	cmd.Flags().String("save", "", "Add the fetched package to build.zig.zon")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"global-cache-dir": carapace.ActionDirectories(),
	})
	carapace.Gen(cmd).PositionalCompletion(carapace.ActionFiles())
	return cmd
}

func initCmd() *cobra.Command {
	cmd := simpleCmd("init", "Initialize a Zig package in the current directory")
	cmd.Flags().Bool("minimal", false, "Create a minimal package")
	return cmd
}

func compileCmd(use, short string, link bool) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	addCompileFlags(cmd)
	if link {
		addLinkFlags(cmd)
	}
	addDebugFlags(cmd)
	addCompileCompletions(cmd)
	return cmd
}

func testCmd() *cobra.Command {
	cmd := compileCmd("test", "Perform unit testing", true)
	addTestFlags(cmd)
	return cmd
}

func fmtCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fmt",
		Short: "Reformat Zig source into canonical form",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	cmd.Flags().Bool("ast-check", false, "Run zig ast-check on every file")
	cmd.Flags().Bool("check", false, "List non-conforming files and exit with an error")
	cmd.Flags().String("color", "", "Enable or disable colored error messages")
	cmd.Flags().String("exclude", "", "Exclude file or directory from formatting")
	cmd.Flags().Bool("stdin", false, "Format code from stdin; output to stdout")
	cmd.Flags().Bool("zon", false, "Treat all input files as ZON")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"color":   actionColorModes(),
		"exclude": carapace.ActionFiles(),
	})
	carapace.Gen(cmd).PositionalAnyCompletion(actionZigSourceFiles())
	return cmd
}

func astCheckCmd() *cobra.Command {
	cmd := simpleCmd("ast-check", "Look for simple compile errors in any set of files")
	carapace.Gen(cmd).PositionalAnyCompletion(actionZigSourceFiles())
	return cmd
}

func reduceCmd() *cobra.Command {
	cmd := simpleCmd("reduce", "Minimize a bug report")
	carapace.Gen(cmd).PositionalAnyCompletion(actionZigSourceFiles())
	return cmd
}

func libcCmd() *cobra.Command {
	cmd := simpleCmd("libc", "Display native libc paths file or validate one")
	carapace.Gen(cmd).PositionalCompletion(carapace.ActionFiles())
	return cmd
}

func passthroughCmd(use, short string) *cobra.Command {
	cmd := simpleCmd(use, short)
	carapace.Gen(cmd).PositionalAnyCompletion(carapace.ActionFiles())
	return cmd
}

func simpleCmd(use, short string) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
}
