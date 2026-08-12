package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_cherryPick_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the cherry-pick, resetting the current cherry-pick state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_cherryPick_restartCmd).Standalone()

	help_revision_cherryPickCmd.AddCommand(help_revision_cherryPick_restartCmd)
}
