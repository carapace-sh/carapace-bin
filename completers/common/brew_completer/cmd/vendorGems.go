package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vendorGemsCmd = &cobra.Command{
	Use:     "vendor-gems",
	Short:   "Install and commit Homebrew's vendored gems",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vendorGemsCmd).Standalone()

	vendorGemsCmd.Flags().Bool("debug", false, "Display any debugging information.")
	vendorGemsCmd.Flags().Bool("help", false, "Show this message.")
	vendorGemsCmd.Flags().Bool("no-commit", false, "Do not generate a new commit upon completion.")
	vendorGemsCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	vendorGemsCmd.Flags().Bool("update", false, "Update the specified list of vendored gems to the latest version.")
	vendorGemsCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(vendorGemsCmd)
}
