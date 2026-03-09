package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_rollbackCmd = &cobra.Command{
	Use:   "rollback <transaction-spec>",
	Short: "undo all transactions after the specified transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_rollbackCmd).Standalone()

	history_rollbackCmd.Flags().Bool("ignore-extras", false, "don't consider extra packages as errors")
	history_rollbackCmd.Flags().Bool("ignore-installed", false, "don't consider mismatches as errors")
	history_rollbackCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages")
	historyCmd.AddCommand(history_rollbackCmd)
}
