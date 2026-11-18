package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundle_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the `Brewfile` in your editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_editCmd).Standalone()

	bundle_editCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	bundle_editCmd.Flags().BoolP("help", "h", false, "Show this message.")
	bundle_editCmd.Flags().Bool("install", false, "Run `install` before editing the `Brewfile`.")
	bundle_editCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	bundle_editCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	bundleCmd.AddCommand(bundle_editCmd)
}
