package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_cherryPickCmd = &cobra.Command{
	Use:   "cherry-pick",
	Short: "Cherry-pick a revision onto the currently synced revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_cherryPickCmd).Standalone()

	help_revisionCmd.AddCommand(help_revision_cherryPickCmd)
}
