package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:                "build",
	Short:              "Build project from build.zig",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionValues(
				"-Doptimize=Debug",
				"-Doptimize=ReleaseSafe",
				"-Doptimize=ReleaseFast",
				"-Doptimize=ReleaseSmall",
				"-Dtarget=native",
				"-Dtarget=x86_64-linux",
				"-Dtarget=x86_64-linux-gnu",
				"-Dtarget=x86_64-linux-musl",
				"-Dtarget=x86_64-windows",
				"-Dtarget=x86_64-windows-gnu",
				"-Dtarget=x86_64-windows-msvc",
				"-Dtarget=x86_64-macos",
				"-Dtarget=aarch64-linux",
				"-Dtarget=aarch64-linux-gnu",
				"-Dtarget=aarch64-linux-musl",
				"-Dtarget=aarch64-macos",
				"-Dtarget=aarch64-windows",
				"-Dtarget=wasm32-wasi",
				"-Dtarget=wasm32-freestanding",
			),
			carapace.ActionValues(
				"--help",
				"--build-file",
				"--cache-dir",
				"--global-cache-dir",
				"--zig-lib-dir",
				"--prefix",
				"--prefix-lib-dir",
				"--prefix-exe-dir",
				"--prefix-include-dir",
				"--release",
				"--verbose",
				"--summary",
				"--prominent-compile-errors",
				"--watch",
				"--fetch",
			),
		).ToA(),
	)
}
