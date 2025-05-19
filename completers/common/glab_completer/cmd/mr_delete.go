package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_deleteCmd = &cobra.Command{
	Use:     "delete [<id> | <branch>]",
	Short:   "Delete a merge request.",
	Aliases: []string{"del"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_deleteCmd).Standalone()

	mrCmd.AddCommand(mr_deleteCmd)

	carapace.Gen(mr_deleteCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_deleteCmd, ""),
	)
}
