package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_revokeCmd = &cobra.Command{
	Use:     "revoke [<id> | <branch>]",
	Short:   "Revoke approval on a merge request.",
	Aliases: []string{"unapprove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_revokeCmd).Standalone()

	mrCmd.AddCommand(mr_revokeCmd)

	carapace.Gen(mr_revokeCmd).PositionalCompletion(
		action.ActionMergeRequestsAndBranches(mr_revokeCmd, ""),
	)
}
