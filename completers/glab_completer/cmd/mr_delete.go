package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_deleteCmd = &cobra.Command{
	Use:     "delete [<id> | <branch>]",
	Short:   "Delete merge requests",
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
