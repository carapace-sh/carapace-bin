package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pre-commit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [hook]",
	Short: "Run hooks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().BoolP("all-files", "a", false, "Run on all the files in the repo.")
	runCmd.Flags().String("checkout-type", "", "Indicates whether the checkout was a branch checkout (changing branches, flag=1) or a file checkout (retrieving a file from the index, flag=0).")
	runCmd.Flags().String("color", "", "Whether to use color in output.")
	runCmd.Flags().String("commit-msg-filename", "", "Filename to check when running during `commit-msg`")
	runCmd.Flags().String("commit-object-name", "", "Commit object name (typically the third argument to .git/hooks/prepare-commit-msg)")
	runCmd.Flags().StringP("config", "c", "", "Path to alternate config file")
	runCmd.Flags().Bool("fail-fast", false, "Stop after the first failing hook.")
	runCmd.Flags().String("files", "", "Specific filenames to run hooks on.")
	runCmd.Flags().String("hook-stage", "", "The stage during which the hook is fired.")
	runCmd.Flags().String("is-squash-merge", "", "During a post-merge hook, indicates whether the merge was a squash merge")
	runCmd.Flags().String("local-branch", "", "Local branch ref used by `git push`.")
	runCmd.Flags().StringP("origin", "o", "", "(for usage with `--from-ref`) -- this option represents the destination ref in a `from_ref...to_ref` diff expression.")
	runCmd.Flags().String("pre-rebase-branch", "", "The branch being rebased, and is not set when  rebasing the current branch.")
	runCmd.Flags().String("pre-rebase-upstream", "", "The upstream from which the series was forked.")
	runCmd.Flags().String("prepare-commit-message-source", "", "Source of the commit message (typically the second argument to .git/hooks/prepare-commit-msg)")
	runCmd.Flags().String("remote-branch", "", "Remote branch ref used by `git push`.")
	runCmd.Flags().String("remote-name", "", "Remote name used by `git push`.")
	runCmd.Flags().String("remote-url", "", "Remote url used by `git push`.")
	runCmd.Flags().String("rewrite-command", "", "During a post-rewrite hook, specifies the command that invoked the rewrite")
	runCmd.Flags().Bool("show-diff-on-failure", false, "When hooks fail, run `git diff` directly afterward.")
	runCmd.Flags().StringP("source", "s", "", "(for usage with `--to-ref`) -- this option represents the original ref in a `from_ref...to_ref` diff expression.")
	runCmd.Flags().BoolP("verbose", "v", false, "")

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"color":               action.ActionColors(),
		"commit-msg-filename": carapace.ActionFiles(),
		"config":              carapace.ActionFiles(".yaml", ".yml"),
		"files":               carapace.ActionFiles(),
		"hook-stage":          action.ActionStages(),
	})

	carapace.Gen(runCmd).PositionalCompletion(
		action.ActionHookIDs(runCmd),
	)

	rootCmd.AddCommand(runCmd)
}
