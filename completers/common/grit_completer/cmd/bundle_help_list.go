package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundle_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List refs recorded in a bundle header",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_help_listCmd).Standalone()

	bundle_helpCmd.AddCommand(bundle_help_listCmd)
}
