package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_undoCmd = &cobra.Command{
	Use:   "undo <transaction-spec>",
	Short: "revert all actions from the specified transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_undoCmd).Standalone()

	history_undoCmd.Flags().Bool("ignore-extras", false, "don't consider extra packages as errors")
	history_undoCmd.Flags().Bool("ignore-installed", false, "don't consider mismatches as errors")
	history_undoCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages")
	historyCmd.AddCommand(history_undoCmd)
}
