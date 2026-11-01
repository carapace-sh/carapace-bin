package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Reorder a commit on the first-parent chain of `HEAD`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_moveCmd).Standalone()

	history_helpCmd.AddCommand(history_help_moveCmd)
}
