package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_subscribeCmd = &cobra.Command{
	Use:     "subscribe [<id> | <branch>]",
	Short:   "Subscribe to a merge request.",
	Aliases: []string{"sub"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_subscribeCmd).Standalone()

	mrCmd.AddCommand(mr_subscribeCmd)

	carapace.Gen(mr_subscribeCmd).PositionalCompletion(
		action.ActionMergeRequestsAndBranches(mr_subscribeCmd, ""),
	)
}
