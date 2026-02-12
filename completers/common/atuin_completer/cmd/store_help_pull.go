package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_help_pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull records from the remote sync server (one way sync)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_help_pullCmd).Standalone()

	store_helpCmd.AddCommand(store_help_pullCmd)
}
