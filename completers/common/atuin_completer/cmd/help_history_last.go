package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_lastCmd = &cobra.Command{
	Use:   "last",
	Short: "Get the last command ran",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_lastCmd).Standalone()

	help_historyCmd.AddCommand(help_history_lastCmd)
}
