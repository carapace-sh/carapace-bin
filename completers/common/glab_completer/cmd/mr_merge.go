package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_mergeCmd = &cobra.Command{
	Use:     "merge {<id> | <branch>}",
	Short:   "Merge or accept a merge request.",
	Aliases: []string{"accept"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_mergeCmd).Standalone()

	mr_mergeCmd.Flags().Bool("auto-merge", false, "Set auto-merge.")
	mr_mergeCmd.Flags().StringP("message", "m", "", "Custom merge commit message.")
	mr_mergeCmd.Flags().BoolP("rebase", "r", false, "Rebase the commits onto the base branch.")
	mr_mergeCmd.Flags().BoolP("remove-source-branch", "d", false, "Remove source branch on merge.")
	mr_mergeCmd.Flags().String("sha", "", "Merge commit SHA.")
	mr_mergeCmd.Flags().BoolP("squash", "s", false, "Squash commits on merge.")
	mr_mergeCmd.Flags().String("squash-message", "", "Custom squash commit message.")
	mr_mergeCmd.Flags().Bool("when-pipeline-succeeds", false, "Merge only when pipeline succeeds")
	mr_mergeCmd.Flags().BoolP("yes", "y", false, "Skip submission confirmation prompt.")
	mr_mergeCmd.Flag("when-pipeline-succeeds").Hidden = true
	mrCmd.AddCommand(mr_mergeCmd)

	carapace.Gen(mr_mergeCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_mergeCmd, "opened"),
	)
}
