package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Create a new, empty change and edit it in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_checkoutCmd).Standalone()

	helpCmd.AddCommand(help_checkoutCmd)
}
