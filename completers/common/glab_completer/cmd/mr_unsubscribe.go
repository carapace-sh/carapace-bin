package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_unsubscribeCmd = &cobra.Command{
	Use:     "unsubscribe [<id> | <branch>]",
	Short:   "Unsubscribe from a merge request.",
	Aliases: []string{"unsub"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_unsubscribeCmd).Standalone()

	mrCmd.AddCommand(mr_unsubscribeCmd)

	carapace.Gen(mr_unsubscribeCmd).PositionalCompletion(
		action.ActionMergeRequestsAndBranches(mr_unsubscribeCmd, ""),
	)
}
