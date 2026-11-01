package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo the last applied history rewrite using the recorded ref plan",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_undoCmd).Standalone()

	help_historyCmd.AddCommand(help_history_undoCmd)
}
