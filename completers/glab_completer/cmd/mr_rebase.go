package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_rebaseCmd = &cobra.Command{
	Use:   "rebase [<id> | <branch>] [flags]",
	Short: "Automatically rebase the source_branch of the merge request against its target_branch.",
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
