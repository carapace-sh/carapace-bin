package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_cherryPick_unresolveCmd = &cobra.Command{
	Use:   "unresolve",
	Short: "Marks the cherry-pick unresolved",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_cherryPick_unresolveCmd).Standalone()

	help_revision_cherryPickCmd.AddCommand(help_revision_cherryPick_unresolveCmd)
}
