package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_endCmd = &cobra.Command{
	Use:   "end",
	Short: "Finishes a new command in the history (adds time, exit code)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_endCmd).Standalone()

	help_historyCmd.AddCommand(help_history_endCmd)
}
