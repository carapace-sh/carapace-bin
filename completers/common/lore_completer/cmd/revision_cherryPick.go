package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var revision_cherryPickCmd = &cobra.Command{
	Use:   "cherry-pick",
	Short: "Cherry-pick a revision onto the currently synced revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_cherryPickCmd).Standalone()

	revision_cherryPickCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_cherryPickCmd.Flags().String("message", "", "Change the message for committing when no conflicts arise from the cherry-pick")
	revision_cherryPickCmd.Flags().Bool("no-commit", false, "Disable auto commits even if no conflicts arise from the cherry-pick")
	revisionCmd.AddCommand(revision_cherryPickCmd)

	carapace.Gen(revision_cherryPickCmd).PositionalCompletion(
		action.ActionRevisions(revision_cherryPickCmd),
	)
}
