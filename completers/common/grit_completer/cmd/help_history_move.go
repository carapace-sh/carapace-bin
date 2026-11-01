package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Reorder a commit on the first-parent chain of `HEAD`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_moveCmd).Standalone()

	help_historyCmd.AddCommand(help_history_moveCmd)
}
