package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:     "log",
	Short:   "Show commit logs",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(logCmd).Standalone()

	logCmd.Flags().StringArrayS("G", "G", nil, "look for differences whose patch text contains added/removed lines that match <regex>")
	logCmd.Flags().StringArrayS("L", "L", nil, "trace the evolution of the line range given by <start>,<end>, or by the function name regex <funcname>, within the <file>")
	logCmd.Flags().Bool("clear-decorations", false, "clear all previous decorate-refs options")
	logCmd.Flags().Bool("combined-all-paths", false, "list the name of the file from all parents in combined diffs for merge commits")
	logCmd.Flags().String("decorate", "", "print out the ref names of any commits that are shown")
	logCmd.Flags().StringArray("decorate-refs", nil, "decorate refs matching pattern")
	logCmd.Flags().StringArray("decorate-refs-exclude", nil, "do not decorate refs matching pattern")
	logCmd.Flags().Bool("default", false, "use argument as default revision")
	logCmd.Flags().String("diff-merges", "", "specify diff format to be used for merge commits")
	logCmd.Flags().Bool("early-output", false, "undocumented")
	logCmd.Flags().Bool("full-diff", false, "with this, the full diff is shown for commits that touch the specified paths")
	logCmd.Flags().BoolP("help", "h", false, "show help")
	logCmd.Flags().Bool("log-size", false, "include a line \"log size <number>\" in the output for each commit")
	logCmd.Flags().Bool("mailmap", false, "use mailmap file to map author and committer names and email addresses to canonical real names and email addresses")
	logCmd.Flags().Bool("no-decorate", false, "do not print out the ref names of any commits that are shown")
	logCmd.Flags().Bool("no-diff-merges", false, "disable output of diffs for merge commits")
	logCmd.Flags().Bool("no-mailmap", false, "do not use mailmap file to map author and committer names and email addresses to canonical real names and email addresses")
	logCmd.Flags().Bool("no-notes", false, "do not show notes")
	logCmd.Flags().Bool("no-relative", false, "do not show pathnames relative to the current directory")
	logCmd.Flags().Bool("no-use-mailmap", false, "do not use mailmap file to map author and committer names and email addresses to canonical real names and email addresses")
	logCmd.Flags().StringArray("notes", nil, "show the notes that annotate the commit")
	logCmd.Flags().Bool("source", false, "print out the ref name given on the command line by which each commit was reached")
	logCmd.Flags().BoolS("t", "t", false, "show the tree objects in the diff output")
	logCmd.Flags().Bool("use-mailmap", false, "use mailmap file to map author and committer names and email addresses to canonical real names and email addresses")

	common.AddBisectionHelperOptions(logCmd)
	common.AddCommitFormattingOptions(logCmd)
	common.AddCommitLimitingOptions(logCmd)
	common.AddCommitOrderingOptions(logCmd)
	common.AddDiffFlags(logCmd)
	common.AddHistorySimplificationOptions(logCmd)
	common.AddObjectTraversalOptions(logCmd)
	rootCmd.AddCommand(logCmd)

	logCmd.Flag("decorate").NoOptDefVal = "short"

	carapace.Gen(logCmd).FlagCompletion(carapace.ActionMap{
		"decorate": carapace.ActionValuesDescribed(
			"short", "do not print ref prefixes",
			"full", "print ref prefixes",
			"auto", "short format when output to terminal",
			"no", "no decoration",
		),
		"diff-merges": carapace.ActionValuesDescribed(
			"off", "Disable output of diffs for merge commits.",
			"none", "Disable output of diffs for merge commits.",
			"on", "This option makes diff output for merge commits to be shown in the default format.",
			"m", "This option makes diff output for merge commits to be shown in the default format.",
			"first-parent", "This option makes merge commits show the full diff with respect to the first parent only.",
			"1", "This option makes merge commits show the full diff with respect to the first parent only.",
			"separate", "This makes merge commits show the full diff with respect to each of the parents.",
			"remerge", "With this option, two-parent merge commits are remerged to create a temporary tree object",
			"r", "With this option, two-parent merge commits are remerged to create a temporary tree object",
			"combined", "With this option, diff output for a merge commit shows the differences from each of the parents to the merge result simultaneously.",
			"c", "With this option, diff output for a merge commit shows the differences from each of the parents to the merge result simultaneously.",
			"dense-combined", "With this option the output produced by --diff-merges=combined is further compressed.",
			"cc", "With this option the output produced by --diff-merges=combined is further compressed.",
		),
	})

	carapace.Gen(logCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			git.ActionRefRanges(git.RefOption{}.Default()).UnlessF(condition.CompletingPath),
		).ToA(),
	)

	carapace.Gen(logCmd).DashAnyCompletion(
		carapace.ActionFiles(),
	)
}
