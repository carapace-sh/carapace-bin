package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "Replay local commits onto a new base (same as top-level `grit rebase`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_rebaseCmd).Standalone()

	history_helpCmd.AddCommand(history_help_rebaseCmd)
}
