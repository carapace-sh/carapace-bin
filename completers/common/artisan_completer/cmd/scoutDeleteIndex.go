package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scoutDeleteIndexCmd = &cobra.Command{
	Use:   "scout:delete-index <name>",
	Short: "Delete an index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scoutDeleteIndexCmd).Standalone()

	rootCmd.AddCommand(scoutDeleteIndexCmd)
}
