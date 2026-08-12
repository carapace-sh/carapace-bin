package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_cherryPick_resolve_mineCmd = &cobra.Command{
	Use:   "mine",
	Short: "Resolve using my changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_cherryPick_resolve_mineCmd).Standalone()

	help_revision_cherryPick_resolveCmd.AddCommand(help_revision_cherryPick_resolve_mineCmd)
}
