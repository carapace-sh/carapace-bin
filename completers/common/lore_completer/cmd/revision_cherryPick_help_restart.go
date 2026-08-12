package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_cherryPick_help_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the cherry-pick, resetting the current cherry-pick state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_cherryPick_help_restartCmd).Standalone()

	revision_cherryPick_helpCmd.AddCommand(revision_cherryPick_help_restartCmd)
}
