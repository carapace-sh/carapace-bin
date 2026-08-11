package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_cherryPick_unresolveCmd = &cobra.Command{
	Use:   "unresolve",
	Short: "Marks the cherry-pick unresolved",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_cherryPick_unresolveCmd).Standalone()

	revision_cherryPick_unresolveCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_cherryPick_unresolveCmd.Flags().String("targets", "", "Path to a targets file")
	revision_cherryPickCmd.AddCommand(revision_cherryPick_unresolveCmd)
}
