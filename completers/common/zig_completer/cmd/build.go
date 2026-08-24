package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build [options]",
	Short:   "Build project from build.zig",
	GroupID: "project",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().String("build-file", "", "Override path to build.zig")
	buildCmd.Flags().String("build-id", "", "Embed a build ID in binaries")
	buildCmd.Flags().String("build-runner", "", "Override path to build runner")
	buildCmd.Flags().String("cache-dir", "", "Override the local cache directory")
	buildCmd.Flags().String("color", "", "Enable or disable colored output")
	buildCmd.Flags().Bool("darling", false, "Integration with system-installed Darling to execute macOS programs on Linux hosts")
	buildCmd.Flags().String("debounce", "", "Delay before rebuilding after changed file detected")
	buildCmd.Flags().Bool("debug-compile-errors", false, "Crash with diagnostics at first compile error")
	buildCmd.Flags().String("debug-log", "", "Enable printing debug/info log messages for scope")
	buildCmd.Flags().Bool("debug-pkg-config", false, "Fail if unknown pkg-config flags encountered")
	buildCmd.Flags().Bool("debug-rt", false, "Debug compiler runtime libraries")
	buildCmd.Flags().String("error-style", "", "Control how build errors are printed")
	buildCmd.Flags().Bool("fallow-so-scripts", false, "Allows .so files to be GNU ld scripts")
	buildCmd.Flags().String("fetch", "", "Fetch all packages and exit")
	buildCmd.Flags().Bool("fno-allow-so-scripts", false, ".so files must be ELF files")
	buildCmd.Flags().String("fork", "", "Override one or more projects from dependency tree")
	buildCmd.Flags().String("fuzz", "", "Continuously search for unit test failures")
	buildCmd.Flags().String("global-cache-dir", "", "Override the global cache directory")
	buildCmd.Flags().BoolP("help", "h", false, "Print help")
	buildCmd.Flags().Bool("incremental", false, "Enable incremental compilation")
	buildCmd.Flags().String("job-count", "", "Limit concurrent jobs")
	buildCmd.Flags().String("libc", "", "Provide a file which specifies libc paths")
	buildCmd.Flags().String("libc-runtimes", "", "Enhances QEMU integration by providing dynamic libc")
	buildCmd.Flags().BoolP("list-steps", "l", false, "Print available steps")
	buildCmd.Flags().String("maxrss", "", "Limit memory usage")
	buildCmd.Flags().String("multiline-errors", "", "Control how multi-line error messages are printed")
	buildCmd.Flags().Bool("no-darling", false, "Disable Darling integration")
	buildCmd.Flags().Bool("no-incremental", false, "Disable incremental compilation")
	buildCmd.Flags().Bool("no-qemu", false, "Disable QEMU integration")
	buildCmd.Flags().Bool("no-rosetta", false, "Disable Rosetta integration")
	buildCmd.Flags().String("no-sys", "", "Disable a system integration")
	buildCmd.Flags().Bool("no-wasmtime", false, "Disable wasmtime integration")
	buildCmd.Flags().Bool("no-wine", false, "Disable Wine integration")
	buildCmd.Flags().String("prefix", "", "Where to install files (default: zig-out)")
	buildCmd.Flags().String("prefix-exe-dir", "", "Where to install executables")
	buildCmd.Flags().String("prefix-include-dir", "", "Where to install C header files")
	buildCmd.Flags().String("prefix-lib-dir", "", "Where to install libraries")
	buildCmd.Flags().Bool("qemu", false, "Integration with system-installed QEMU to execute foreign-architecture programs on Linux hosts")
	buildCmd.Flags().Bool("reference-trace", false, "Enable reference trace (256 lines)")
	buildCmd.Flags().String("release", "", "Request release mode")
	buildCmd.Flags().Bool("rosetta", false, "Rely on Rosetta to execute x86_64 programs on ARM64 macOS hosts")
	buildCmd.Flags().String("search-prefix", "", "Add a path to look for binaries, libraries, headers")
	buildCmd.Flags().String("seed", "", "Override the random seed for the build")
	buildCmd.Flags().Bool("skip-oom-steps", false, "Skip steps that would exceed --maxrss")
	buildCmd.Flags().String("summary", "", "Control the printing of the build summary")
	buildCmd.Flags().String("sys", "", "Enable a system integration")
	buildCmd.Flags().String("sysroot", "", "Set the system root directory (usually /)")
	buildCmd.Flags().String("system", "", "Disable package fetching; enable all integrations")
	buildCmd.Flags().String("test-timeout", "", "Limit execution time of unit tests")
	buildCmd.Flags().Bool("time-report", false, "Force full rebuild and provide detailed information on compilation time")
	buildCmd.Flags().Bool("verbose", false, "Print commands before executing them")
	buildCmd.Flags().Bool("verbose-air", false, "Enable compiler debug output for Zig AIR")
	buildCmd.Flags().Bool("verbose-cc", false, "Display C compiler invocations")
	buildCmd.Flags().Bool("verbose-cimport", false, "Enable compiler debug output for C imports")
	buildCmd.Flags().Bool("verbose-generic-instances", false, "Enable compiler debug output for generic instance generation")
	buildCmd.Flags().Bool("verbose-intern-pool", false, "Enable compiler debug output for InternPool")
	buildCmd.Flags().Bool("verbose-link", false, "Display linker invocations")
	buildCmd.Flags().String("verbose-llvm-bc", "", "Enable compiler debug output for unoptimized LLVM BC")
	buildCmd.Flags().Bool("verbose-llvm-cpu-features", false, "Enable compiler debug output for LLVM CPU features")
	buildCmd.Flags().String("verbose-llvm-ir", "", "Enable compiler debug output for unoptimized LLVM IR")
	buildCmd.Flags().Bool("wasmtime", false, "Integration with system-installed wasmtime to execute WASI binaries")
	buildCmd.Flags().Bool("watch", false, "Continuously rebuild when source files are modified")
	buildCmd.Flags().String("webui", "", "Enable the web interface on the given IP address")
	buildCmd.Flags().Bool("wine", false, "Integration with system-installed Wine to execute Windows programs on Linux hosts")
	buildCmd.Flags().String("zig-lib-dir", "", "Override path to Zig installation lib directory")

	buildCmd.Flags().Lookup("build-id").NoOptDefVal = " "
	buildCmd.Flags().Lookup("fuzz").NoOptDefVal = " "
	buildCmd.Flags().Lookup("reference-trace").NoOptDefVal = " "
	buildCmd.Flags().Lookup("release").NoOptDefVal = " "
	buildCmd.Flags().Lookup("webui").NoOptDefVal = " "

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"build-file":         carapace.ActionFiles("build.zig"),
		"build-id":           carapace.ActionValues("fast", "sha1", "tree", "md5", "uuid", "none"),
		"build-runner":       carapace.ActionFiles(),
		"cache-dir":          carapace.ActionDirectories(),
		"color":              carapace.ActionValues("auto", "off", "on").StyleF(style.ForKeyword),
		"error-style":        carapace.ActionValues("verbose", "minimal", "verbose_clear", "minimal_clear"),
		"fetch":              carapace.ActionValues("needed", "all"),
		"fork":               carapace.ActionFiles(),
		"global-cache-dir":   carapace.ActionDirectories(),
		"libc":               carapace.ActionFiles(),
		"libc-runtimes":      carapace.ActionDirectories(),
		"multiline-errors":   carapace.ActionValues("indent", "newline", "none"),
		"prefix":             carapace.ActionDirectories(),
		"prefix-exe-dir":     carapace.ActionDirectories(),
		"prefix-include-dir": carapace.ActionDirectories(),
		"prefix-lib-dir":     carapace.ActionDirectories(),
		"release":            carapace.ActionValues("fast", "safe", "small"),
		"search-prefix":      carapace.ActionDirectories(),
		"summary":            carapace.ActionValues("all", "new", "failures", "line", "none"),
		"sysroot":            carapace.ActionDirectories(),
		"system":             carapace.ActionDirectories(),
		"verbose-llvm-bc":    carapace.ActionFiles(),
		"verbose-llvm-ir":    carapace.ActionFiles(),
		"zig-lib-dir":        carapace.ActionDirectories(),
	})
}
