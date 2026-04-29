package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Stream history events from the daemon as they are received",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_tailCmd).Standalone()

	help_historyCmd.AddCommand(help_history_tailCmd)
}
