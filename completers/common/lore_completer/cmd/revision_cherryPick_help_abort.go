package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_cherryPick_help_abortCmd = &cobra.Command{
	Use:   "abort",
	Short: "Abort a cherry-pick",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_cherryPick_help_abortCmd).Standalone()

	revision_cherryPick_helpCmd.AddCommand(revision_cherryPick_help_abortCmd)
}
