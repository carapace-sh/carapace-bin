package cmd

import (
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use: "commit",
	Short: "Record changes to the repository",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {
	commitCmd.Flags().BoolP("all", "a", false, "commit all changed files")
	commitCmd.Flags().Bool("ahead-behind", false, "compute full ahead/behind values")
	commitCmd.Flags().Bool("amend", false, "amend previous commit")
	commitCmd.Flags().String("author", "", "override author for commit")
	commitCmd.Flags().Bool("branch", false, "show branch information")
	commitCmd.Flags().String("cleanup", "", "how to strip spaces and #comments from message")
	commitCmd.Flags().BoolP("reedit-message", "c", false, "<commit>    reuse and edit message from specified commit")
	commitCmd.Flags().BoolP("reuse-message", "C", false, "<commit>    reuse message from specified commit")
	commitCmd.Flags().String("date", "", "override date for commit")
	commitCmd.Flags().Bool("dry-run", false, "show what would be committed")
	commitCmd.Flags().BoolP("edit", "e", false, "force edit of commit")
	commitCmd.Flags().BoolP("file", "F", false, "<file>     read message from file")
	commitCmd.Flags().String("fixup", "", "use autosquash formatted message to fixup specified commit")
	commitCmd.Flags().BoolP("include", "i", false, "add specified files to index for commit")
	commitCmd.Flags().Bool("interactive", false, "interactively add files")
	commitCmd.Flags().Bool("long", false, "show status in long format (default)")
	commitCmd.Flags().BoolP("message", "m", false, "<message>    commit message")
	commitCmd.Flags().BoolP("no-verify", "n", false, "bypass pre-commit and commit-msg hooks")
	commitCmd.Flags().Bool("no-post-rewrite", false, "bypass post-rewrite hook")
	commitCmd.Flags().BoolP("only", "o", false, "commit only specified files")
	commitCmd.Flags().Bool("pathspec-file-nul", false, "with --pathspec-from-file, pathspec elements are separated with NUL character")
	commitCmd.Flags().String("pathspec-from-file", "", "read pathspec from file")
	commitCmd.Flags().Bool("porcelain", false, "machine-readable output")
	commitCmd.Flags().BoolP("patch", "p", false, "interactively add changes")
	commitCmd.Flags().BoolP("quiet", "q", false, "suppress summary after successful commit")
	commitCmd.Flags().Bool("reset-author", false, "the commit is authored by me now (used with -C/-c/--amend)")
	commitCmd.Flags().StringP("gpg-sign", "S", "", "GPG sign commit")
	commitCmd.Flags().Bool("short", false, "show status concisely")
	commitCmd.Flags().String("squash", "", "use autosquash formatted message to squash specified commit")
	commitCmd.Flags().BoolP("signoff", "s", false, "add Signed-off-by:")
	commitCmd.Flags().Bool("status", false, "include status in commit message template")
	commitCmd.Flags().BoolP("template", "t", false, "<file>    use specified template file")
	commitCmd.Flags().StringP("untracked-files", "u", "", "show untracked files, optional modes: all, normal, no. (Default: all)")
	commitCmd.Flags().BoolP("verbose", "v", false, "show diff in commit message template")
	commitCmd.Flags().BoolP("null", "z", false, "terminate entries with NUL")
    rootCmd.AddCommand(commitCmd)
}
