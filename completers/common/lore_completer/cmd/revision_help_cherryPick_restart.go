package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_cherryPick_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the cherry-pick, resetting the current cherry-pick state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_cherryPick_restartCmd).Standalone()

	revision_help_cherryPickCmd.AddCommand(revision_help_cherryPick_restartCmd)
}
