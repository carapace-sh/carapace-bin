package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dcommitCmd = &cobra.Command{
	Use:   "dcommit",
	Short: "Commit several diffs to merge with upstream",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dcommitCmd).Standalone()

	dcommitCmd.Flags().Bool("add-author-from", false, "Add author from")
	dcommitCmd.Flags().Bool("all", false, "Fetch from all SVN remotes")
	dcommitCmd.Flags().StringArrayP("authors-file", "A", []string{}, "Authors file")
	dcommitCmd.Flags().String("authors-prog", "", "Authors program")
	dcommitCmd.Flags().String("commit-url", "", "Commit to this SVN URL")
	dcommitCmd.Flags().String("config-dir", "", "SVN configuration directory")
	dcommitCmd.Flags().IntP("copy-similarity", "C", 0, "Copy similarity threshold")
	dcommitCmd.Flags().BoolP("dry-run", "n", false, "Dry run")
	dcommitCmd.Flags().BoolP("edit", "e", false, "Edit commit message")
	dcommitCmd.Flags().Bool("find-copies-harder", false, "Find copies harder")
	dcommitCmd.Flags().Bool("follow-parent", false, "Follow parent")
	dcommitCmd.Flags().String("ignore-paths", "", "Regex of paths to ignore")
	dcommitCmd.Flags().String("ignore-refs", "", "Regex of SVN refs to ignore")
	dcommitCmd.Flags().String("include-paths", "", "Regex of paths to include")
	dcommitCmd.Flags().Bool("interactive", false, "Interactive")
	dcommitCmd.Flags().IntS("l", "l", 0, "Rename limit")
	dcommitCmd.Flags().Bool("localtime", false, "Store Git commit times in local time zone")
	dcommitCmd.Flags().Int("log-window-size", 100, "Log window size")
	dcommitCmd.Flags().BoolP("merge", "m", false, "Merge")
	dcommitCmd.Flags().String("mergeinfo", "", "Merge information")
	dcommitCmd.Flags().Bool("no-auth-cache", false, "Disable SVN authentication caching")
	dcommitCmd.Flags().Bool("no-checkout", false, "No checkout")
	dcommitCmd.Flags().Bool("no-rebase", false, "After committing, do not rebase or reset")
	dcommitCmd.Flags().Bool("noMetadata", false, "Disable metadata")
	dcommitCmd.Flags().BoolP("quiet", "q", false, "Quiet")
	dcommitCmd.Flags().IntP("revision", "r", 0, "Revision to commit to")
	dcommitCmd.Flags().Bool("rmdir", false, "Remove empty directories")
	dcommitCmd.Flags().String("set-svn-props", "", "Set SVN properties")
	dcommitCmd.Flags().StringP("strategy", "s", "", "Rebase strategy")
	dcommitCmd.Flags().Bool("use-log-author", false, "Use log author")
	dcommitCmd.Flags().Bool("useSvmProps", false, "Use SVM properties")
	dcommitCmd.Flags().Bool("useSvnsyncProps", false, "Use SVNSync properties")
	dcommitCmd.Flags().String("username", "", "SVN username")
	dcommitCmd.Flags().BoolP("verbose", "v", false, "Verbose")
	rootCmd.AddCommand(dcommitCmd)

	carapace.Gen(dcommitCmd).FlagCompletion(carapace.ActionMap{
		"authors-file":  carapace.ActionFiles(),
		"authors-prog":  carapace.ActionFiles(),
		"commit-url":    carapace.ActionValues(),
		"config-dir":    carapace.ActionFiles(),
		"ignore-paths":  carapace.ActionValues(),
		"ignore-refs":   carapace.ActionValues(),
		"include-paths": carapace.ActionValues(),
		"mergeinfo":     carapace.ActionValues(),
		"revision":      carapace.ActionValues(),
		"set-svn-props": carapace.ActionValues(),
		"strategy":      carapace.ActionValues("recursive", "resolve", "ours", "theirs", "subtree"),
		"username":      carapace.ActionValues(),
	})

	carapace.Gen(dcommitCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
