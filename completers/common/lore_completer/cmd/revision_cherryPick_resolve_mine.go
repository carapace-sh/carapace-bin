package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_cherryPick_resolve_mineCmd = &cobra.Command{
	Use:   "mine",
	Short: "Resolve using my changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_cherryPick_resolve_mineCmd).Standalone()

	revision_cherryPick_resolve_mineCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_cherryPick_resolve_mineCmd.Flags().String("targets", "", "Path to a targets file")
	revision_cherryPick_resolveCmd.AddCommand(revision_cherryPick_resolve_mineCmd)
}
