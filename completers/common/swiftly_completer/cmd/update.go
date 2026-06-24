package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [toolchain]",
	Short: "Update an installed toolchain to a newer version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("assume-yes", false, "Disable confirmation prompts by assuming 'yes'")
	updateCmd.Flags().String("format", "text", "Output format (text|json)")
	updateCmd.Flags().BoolP("help", "h", false, "Show help information")
	updateCmd.Flags().Bool("no-verify", false, "Do not verify the toolchain's PGP signature before installation")
	updateCmd.Flags().String("post-install-file", "", "File path to a location for a post installation script")
	updateCmd.Flags().Bool("verify", true, "Verify the toolchain's PGP signature before installation")

	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"format":            actionFormat(),
		"post-install-file": carapace.ActionFiles(),
	})

	carapace.Gen(updateCmd).PositionalCompletion(
		actionToolchains(),
	)
}
