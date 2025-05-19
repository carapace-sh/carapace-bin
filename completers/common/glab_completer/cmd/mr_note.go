package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_noteCmd = &cobra.Command{
	Use:     "note [<id> | <branch>]",
	Short:   "Add a comment or note to a merge request.",
	Aliases: []string{"comment"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_noteCmd).Standalone()

	mr_noteCmd.Flags().StringP("message", "m", "", "Comment or note message.")
	mr_noteCmd.Flags().Bool("unique", false, "Don't create a comment or note if it already exists.")
	mrCmd.AddCommand(mr_noteCmd)

	carapace.Gen(mr_noteCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_noteCmd, ""),
	)
}
