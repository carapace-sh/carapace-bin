package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_cherryPickCmd = &cobra.Command{
	Use:   "cherry-pick",
	Short: "Cherry-pick a revision onto the currently synced revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_cherryPickCmd).Standalone()

	revision_helpCmd.AddCommand(revision_help_cherryPickCmd)
}
