package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_cherryPick_resolve_theirsCmd = &cobra.Command{
	Use:   "theirs",
	Short: "Resolve using the incoming changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_cherryPick_resolve_theirsCmd).Standalone()

	help_revision_cherryPick_resolveCmd.AddCommand(help_revision_cherryPick_resolve_theirsCmd)
}
