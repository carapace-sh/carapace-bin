package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge current branch into target",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mergeCmd).Standalone()

	mergeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	mergeCmd.Flags().Bool("no-commit", false, "Skip commit and squash")
	mergeCmd.Flags().Bool("no-ff", false, "Create a merge commit (no fast-forward)")
	mergeCmd.Flags().Bool("no-rebase", false, "Skip rebase (fail if not already rebased)")
	mergeCmd.Flags().Bool("no-remove", false, "Keep worktree after merge")
	mergeCmd.Flags().Bool("no-squash", false, "Skip commit squashing")
	mergeCmd.Flags().Bool("no-verify", false, "Skip hooks")
	mergeCmd.Flags().String("stage", "", "What to stage before committing [default: all]")
	mergeCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	rootCmd.AddCommand(mergeCmd)

	carapace.Gen(mergeCmd).FlagCompletion(carapace.ActionMap{
		"stage": carapace.ActionValues("all", "tracked", "none").StyleF(style.ForKeyword),
	})

	carapace.Gen(mergeCmd).PositionalCompletion(
		wt.ActionBranches(), // TODO what are valid targets? (likely only local worktrees)
	)
}
