package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var historyUndoCmd = &cobra.Command{
	Use:   "undo [options] <transaction-spec>",
	Short: "revert all actions from the specified transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(historyUndoCmd).Standalone()

	historyUndoCmd.Flags().Bool("ignore-extras", false, "Ignore extra packages")
	historyUndoCmd.Flags().Bool("ignore-installed", false, "Ignore installed packages")
	historyUndoCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	historyUndoCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	historyUndoCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	historyUndoCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	historyCmd.AddCommand(historyUndoCmd)

	carapace.Gen(historyUndoCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})
}
