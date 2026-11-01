package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sharedStore_help_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sharedStore_help_infoCmd).Standalone()

	sharedStore_helpCmd.AddCommand(sharedStore_help_infoCmd)
}
