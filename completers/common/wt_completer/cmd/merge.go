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

	mergeCmd.Flags().Bool("commit", false, "Force commit and squash")
	mergeCmd.Flags().Bool("ff", false, "Allow fast-forward (default)")
	mergeCmd.Flags().String("format", "", "Output format (text, json)")
	mergeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	mergeCmd.Flags().Bool("no-commit", false, "Skip commit and squash")
	mergeCmd.Flags().Bool("no-ff", false, "Create a merge commit (no fast-forward)")
	mergeCmd.Flags().Bool("no-hooks", false, "Skip hooks")
	mergeCmd.Flags().Bool("no-rebase", false, "Skip rebase; require the target to fast-forward to the resulting tip")
	mergeCmd.Flags().Bool("no-remove", false, "Keep worktree after merge")
	mergeCmd.Flags().Bool("no-squash", false, "Skip commit squashing")
	mergeCmd.Flags().Bool("no-verify", false, "Skip hooks (deprecated alias for --no-hooks)")
	mergeCmd.Flags().Bool("rebase", false, "Force rebasing onto target")
	mergeCmd.Flags().Bool("remove", false, "Force worktree removal after merge")
	mergeCmd.Flags().String("stage", "", "What to stage before committing [default: all]")
	mergeCmd.Flags().Bool("verify", false, "Force running hooks")
	mergeCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	rootCmd.AddCommand(mergeCmd)

	carapace.Gen(mergeCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("text", "json"),
		"stage": carapace.ActionValuesDescribed(
			"all", "Stage everything: untracked files + unstaged tracked changes",
			"tracked", "Stage tracked changes only (like git add -u)",
			"none", "Stage nothing, commit only what's already in the index",
		).StyleF(style.ForKeyword),
	})

	carapace.Gen(mergeCmd).PositionalCompletion(
		wt.ActionBranches(),
	)
}
