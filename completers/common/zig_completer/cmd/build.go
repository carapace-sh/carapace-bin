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
	buildCmd.Flags().String("build-runner", "", "Override path to build runner")
	buildCmd.Flags().String("cache-dir", "", "Override the local cache directory")
	buildCmd.Flags().String("color", "", "Enable or disable colored output")
	buildCmd.Flags().Bool("debug-compile-errors", false, "Crash with diagnostics at first compile error")
	buildCmd.Flags().String("debug-log", "", "Enable printing debug/info log messages for scope")
	buildCmd.Flags().String("fetch", "", "Fetch all packages and exit")
	buildCmd.Flags().String("global-cache-dir", "", "Override the global cache directory")
	buildCmd.Flags().BoolP("help", "h", false, "Print help")
	buildCmd.Flags().String("job-count", "", "Limit concurrent jobs")
	buildCmd.Flags().Bool("reference-trace", false, "Enable reference trace (256 lines)")
	buildCmd.Flags().String("seed", "", "Override the random seed for the build")
	buildCmd.Flags().String("system", "", "Set system package directory path")
	buildCmd.Flags().Bool("verbose-air", false, "Enable compiler debug output for Zig AIR")
	buildCmd.Flags().Bool("verbose-cc", false, "Display C compiler invocations")
	buildCmd.Flags().Bool("verbose-cimport", false, "Enable compiler debug output for C imports")
	buildCmd.Flags().Bool("verbose-generic-instances", false, "Enable compiler debug output for generic instance generation")
	buildCmd.Flags().Bool("verbose-intern-pool", false, "Enable compiler debug output for InternPool")
	buildCmd.Flags().Bool("verbose-link", false, "Display linker invocations")
	buildCmd.Flags().String("verbose-llvm-bc", "", "Enable compiler debug output for unoptimized LLVM BC")
	buildCmd.Flags().Bool("verbose-llvm-cpu-features", false, "Enable compiler debug output for LLVM CPU features")
	buildCmd.Flags().String("verbose-llvm-ir", "", "Enable compiler debug output for unoptimized LLVM IR")
	buildCmd.Flags().String("zig-lib-dir", "", "Override path to Zig installation lib directory")

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"build-file":       carapace.ActionFiles("build.zig"),
		"build-runner":     carapace.ActionFiles(),
		"cache-dir":        carapace.ActionDirectories(),
		"color":            carapace.ActionValues("auto", "off", "on").StyleF(style.ForKeyword),
		"fetch":            carapace.ActionValues("needed", "all"),
		"global-cache-dir": carapace.ActionDirectories(),
		"system":           carapace.ActionDirectories(),
		"verbose-llvm-bc":  carapace.ActionFiles(),
		"verbose-llvm-ir":  carapace.ActionFiles(),
		"zig-lib-dir":      carapace.ActionDirectories(),
	})
}
