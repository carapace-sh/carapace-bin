package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_cherryPick_resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve conflicts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_cherryPick_resolveCmd).Standalone()

	revision_cherryPick_resolveCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_cherryPick_resolveCmd.Flags().String("targets", "", "Path to a targets file")
	revision_cherryPickCmd.AddCommand(revision_cherryPick_resolveCmd)

	carapace.Gen(revision_cherryPick_resolveCmd).FlagCompletion(carapace.ActionMap{
		"targets": carapace.ActionFiles(),
	})

	carapace.Gen(revision_cherryPick_resolveCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
