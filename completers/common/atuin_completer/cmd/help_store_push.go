package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_store_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push all records to the remote sync server (one way sync)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_store_pushCmd).Standalone()

	help_storeCmd.AddCommand(help_store_pushCmd)
}
