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

	vendorGemsCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	vendorGemsCmd.Flags().BoolP("help", "h", false, "Show this message.")
	vendorGemsCmd.Flags().Bool("no-commit", false, "Do not generate a new commit upon completion.")
	vendorGemsCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	vendorGemsCmd.Flags().String("update", "", "Update the specified list of vendored gems to the latest version.")
	vendorGemsCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(vendorGemsCmd)
}
