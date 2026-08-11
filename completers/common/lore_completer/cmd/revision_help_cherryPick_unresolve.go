package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_cherryPick_unresolveCmd = &cobra.Command{
	Use:   "unresolve",
	Short: "Marks the cherry-pick unresolved",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_cherryPick_unresolveCmd).Standalone()

	revision_help_cherryPickCmd.AddCommand(revision_help_cherryPick_unresolveCmd)
}
