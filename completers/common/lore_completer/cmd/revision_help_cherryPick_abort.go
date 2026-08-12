package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_cherryPick_abortCmd = &cobra.Command{
	Use:   "abort",
	Short: "Abort a cherry-pick",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_cherryPick_abortCmd).Standalone()

	revision_help_cherryPickCmd.AddCommand(revision_help_cherryPick_abortCmd)
}
