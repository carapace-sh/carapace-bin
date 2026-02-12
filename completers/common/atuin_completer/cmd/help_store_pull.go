package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_store_pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull records from the remote sync server (one way sync)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_store_pullCmd).Standalone()

	help_storeCmd.AddCommand(help_store_pullCmd)
}
