package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generateCiCmd = &cobra.Command{
	Use:   "generate-ci",
	Short: "Generate CI configuration",
	Args:  cobra.ExactArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateCiCmd).Standalone()

	generateCiCmd.Flags().StringP("manifest-path", "m", "", "Path to Cargo.toml")
	generateCiCmd.Flags().StringP("output", "o", "-", "Output path")
	generateCiCmd.Flags().StringArray("platform", nil, "Platform support")
	generateCiCmd.Flags().Bool("pytest", false, "Enable pytest")
	generateCiCmd.Flags().Bool("skip-attestation", false, "Skip artifact attestation")
	generateCiCmd.Flags().CountP("verbose", "v", "Use verbose output")
	generateCiCmd.Flags().Bool("zig", false, "Use zig to do cross compilation")
	rootCmd.AddCommand(generateCiCmd)

	carapace.Gen(generateCiCmd).FlagCompletion(carapace.ActionMap{
		"output":   carapace.ActionFiles(),
		"platform": carapace.ActionValues("all", "manylinux", "musllinux", "windows", "macos", "emscripten", "android"),
	})

	carapace.Gen(generateCiCmd).PositionalCompletion(
		carapace.ActionValues("github"),
	)
}
