package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_todoCmd = &cobra.Command{
	Use:     "todo [<id> | <branch>]",
	Short:   "Add a to-do item to merge request.",
	Aliases: []string{"add-todo"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_todoCmd).Standalone()

	mrCmd.AddCommand(mr_todoCmd)

	carapace.Gen(mr_todoCmd).PositionalCompletion(
		action.ActionMergeRequestsAndBranches(mr_todoCmd, ""),
	)
}
