package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_reopenCmd = &cobra.Command{
	Use:     "reopen [<id>... | <branch>...]",
	Short:   "Reopen a merge request.",
	Aliases: []string{"open"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_reopenCmd).Standalone()

	mrCmd.AddCommand(mr_reopenCmd)

	carapace.Gen(mr_reopenCmd).PositionalCompletion(
		action.ActionMergeRequestsAndBranches(mr_reopenCmd, "closed"),
	)
}
