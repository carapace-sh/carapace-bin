package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_rebaseCmd = &cobra.Command{
	Use:   "rebase [<id> | <branch>] [flags]",
	Short: "Rebase the source branch of a merge request against its target branch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_rebaseCmd).Standalone()

	mr_rebaseCmd.Flags().Bool("skip-ci", false, "Rebase merge request while skipping CI/CD pipeline.")
	mrCmd.AddCommand(mr_rebaseCmd)

	carapace.Gen(mr_rebaseCmd).PositionalCompletion(
		action.ActionMergeRequestsAndBranches(mr_rebaseCmd, ""),
	)
}
