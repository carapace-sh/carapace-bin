package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// addBuildFlags registers common build flags shared by `zig build-exe`,
// `zig build-lib`, and `zig build-obj`.
func addBuildFlags(c *cobra.Command) {
	c.Flags().StringP("name", "", "", "Override root name")
	c.Flags().StringP("target", "", "", "<arch><sub>-<os>-<abi>")
	c.Flags().StringP("mcpu", "", "", "Specify target CPU and feature set")
	c.Flags().StringP("O", "", "", "Choose what to optimize for")
	c.Flags().Bool("strip", false, "Omit debug symbols")
	c.Flags().Bool("single-threaded", false, "Code assumes there is only one thread")
	c.Flags().Bool("static", false, "Output will be statically linked")
	c.Flags().Bool("dynamic", false, "Force output to be dynamically linked")
	c.Flags().Bool("dynamic-linker", false, "Set the dynamic interpreter path")
	c.Flags().StringP("femit-bin", "", "", "Output machine code")
	c.Flags().StringP("femit-asm", "", "", "Output assembly code")
	c.Flags().StringP("femit-llvm-ir", "", "", "Output the LLVM IR")
	c.Flags().StringP("femit-h", "", "", "Output a C header file")
	c.Flags().BoolP("help", "h", false, "Print help")

	carapace.Gen(c).FlagCompletion(carapace.ActionMap{
		"O": carapace.ActionValues(
			"Debug",
			"ReleaseSafe",
			"ReleaseFast",
			"ReleaseSmall",
		),
	})
}
