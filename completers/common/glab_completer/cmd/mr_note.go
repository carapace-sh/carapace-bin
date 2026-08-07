package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_noteCmd = &cobra.Command{
	Use:     "note [<id> | <branch>]",
	Short:   "Manage comments and discussions on a merge request.",
	Aliases: []string{"comment"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_noteCmd).Standalone()

	mr_noteCmd.Flags().StringP("message", "m", "", "Comment or note message.")
	mr_noteCmd.Flags().String("resolve", "", "Resolve the discussion containing the specified note ID.")
	mr_noteCmd.Flags().Bool("unique", false, "Do not create a comment if a comment with the same body already exists.")
	mr_noteCmd.Flags().String("unresolve", "", "Unresolve the discussion containing the specified note ID.")
	mr_noteCmd.Flag("message").Hidden = true
	mr_noteCmd.Flag("resolve").Hidden = true
	mr_noteCmd.Flag("unique").Hidden = true
	mr_noteCmd.Flag("unresolve").Hidden = true
	mrCmd.AddCommand(mr_noteCmd)

	carapace.Gen(mr_noteCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_noteCmd, ""),
	)
}
