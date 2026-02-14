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

	logCmd.Flags().StringS("G", "G", "", "Look for differences whose patch text contains added/removed lines that match <regex>.")
	logCmd.Flags().StringS("L", "L", "", "Trace the evolution of the line range given by <start>,<end>, or by the function name regex <funcname>, within the <file>.")
	logCmd.Flags().Bool("ancestry-path", false, "When given a range of commits to display, only display commits that exist directly on the ancestry chain.")
	logCmd.Flags().Bool("author-date-order", false, "Show no parents before all of its children are shown, but otherwise show commits in the author timestamp order.")
	logCmd.Flags().Bool("bisect", false, "Pretend as if the bad bisection ref refs/bisect/bad was listed and as if it was followed by --not and the good bisection refs refs/bisect/good-* on the command line.")
	logCmd.Flags().Bool("combined-all-paths", false, "This flag causes combined diffs (used for merge commits) to list the name of the file from all parents.")
	logCmd.Flags().Bool("date-order", false, "Show no parents before all of its children are shown, but otherwise show commits in the commit timestamp order.")
	logCmd.Flags().String("decorate", "", "Print out the ref names of any commits that are shown.")
	logCmd.Flags().String("decorate-refs", "", "Decorate refs matching pattern.")
	logCmd.Flags().String("decorate-refs-exclude", "", "Do not decorete refs matching pattern.")
	logCmd.Flags().Bool("default", false, "use argument as default revision")
	logCmd.Flags().Bool("dense", false, "Only the selected commits are shown")
	logCmd.Flags().String("diff-merges", "", "Specify diff format to be used for merge commits.")
	logCmd.Flags().Bool("early-output", false, "undocumented")
	logCmd.Flags().Bool("full-diff", false, "With this, the full diff is shown for commits that touch the specified paths")
	logCmd.Flags().Bool("full-history", false, "Same as the default mode, but does not prune some history.")
	logCmd.Flags().BoolP("help", "h", false, "Show help.")
	logCmd.Flags().Bool("log-size", false, "Include a line “log size <number>” in the output for each commit.")
	logCmd.Flags().Bool("mailmap", false, "Use mailmap file to map author and committer names and email addresses to canonical real names and email addresses.")
	logCmd.Flags().Bool("no-decorate", false, "Print out the ref names of any commits that are shown.")
	logCmd.Flags().Bool("no-diff-merges", false, "Specify diff format to be used for merge commits.")
	logCmd.Flags().Bool("no-mailmap", false, "Do not use mailmap file to map author and committer names and email addresses to canonical real names and email addresses.")
	logCmd.Flags().Bool("no-notes", false, "Do not show notes.")
	logCmd.Flags().Bool("no-relative", false, "When run from a subdirectory of the project, it can be told to exclude changes outside the directory and show pathnames relative to it with this option.")
	logCmd.Flags().Bool("no-use-mailmap", false, "Do not use mailmap file to map author and committer names and email addresses to canonical real names and email addresses.")
	logCmd.Flags().String("notes", "", "Show the notes that annotate the commit")
	logCmd.Flags().Bool("reverse", false, "Output the commits chosen to be shown (see Commit Limiting section above) in reverse order. Cannot be combined with --walk-reflogs.")
	logCmd.Flags().String("rotate-to", "", "Discard the files before the named <file> from the output (i.e.  skip to), or move them to the end of the output (i.e.  rotate to).")
	logCmd.Flags().Bool("show-pulls", false, "Include all commits from the default mode, but also any merge commits that are not TREESAME to the first parent but are TREESAME to a later parent.")
	logCmd.Flags().Bool("simplify-by-decoration", false, "Commits that are referred by some branch or tag are selected.")
	logCmd.Flags().Bool("simplify-merges", false, "First, build a history graph in the same way that --full-history with parent rewriting does.")
	logCmd.Flags().String("skip-to", "", "Discard the files before the named <file> from the output (i.e.  skip to), or move them to the end of the output (i.e.  rotate to).")
	logCmd.Flags().Bool("source", false, "Print out the ref name given on the command line by which each commit was reached.")
	logCmd.Flags().Bool("sparse", false, "All commits that are walked are included.")
	logCmd.Flags().BoolS("t", "t", false, "Show the tree objects in the diff output.")
	logCmd.Flags().Bool("topo-order", false, "Show no parents before all of its children are shown, and avoid showing commits on multiple lines of history intermixed.")
	logCmd.Flags().Bool("use-mailmap", false, "Use mailmap file to map author and committer names and email addresses to canonical real names and email addresses.")
	common.AddCommitFormattingOptions(logCmd)
	common.AddCommitLimitingOptions(logCmd)
	common.AddDiffFlags(logCmd)
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
