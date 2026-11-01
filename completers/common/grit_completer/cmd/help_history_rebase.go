package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "Replay local commits onto a new base (same as top-level `grit rebase`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_rebaseCmd).Standalone()

	help_historyCmd.AddCommand(help_history_rebaseCmd)
}
