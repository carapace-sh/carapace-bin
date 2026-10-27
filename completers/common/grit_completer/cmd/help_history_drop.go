package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Remove a commit and replay descendants onto its parent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_dropCmd).Standalone()

	help_historyCmd.AddCommand(help_history_dropCmd)
}
