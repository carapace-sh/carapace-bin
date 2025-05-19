package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_approveCmd = &cobra.Command{
	Use:   "approve {<id> | <branch>}",
	Short: "Approve merge requests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_approveCmd).Standalone()

	mr_approveCmd.Flags().StringP("sha", "s", "", "SHA, which must match the SHA of the HEAD commit of the merge request.")
	mrCmd.AddCommand(mr_approveCmd)

	carapace.Gen(mr_approveCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_approveCmd, ""),
	)
}
