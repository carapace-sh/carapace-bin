package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_rewordCmd = &cobra.Command{
	Use:   "reword",
	Short: "Replace a commit message and replay descendants onto the new history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_rewordCmd).Standalone()

	help_historyCmd.AddCommand(help_history_rewordCmd)
}
