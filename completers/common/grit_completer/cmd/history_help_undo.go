package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo the last applied history rewrite using the recorded ref plan",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_undoCmd).Standalone()

	history_helpCmd.AddCommand(history_help_undoCmd)
}
