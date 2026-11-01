package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_cherryPick_resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve conflicts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_cherryPick_resolveCmd).Standalone()

	revision_help_cherryPickCmd.AddCommand(revision_help_cherryPick_resolveCmd)
}
