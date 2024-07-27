package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scoutDeleteAllIndexesCmd = &cobra.Command{
	Use:   "scout:delete-all-indexes",
	Short: "Delete all indexes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scoutDeleteAllIndexesCmd).Standalone()

	rootCmd.AddCommand(scoutDeleteAllIndexesCmd)
}
