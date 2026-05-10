package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

func newBuildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build [steps] [options]",
		Short: "Build project from build.zig",
	}

	cmd.Flags().String("build-file", "", "Override path to build.zig")
	cmd.Flags().String("cache-dir", "", "Override the local cache directory")
	cmd.Flags().String("global-cache-dir", "", "Override the global cache directory")
	cmd.Flags().String("zig-lib-dir", "", "Override path to Zig lib directory")
	cmd.Flags().Bool("build-runner", false, "Override path to build runner")
	cmd.Flags().BoolP("verbose", "v", false, "Print commands before executing them")
	cmd.Flags().Bool("color", false, "Enable or disable colored error messages")
	cmd.Flags().StringP("jobs", "j", "", "Limit concurrent jobs")
	cmd.Flags().Bool("prominent-compile-errors", false, "Output compile errors formatted for a human to read")
	cmd.Flags().Bool("fetch", false, "Exit after fetching dependency tree")
	cmd.Flags().String("prefix", "", "Override default install prefix")
	cmd.Flags().String("prefix-lib-dir", "", "Override default library directory within prefix")
	cmd.Flags().String("prefix-exe-dir", "", "Override default executable directory within prefix")
	cmd.Flags().String("prefix-include-dir", "", "Override default include directory within prefix")
	cmd.Flags().String("sysroot", "", "Set the system root directory")
	cmd.Flags().String("search-prefix", "", "Add a path to look for binaries, libraries, headers")
	cmd.Flags().String("libc", "", "Provide a file which specifies libc paths")
	cmd.Flags().String("target", "", "The CPU architecture, OS, and ABI to build for")
	cmd.Flags().String("mcpu", "", "Specify target CPU and feature set")
	cmd.Flags().String("dynamic-linker", "", "Set the dynamic interpreter path")
	cmd.Flags().String("watch", "", "Hot-reload mode")
	cmd.Flags().Bool("fqemu", false, "Enable QEMU integration")
	cmd.Flags().Bool("fno-qemu", false, "Disable QEMU integration")
	cmd.Flags().String("qemu-user", "", "QEMU user-mode executable")
	cmd.Flags().Bool("fwine", false, "Enable Wine integration")
	cmd.Flags().Bool("fno-wine", false, "Disable Wine integration")
	cmd.Flags().Bool("fwasmtime", false, "Enable Wasmtime integration")
	cmd.Flags().Bool("fno-wasmtime", false, "Disable Wasmtime integration")
	cmd.Flags().Bool("frosetta", false, "Enable Rosetta integration")
	cmd.Flags().Bool("fno-rosetta", false, "Disable Rosetta integration")
	cmd.Flags().String("summary", "", "Control the printing of the build summary")
	cmd.Flags().Bool("help", false, "Print this help and exit")
	cmd.Flags().Bool("verbose-link", false, "Display linker invocations")
	cmd.Flags().Bool("verbose-air", false, "Enable compiler debug output for Zig AIR")
	cmd.Flags().Bool("verbose-llvm-ir", false, "Enable compiler debug output for LLVM IR")
	cmd.Flags().Bool("verbose-llvm-bc", false, "Enable compiler debug output for LLVM BC")
	cmd.Flags().Bool("verbose-cimport", false, "Enable compiler debug output for C imports")
	cmd.Flags().Bool("verbose-cc", false, "Enable compiler debug output for C compilation")
	cmd.Flags().Bool("verbose-llvm-cpu-features", false, "Enable compiler debug output for LLVM CPU features")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"build-file": carapace.ActionFiles("zig"),
		"cache-dir":  carapace.ActionDirectories(),
		"global-cache-dir": carapace.ActionDirectories(),
		"zig-lib-dir":      carapace.ActionDirectories(),
		"prefix":           carapace.ActionDirectories(),
		"prefix-lib-dir":   carapace.ActionDirectories(),
		"prefix-exe-dir":   carapace.ActionDirectories(),
		"prefix-include-dir": carapace.ActionDirectories(),
		"sysroot":          carapace.ActionDirectories(),
		"search-prefix":    carapace.ActionDirectories(),
		"libc":             carapace.ActionFiles(),
		"color": carapace.ActionValues("auto", "off", "on").StyleF(style.ForKeyword),
		"summary": carapace.ActionValues("all", "new", "failures", "none"),
	})

	return cmd
}
