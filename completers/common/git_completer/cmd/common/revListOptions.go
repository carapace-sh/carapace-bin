package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

// AddCommitLimitingOptions adds commit limiting options.
//
//	https://git-scm.com/docs/git-rev-list#_commit_limiting
func AddCommitLimitingOptions(cmd *cobra.Command) {
	cmd.Flags().String("after", "", "show commits more recent than a specific date")
	cmd.Flags().Bool("all", false, "pretend as if all the refs in refs/ are listed on the command line")
	cmd.Flags().Bool("all-match", false, "limit output to ones matching all given --grep")
	cmd.Flags().Bool("alternate-refs", false, "pretend as if all alternate ref tips are listed on the command line")
	cmd.Flags().StringArray("author", nil, "limit output to ones with matching author header")
	cmd.Flags().Bool("basic-regexp", false, "consider the limiting patterns to be basic regular expressions")
	cmd.Flags().String("before", "", "show commits older than a specific date")
	cmd.Flags().Bool("boundary", false, "output excluded boundary commits")
	cmd.Flags().StringArray("branches", nil, "pretend as if all the refs in refs/heads are listed on the command line")
	cmd.Flags().Bool("cherry", false, "a synonym for --right-only --cherry-mark --no-merges")
	cmd.Flags().Bool("cherry-mark", false, "like --cherry-pick but mark commits rather than omitting them")
	cmd.Flags().Bool("cherry-pick", false, "omit any commit that introduces the same change as another commit")
	cmd.Flags().StringArray("committer", nil, "limit output to ones with matching committer header")
	cmd.Flags().String("disk-usage", "", "suppress normal output and print disk usage")
	cmd.Flags().StringArray("exclude", nil, "do not include matching refs the next option would consider")
	cmd.Flags().Bool("exclude-first-parent-only", false, "for excluded commits follow only the first parent commit upon merge")
	cmd.Flags().String("exclude-hidden", "", "do not include refs that would be hidden")
	cmd.Flags().BoolP("extended-regexp", "E", false, "consider the limiting patterns to be extended regular expressions")
	cmd.Flags().Bool("first-parent", false, "for included commits follow only the first parent commit upon merge")
	cmd.Flags().BoolP("fixed-strings", "F", false, "consider the limiting patterns to be fixed strings")
	cmd.Flags().StringArray("glob", nil, "pretend as if all matching refs are listed on the command line")
	cmd.Flags().StringArray("grep", nil, "limit output to ones with matching log message")
	cmd.Flags().StringArray("grep-reflog", nil, "limit output to ones with matching reflog entries")
	cmd.Flags().Bool("ignore-missing", false, "pretend as if the bad input was not given")
	cmd.Flags().Bool("invert-grep", false, "limit output to ones with non-matching log message")
	cmd.Flags().Bool("left-only", false, "list only commits on the left side of a symmetric difference")
	cmd.Flags().String("max-age", "", "limit the commits output to specified time range")
	cmd.Flags().StringP("max-count", "n", "", "limit the number of commits to output")
	cmd.Flags().String("max-count-oldest", "", "limit the number of commits to output, picking the oldest N")
	cmd.Flags().String("max-parents", "", "show only commits which have at most that many parent commits")
	cmd.Flags().Bool("merge", false, "show refs that touch files having a conflict and don't exist on all heads to merge")
	cmd.Flags().Bool("merges", false, "print only merge commits")
	cmd.Flags().Bool("maximal-only", false, "restrict output to commits not reachable from other commits in the range")
	cmd.Flags().String("min-age", "", "limit the commits output to specified time range")
	cmd.Flags().String("min-parents", "", "show only commits which have at least that many parent commits")
	cmd.Flags().Bool("no-max-parents", false, "reset max-parents limit")
	cmd.Flags().Bool("no-merges", false, "do not print commits with more than one parent")
	cmd.Flags().Bool("no-min-parents", false, "reset min-parents limit")
	cmd.Flags().Bool("not", false, "reverses the meaning of the ^ prefix")
	cmd.Flags().BoolP("perl-regexp", "P", false, "consider the limiting patterns to be Perl-compatible regular expressions")
	cmd.Flags().String("progress", "", "show progress reports on stderr as objects are considered")
	cmd.Flags().Bool("quiet", false, "don't print anything to standard output")
	cmd.Flags().Bool("reflog", false, "pretend as if all reflog objects are listed on the command line")
	cmd.Flags().BoolP("regexp-ignore-case", "i", false, "match the regular expression case insensitive")
	cmd.Flags().StringArray("remotes", nil, "pretend as if all the refs in refs/remotes are listed on the command line")
	cmd.Flags().Bool("remove-empty", false, "stop when a given path disappears from the tree")
	cmd.Flags().Bool("right-only", false, "list only commits on the right side of a symmetric difference")
	cmd.Flags().String("since", "", "show commits more recent than a specific date")
	cmd.Flags().String("since-as-filter", "", "show all commits more recent than a specific date")
	cmd.Flags().Bool("single-worktree", false, "examine the current working tree only")
	cmd.Flags().String("skip", "", "skip number commits before starting to show the commit output")
	cmd.Flags().Bool("stdin", false, "read arguments from standard input as well")
	cmd.Flags().StringArray("tags", nil, "pretend as if all the refs in refs/tags are listed on the command line")
	cmd.Flags().String("until", "", "show commits older than a specific date")
	cmd.Flags().Bool("use-bitmap-index", false, "try to speed up the traversal using the pack bitmap index")
	cmd.Flags().BoolP("walk-reflogs", "g", false, "walk reflog entries from the most recent one to older ones")

	cmd.Flag("branches").NoOptDefVal = " "
	cmd.Flag("disk-usage").NoOptDefVal = " "
	cmd.Flag("exclude-hidden").NoOptDefVal = " "
	cmd.Flag("remotes").NoOptDefVal = " "
	cmd.Flag("tags").NoOptDefVal = " "

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"after":      time.ActionDate(),
		"author":     git.ActionAuthors(),
		"before":     time.ActionDate(),
		"branches":   git.ActionLocalBranches(),
		"committer":  git.ActionCommitters(),
		"disk-usage": carapace.ActionValues("human"),
		"exclude": carapace.Batch(
			git.ActionLocalBranches(),
			git.ActionTags(),
			git.ActionRemotes(),
		).ToA(),
		"exclude-hidden":  carapace.ActionValues("receive", "uploadpack"),
		"glob":            carapace.ActionValues(), // TODO
		"grep":            carapace.ActionValues(), // TODO
		"grep-reflog":     carapace.ActionValues(), // TODO
		"max-age":         carapace.ActionValues(), // TODO
		"max-count":       carapace.ActionValues(), // TODO
		"max-parents":     carapace.ActionValues(), // TODO
		"min-age":         carapace.ActionValues(), // TODO
		"min-parents":     carapace.ActionValues(), // TODO
		"progress":        carapace.ActionValues(), // TODO
		"remotes":         git.ActionRemotes(),
		"since":           time.ActionDate(),
		"since-as-filter": time.ActionDate(),
		"skip":            carapace.ActionValues(), // TODO
		"tags":            git.ActionTags(),
		"until":           time.ActionDate(),
	})
}

// AddHistorySimplificationOptions adds history simplification options.
//
//	https://git-scm.com/docs/git-rev-list#_history_simplification
func AddHistorySimplificationOptions(cmd *cobra.Command) {
	cmd.Flags().String("ancestry-path", "", "only display commits in that range that are ancestors of <commit>, descendants of <commit>, or <commit> itself")
	cmd.Flags().Bool("dense", false, "only the selected commits are shown, plus some to have a meaningful history")
	cmd.Flags().Bool("full-history", false, "same as the default mode, but does not prune some history")
	cmd.Flags().Bool("show-pulls", false, "include also any merge commits that are not TREESAME to the first parent but are TREESAME to a later parent")
	cmd.Flags().Bool("simplify-by-decoration", false, "commits that are referred by some branch or tag are selected")
	cmd.Flags().Bool("simplify-merges", false, "additional option to --full-history to remove some needless merges from the resulting history")
	cmd.Flags().Bool("sparse", false, "all commits in the simplified history are shown")

	cmd.Flag("ancestry-path").NoOptDefVal = " "

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"ancestry-path": git.ActionRefs(git.RefOption{}.Default()),
	})
}

// AddBisectionHelperOptions completes bisection helper options.
//
//	https://git-scm.com/docs/git-rev-list#_bisection_helpers
func AddBisectionHelperOptions(cmd *cobra.Command) {
	cmd.Flags().Bool("bisect", false, "limit output to the one commit object which is roughly halfway between included and excluded commits")
	cmd.Flags().Bool("bisect-all", false, "output all the commit objects between the included and excluded commits")
	cmd.Flags().Bool("bisect-vars", false, "calculate the same as --bisect, except that refs in refs/bisect/ are not used")
}

// AddCommitOrderingOptions adds commit ordering options.
//
//	https://git-scm.com/docs/git-rev-list#_commit_ordering
func AddCommitOrderingOptions(cmd *cobra.Command) {
	cmd.Flags().Bool("author-date-order", false, "show commits in the author timestamp order")
	cmd.Flags().Bool("date-order", false, "show commits in the commit timestamp order")
	cmd.Flags().Bool("reverse", false, "output the commits in reverse order")
	cmd.Flags().Bool("topo-order", false, "avoid showing commits on multiple lines of history intermixed")
}

// AddCommitFormattingOptions adds commit formatting options.
//
//	https://git-scm.com/docs/git-rev-list#_commit_formatting
func AddCommitFormattingOptions(cmd *cobra.Command) {
	cmd.Flags().Bool("abbrev-commit", false, "show a prefix that names the object uniquely")
	cmd.Flags().Bool("children", false, "print also the children of the commit")
	cmd.Flags().Bool("commit-header", false, "overrides a previous --no-commit-header")
	cmd.Flags().Bool("count", false, "only print a number stating how many commits would have been listed")
	cmd.Flags().String("date", "", "date format")
	cmd.Flags().String("encoding", "", "re-code the commit log message in given encoding")
	cmd.Flags().String("expand-tabs", "", "perform a tab expansion")
	cmd.Flags().String("format", "", "pretty-print the contents of the commit logs in a given format")
	cmd.Flags().Bool("graph", false, "draw a text-based graphical representation of the commit history")
	cmd.Flags().String("graph-lane-limit", "", "limit the number of lanes in graph output")
	cmd.Flags().Bool("header", false, "print the contents of the commit in raw-format")
	cmd.Flags().Bool("left-right", false, "mark which side of a symmetric difference a commit is reachable from")
	cmd.Flags().Bool("no-abbrev-commit", false, "show the full 40-byte hexadecimal commit object name")
	cmd.Flags().Bool("no-commit-header", false, "suppress the header line containing \"commit\"")
	cmd.Flags().Bool("no-expand-tabs", false, "do not perform a tab expansion")
	cmd.Flags().Bool("oneline", false, "shorthand for \"--pretty=oneline --abbrev-commit\" used together")
	cmd.Flags().Bool("parents", false, "print also the parents of the commit")
	cmd.Flags().String("pretty", "", "pretty-print the contents of the commit logs in a given format")
	cmd.Flags().Bool("relative-date", false, "synonym for --date=relative")
	cmd.Flags().String("show-linear-break", "", "put a barrier in between branches in graph output")
	cmd.Flags().Bool("show-signature", false, "check the validity of a signed commit object")
	cmd.Flags().Bool("timestamp", false, "print the raw commit timestamp")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"date": carapace.ActionValuesDescribed(
			"relative", "shows dates relative to the current time",
			"local", "is an alias for --date=default-local",
			"iso", "shows timestamps in a ISO 8601-like format",
			"iso8601", "shows timestamps in a ISO 8601-like format",
			"iso-strict", "shows timestamps in strict ISO 8601 format",
			"iso8601-strict", "shows timestamps in strict ISO 8601 format",
			"rfc", "shows timestamps in RFC 2822 format",
			"rfc2822", "shows timestamps in RFC 2822 format",
			"short", "shows only the date, but not the time, in YYYY-MM-DD format",
			"raw", "shows the date as seconds since the epoch (1970-01-01 00:00:00 UTC)",
			"human", "shows the date in human readable format",
			"unix", "shows the date as a Unix epoch timestamp (seconds since 1970)",
			"format:", "feeds the given format to your system strftime",
			"format:%c", "show the date in your system locale’s preferred format",
			"default", "the default format",
		),
		"encoding":          carapace.ActionValues(), // TODO
		"expand-tabs":       carapace.ActionValues(), // TODO
		"format":            carapace.ActionValues("oneline", "short", "medium", "full", "fuller", "reference", "email", "raw", "format:"),
		"pretty":            carapace.ActionValues("oneline", "short", "medium", "full", "fuller", "reference", "email", "raw", "format:"),
		"show-linear-break": carapace.ActionValues(), // TODO
	})
}

// AddObjectTraversalOptions adds object traversal options.
//
//	https://git-scm.com/docs/git-rev-list#_object_traversal
func AddObjectTraversalOptions(cmd *cobra.Command) {
	cmd.Flags().Bool("do-walk", false, "overrides a previous --no-walk")
	cmd.Flags().Bool("exclude-promisor-objects", false, "prefilter object traversal at promisor boundary")
	cmd.Flags().String("filter", "", "omit objects from the list of printed objects")
	cmd.Flags().Bool("filter-print-omitted", false, "print a list of the objects omitted by the filter")
	cmd.Flags().Bool("filter-provided-objects", false, "filter the list of explicitly provided objects")
	cmd.Flags().Bool("in-commit-order", false, "print tree and blob ids in order of the commits")
	cmd.Flags().Bool("indexed-objects", false, "pretend as if all trees and blobs used by the index are listed on the command line")
	cmd.Flags().String("missing", "", "specify how missing objects are handled")
	cmd.Flags().Bool("no-filter", false, "turn off any previous --filter= argument")
	cmd.Flags().Bool("no-object-names", false, "do not print the names of the object IDs that are found")
	cmd.Flags().String("no-walk", "", "only show the given commits, but do not traverse their ancestors")
	cmd.Flags().Bool("object-names", false, "print the names of the object IDs that are found")
	cmd.Flags().Bool("objects", false, "print the object IDs of any object referenced by the listed commits")
	cmd.Flags().Bool("objects-edge", false, "similar to --objects, but also print the IDs of excluded commits")
	cmd.Flags().Bool("objects-edge-aggressive", false, "similar to --objects-edge, but tries harder to find excluded commits")
	cmd.Flags().Bool("unpacked", false, "print the object IDs that are not in packs")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionValues(), // TODO
		"missing": carapace.ActionValues(
			"error", "requests that rev-list stop with an error if a missing object is encountered",
			"allow-any", "allow object traversal to continue if a missing object is encountered",
			"allow-promisor", "is like allow-any, but will only allow object traversal to continue for EXPECTED promisor missing objects",
			"print", "is like allow-any, but will also print a list of the missing objects",
		),
		"no-walk": carapace.ActionValues("sorted", "unsorted"),
	})
}
