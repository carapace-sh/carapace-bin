package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_checkoutCmd = &cobra.Command{
	Use:   "checkout [<id> | <branch> | <url>]",
	Short: "Check out an open merge request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_checkoutCmd).Standalone()

	mr_checkoutCmd.Flags().StringP("branch", "b", "", "Check out merge request with name <branch>.")
	mr_checkoutCmd.Flags().StringP("set-upstream-to", "u", "", "Set tracking of checked-out branch to [REMOTE/]BRANCH.")
	mr_checkoutCmd.Flags().BoolP("track", "t", false, "Set checked out branch to track the remote branch.")
	mr_checkoutCmd.Flag("track").Hidden = true
	mrCmd.AddCommand(mr_checkoutCmd)

	carapace.Gen(mr_checkoutCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(mr_checkoutCmd),
		// TODO "set-upstream-to":
	})

	carapace.Gen(mr_checkoutCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_checkoutCmd, ""),
	)
}
