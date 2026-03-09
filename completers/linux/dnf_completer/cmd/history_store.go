package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_storeCmd = &cobra.Command{
	Use:   "store [transaction-spec]",
	Short: "store the transaction into a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_storeCmd).Standalone()

	historyCmd.AddCommand(history_storeCmd)
}
