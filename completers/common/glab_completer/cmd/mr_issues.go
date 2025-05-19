package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_issuesCmd = &cobra.Command{
	Use:     "issues [<id> | <branch>]",
	Short:   "Get issues related to a particular merge request.",
	Aliases: []string{"issue"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_issuesCmd).Standalone()

	mrCmd.AddCommand(mr_issuesCmd)

	carapace.Gen(mr_issuesCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_issuesCmd, ""),
	)
}
