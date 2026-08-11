package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_cherryPick_help_unresolveCmd = &cobra.Command{
	Use:   "unresolve",
	Short: "Marks the cherry-pick unresolved",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_cherryPick_help_unresolveCmd).Standalone()

	revision_cherryPick_helpCmd.AddCommand(revision_cherryPick_help_unresolveCmd)
}
