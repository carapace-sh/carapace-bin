package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_cherryPick_resolve_help_theirsCmd = &cobra.Command{
	Use:   "theirs",
	Short: "Resolve using the incoming changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_cherryPick_resolve_help_theirsCmd).Standalone()

	revision_cherryPick_resolve_helpCmd.AddCommand(revision_cherryPick_resolve_help_theirsCmd)
}
