package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pre-commit_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var tryRepoCmd = &cobra.Command{
	Use:   "try-repo [repo] [hook]",
	Short: "Try the hooks in a repository, useful for developing new hooks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tryRepoCmd).Standalone()

	tryRepoCmd.Flags().BoolP("all-files", "a", false, "Run on all the files in the repo.")
	tryRepoCmd.Flags().String("checkout-type", "", "Indicates whether the checkout was a branch checkout (changing branches, flag=1) or a file checkout (retrieving a file from the index, flag=0).")
	tryRepoCmd.Flags().String("color", "", "Whether to use color in output.")
	tryRepoCmd.Flags().String("commit-msg-filename", "", "Filename to check when running during `commit-msg`")
	tryRepoCmd.Flags().String("commit-object-name", "", "Commit object name (typically the third argument to .git/hooks/prepare-commit-msg)")
	tryRepoCmd.Flags().StringP("config", "c", "", "Path to alternate config file")
	tryRepoCmd.Flags().Bool("fail-fast", false, "Stop after the first failing hook.")
	tryRepoCmd.Flags().String("files", "", "Specific filenames to run hooks on.")
	tryRepoCmd.Flags().String("hook-stage", "", "The stage during which the hook is fired.")
	tryRepoCmd.Flags().String("is-squash-merge", "", "During a post-merge hook, indicates whether the merge was a squash merge")
	tryRepoCmd.Flags().String("local-branch", "", "Local branch ref used by `git push`.")
	tryRepoCmd.Flags().StringP("origin", "o", "", "(for usage with `--from-ref`) -- this option represents the destination ref in a `from_ref...to_ref` diff expression.")
	tryRepoCmd.Flags().String("pre-rebase-branch", "", "The branch being rebased, and is not set when  rebasing the current branch.")
	tryRepoCmd.Flags().String("pre-rebase-upstream", "", "The upstream from which the series was forked.")
	tryRepoCmd.Flags().String("prepare-commit-message-source", "", "Source of the commit message (typically the second argument to .git/hooks/prepare-commit-msg)")
	tryRepoCmd.Flags().String("remote-branch", "", "Remote branch ref used by `git push`.")
	tryRepoCmd.Flags().String("remote-name", "", "Remote name used by `git push`.")
	tryRepoCmd.Flags().String("remote-url", "", "Remote url used by `git push`.")
	tryRepoCmd.Flags().String("rev", "", "Manually select a rev to run against, otherwise the `HEAD` revision will be used.")
	tryRepoCmd.Flags().String("rewrite-command", "", "During a post-rewrite hook, specifies the command that invoked the rewrite")
	tryRepoCmd.Flags().Bool("show-diff-on-failure", false, "When hooks fail, run `git diff` directly afterward.")
	tryRepoCmd.Flags().StringP("source", "s", "", "(for usage with `--to-ref`) -- this option represents the original ref in a `from_ref...to_ref` diff expression.")
	tryRepoCmd.Flags().BoolP("verbose", "v", false, "")

	carapace.Gen(tryRepoCmd).FlagCompletion(carapace.ActionMap{
		"color":               action.ActionColors(),
		"commit-msg-filename": carapace.ActionFiles(),
		"config":              carapace.ActionFiles(".yaml", ".yml"),
		"files":               carapace.ActionFiles(),
		"hook-stage":          action.ActionStages(),
		"rev":                 git.ActionRefs(git.RefOption{}.Default()),
	})

	carapace.Gen(tryRepoCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		action.ActionHookIDs(tryRepoCmd),
	)

	rootCmd.AddCommand(tryRepoCmd)
}
