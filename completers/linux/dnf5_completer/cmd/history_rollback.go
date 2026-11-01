package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var historyRollbackCmd = &cobra.Command{
	Use:   "rollback [options] <transaction-spec>",
	Short: "undo all transactions performed after the specified transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(historyRollbackCmd).Standalone()

	historyRollbackCmd.Flags().Bool("ignore-extras", false, "Ignore extra packages")
	historyRollbackCmd.Flags().Bool("ignore-installed", false, "Ignore installed packages")
	historyRollbackCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	historyRollbackCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	historyRollbackCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	historyRollbackCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	historyCmd.AddCommand(historyRollbackCmd)

	carapace.Gen(historyRollbackCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})
}
