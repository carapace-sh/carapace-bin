package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var historyRedoCmd = &cobra.Command{
	Use:   "redo [options] <transaction-spec>",
	Short: "repeat all actions from the specified transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(historyRedoCmd).Standalone()

	historyRedoCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	historyRedoCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	historyRedoCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	historyRedoCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	historyCmd.AddCommand(historyRedoCmd)

	carapace.Gen(historyRedoCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})
}
