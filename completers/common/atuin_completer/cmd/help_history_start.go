package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Begins a new command in the history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_startCmd).Standalone()

	help_historyCmd.AddCommand(help_history_startCmd)
}
