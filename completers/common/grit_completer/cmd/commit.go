package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/grit_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:     "commit",
	Short:   "Stage every change and record a new commit",
	GroupID: "change",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commitCmd).Standalone()

	commitCmd.Flags().BoolP("all", "a", false, "Stage every change first, then commit. This is the default when `commit.autoStage` is true")
	commitCmd.Flags().Bool("allow-empty", false, "Record a commit even when the tree matches the parent (same as `git commit --allow-empty`)")
	commitCmd.Flags().Bool("amend", false, "Replace the tip commit instead of creating a new one")
	commitCmd.Flags().String("author", "", "Author for this commit (`Name <email>`); committer is unchanged")
	commitCmd.Flags().String("date", "", "Override the author date (Git date formats, RFC 3339, `@epoch`, …)")
	commitCmd.Flags().Bool("dry-run", false, "Compute the commit without writing objects, updating refs, or running hooks")
	commitCmd.Flags().BoolP("gpg-sign", "S", false, "Sign the commit with the configured signing key")
	commitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	commitCmd.Flags().StringSlice("hunks", nil, "Commit only the listed hunk ids from `grit diff --json` (comma-separated, repeatable)")
	commitCmd.Flags().StringSlice("lines", nil, "Commit only line ranges from the unstaged diff (same format as `grit add --lines`)")
	commitCmd.Flags().StringP("message", "m", "", "Commit message")
	commitCmd.Flags().StringP("message-file", "F", "", "Read the commit message from a file (`-` for stdin)")
	commitCmd.Flags().Bool("no-gpg-sign", false, "Do not sign this commit, even when `commit.gpgsign` is enabled")
	commitCmd.Flags().Bool("no-verify", false, "Skip the `pre-commit` and `commit-msg` hooks")
	commitCmd.Flags().String("patch-file", "", "Apply a unified diff to the index, then commit only that staged content (`git apply --cached` + commit)")
	commitCmd.Flags().Bool("pathspec-file-nul", false, "`--pathspec-from-file` entries are NUL-separated")
	commitCmd.Flags().String("pathspec-from-file", "", "Read pathspecs from a file (`-` for stdin) instead of or in addition to argv paths")
	commitCmd.Flags().Bool("reset-author", false, "Set the author to the committer with a fresh author date")
	commitCmd.Flags().StringP("reuse-message", "C", "", "Reuse the log message and author from a commit (`-C` / `--reuse-message`)")
	commitCmd.Flags().BoolP("signoff", "s", false, "Append a `Signed-off-by` trailer for the committer")
	commitCmd.Flags().Bool("staged", false, "Commit only what is already staged (same as `git commit` without `-a`)")
	commitCmd.Flags().StringSlice("trailer", nil, "Append a commit trailer (`key=value` or `key: value`); repeat for multiple trailers")
	rootCmd.AddCommand(commitCmd)

	carapace.Gen(commitCmd).FlagCompletion(carapace.ActionMap{
		"hunks": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionHunkIds("unstaged")
		}),
		"lines":         action.ActionLineSelections("unstaged"),
		"reuse-message": git.ActionRefs(git.RefOption{}.Default()),
	})
}
