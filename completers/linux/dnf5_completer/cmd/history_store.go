package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var historyStoreCmd = &cobra.Command{
	Use:   "store [options] [<transaction-spec>...]",
	Short: "store transaction to a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(historyStoreCmd).Standalone()

	historyStoreCmd.Flags().StringP("output", "o", "", "Path to a directory for storing the transaction")

	historyCmd.AddCommand(historyStoreCmd)

	carapace.Gen(historyStoreCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionDirectories(),
	})
}
