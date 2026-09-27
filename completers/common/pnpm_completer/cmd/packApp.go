package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packAppCmd = &cobra.Command{
	Use:   "pack-app",
	Short: "Pack a `CommonJS` entry file into a standalone executable for one or more target platforms",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packAppCmd).Standalone()

	packAppCmd.Flags().String("entry", "", "Path to the CJS entry file to embed in the executable")
	packAppCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	packAppCmd.Flags().StringP("output-dir", "o", "", "Output directory for the built executables. Defaults to `dist-app`")
	packAppCmd.Flags().String("output-name", "", "Name for the output executable (without extension). Defaults to the unscoped package name")
	packAppCmd.Flags().String("runtime", "", "Runtime to embed, as a `<name>@<version>` spec (e.g. `node@25`, `node@25.5.0`). Only `node` is supported, and the version must be >= v25.5. Defaults to the minimum SEA-capable version (v25.5.0)")
	packAppCmd.Flags().StringSliceP("target", "t", nil, "Target to build for. May be specified multiple times. Supported: linux-x64, linux-x64-musl, linux-arm64, linux-arm64-musl, darwin-x64, darwin-arm64, win32-x64, win32-arm64")
	rootCmd.AddCommand(packAppCmd)
}
