package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_closeCmd = &cobra.Command{
	Use:   "close [<id> | <branch>]",
	Short: "Close a merge request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_closeCmd).Standalone()

	mrCmd.AddCommand(mr_closeCmd)

	carapace.Gen(mr_closeCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_closeCmd, ""),
	)
}
