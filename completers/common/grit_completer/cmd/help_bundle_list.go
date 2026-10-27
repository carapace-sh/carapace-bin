package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_bundle_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List refs recorded in a bundle header",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_bundle_listCmd).Standalone()

	help_bundleCmd.AddCommand(help_bundle_listCmd)
}
