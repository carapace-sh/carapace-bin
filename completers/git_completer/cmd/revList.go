package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var revListCmd = &cobra.Command{
	Use:     "rev-list",
	Short:   "Lists commit objects in reverse chronological order",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(revListCmd).Standalone()

	// TODO a lot of these flags are repeatable
	revListCmd.Flags().Bool("abbrev-commit", false, "Instead of showing the full 40-byte hexadecimal commit object name")
	revListCmd.Flags().String("after", "", "Show commits more recent than a specific date")
	revListCmd.Flags().Bool("all", false, "Pretend as if all the refs in refs/ are listed on the command line as <commit>")
	revListCmd.Flags().Bool("all-match", false, "Limit the commits output to ones that match all given --grep")
	revListCmd.Flags().Bool("alternate-refs", false, "Pretend as if all objects mentioned as ref tips of alternate repositories were listed on the command line")
	revListCmd.Flags().String("ancestry-path", "", "Limit the displayed commits to those which are an ancestor of <commit>")
	revListCmd.Flags().String("author", "", "Limit the commits output to ones with author")
	revListCmd.Flags().Bool("author-date-order", false, "Show no parents before all of its children are shown")
	revListCmd.Flags().Bool("basic-regexp", false, "Consider the limiting patterns to be basic regular expressions")
	revListCmd.Flags().String("before", "", "Show commits older than a specific date")
	revListCmd.Flags().Bool("bisect", false, "Limit output to the one commit object which is roughly halfway between included and excluded commits")
	revListCmd.Flags().Bool("bisect-all", false, "This outputs all the commit objects between the included and excluded commits")
	revListCmd.Flags().Bool("bisect-vars", false, "This calculates the same as --bisect, except that refs in refs/bisect/ are not used")
	revListCmd.Flags().Bool("boundary", false, "Output excluded boundary commits")
	revListCmd.Flags().String("branches", "", "Pretend as if all the refs in refs/heads are listed on the command line as <commit>")
	revListCmd.Flags().Bool("cherry", false, "A synonym for --right-only --cherry-mark --no-merges")
	revListCmd.Flags().Bool("cherry-mark", false, "Like --cherry-pick but mark equivalent commits with = rather than omitting them")
	revListCmd.Flags().Bool("cherry-pick", false, "Omit any commit that introduces the same change as another commit")
	revListCmd.Flags().Bool("children", false, "Print also the children of the commit")
	revListCmd.Flags().Bool("commit-header", false, "Overrides a previous --no-commit-header")
	revListCmd.Flags().String("committer", "", "Limit the commits output to ones with committer")
	revListCmd.Flags().Bool("count", false, "Print a number stating how many commits would have been listed")
	revListCmd.Flags().String("date", "", "Only takes effect for dates shown in human-readable format")
	revListCmd.Flags().Bool("date-order", false, "Show no parents before all of its children are shown")
	revListCmd.Flags().Bool("dense", false, "Commits that are walked are included if they are not TREESAME to any parent")
	revListCmd.Flags().String("disk-usage", "", "Suppress normal output")
	revListCmd.Flags().Bool("do-walk", false, "Overrides a previous --no-walk")
	revListCmd.Flags().String("encoding", "", "Commit objects record the character encoding used for the log message in their encoding header")
	revListCmd.Flags().String("exclude", "", "Do not include refs matching <glob-pattern>")
	revListCmd.Flags().String("exclude-first-parent-only", "", "Follow only the first parent commit upon seeing a merge commit")
	revListCmd.Flags().String("exclude-hidden", "", "Do not include refs that would be hidden by git-receive-pack or git-upload-pack")
	revListCmd.Flags().Bool("exclude-promisor-objects", false, "Prefilter object traversal at promisor boundary")
	revListCmd.Flags().String("expand-tabs", "", "Perform a tab expansion")
	revListCmd.Flags().BoolP("extended-regexp", "E", false, "Consider the limiting patterns to be extended regular expressions")
	revListCmd.Flags().String("filter", "", "Omits objects from the list of printed objects")
	revListCmd.Flags().Bool("filter-print-omitted", false, "Prints a list of the objects omitted by the filter")
	revListCmd.Flags().Bool("filter-provided-objects", false, "Filter the list of explicitly provided objects")
	revListCmd.Flags().Bool("first-parent", false, "Follow only the first parent commit upon seeing a merge commit")
	revListCmd.Flags().BoolP("fixed-strings", "F", false, "Consider the limiting patterns to be fixed strings")
	revListCmd.Flags().String("format", "", "Pretty-print the contents of the commit logs in a given format")
	revListCmd.Flags().Bool("full-history", false, "Same as the default mode, but does not prune some history")
	revListCmd.Flags().String("glob", "", "Pretend as if all the refs matching shell glob <glob-pattern> are listed on the command line as <commit>")
	revListCmd.Flags().Bool("graph", false, "Draw a text-based graphical representation of the commit history")
	revListCmd.Flags().String("grep", "", "Limit the commits output to ones with log message that matches the specified pattern")
	revListCmd.Flags().String("grep-reflog", "", "Limit the commits output to ones with reflog entries that match the specified pattern")
	revListCmd.Flags().Bool("header", false, "Print the contents of the commit in raw-format")
	revListCmd.Flags().Bool("ignore-missing", false, "Ignore invalid object names in the input")
	revListCmd.Flags().Bool("in-commit-order", false, "Print tree and blob ids in order of the commits")
	revListCmd.Flags().Bool("indexed-objects", false, "Pretend as if all trees and blobs used by the index are listed on the command line")
	revListCmd.Flags().Bool("invert-grep", false, "Limit the commits output to ones with log message that do not match the pattern specified with --grep")
	revListCmd.Flags().Bool("left-only", false, "List only commits on the left side of a symmetric difference")
	revListCmd.Flags().Bool("left-right", false, "Mark which side of a symmetric difference a commit is reachable from")
	revListCmd.Flags().String("max-age", "", "Limit the commits output to specified time range")
	revListCmd.Flags().StringP("max-count", "n", "", "Limit the number of commits to output")
	revListCmd.Flags().String("max-parents", "", "Show only commits which have at most that many parent commits")
	revListCmd.Flags().Bool("merge", false, "Show refs that touch files having a conflict and don’t exist on all heads to merge")
	revListCmd.Flags().Bool("merges", false, "Print only merge commits")
	revListCmd.Flags().String("min-age", "", "Limit the commits output to specified time range")
	revListCmd.Flags().String("min-parents", "", "Show only commits which have at least that many parent commits")
	revListCmd.Flags().String("missing", "", "A debug option to help with future \"partial clone\" development")
	revListCmd.Flags().Bool("no-abbrev-commit", false, "Show the full 40-byte hexadecimal commit object name")
	revListCmd.Flags().Bool("no-commit-header", false, "Suppress the header line containing \"commit\" and the object ID printed before the specified format")
	revListCmd.Flags().Bool("no-expand-tabs", false, "Do not perform a tab expansion")
	revListCmd.Flags().Bool("no-filter", false, "Turn off any previous --filter= argument")
	revListCmd.Flags().Bool("no-max-parents", false, "Reset maximum parents limit")
	revListCmd.Flags().Bool("no-merges", false, "Do not print commits with more than one parent")
	revListCmd.Flags().Bool("no-min-parents", false, "Reset minimum parents limit")
	revListCmd.Flags().Bool("no-object-names", false, "Does not print the names of the object IDs that are found")
	revListCmd.Flags().String("no-walk", "", "Only show the given commits, but do not traverse their ancestors")
	revListCmd.Flags().Bool("not", false, "Reverses the meaning of the ^ prefix (or lack thereof) for all following revision specifiers")
	revListCmd.Flags().Bool("object-names", false, "Print the names of the object IDs that are found")
	revListCmd.Flags().Bool("objects", false, "Print the object IDs of any object referenced by the listed commits")
	revListCmd.Flags().Bool("objects-edge", false, "Similar to --objects, but also print the IDs of excluded commits prefixed with a “-” character")
	revListCmd.Flags().Bool("objects-edge-aggressive", false, "Similar to --objects-edge, but it tries harder to find excluded commits")
	revListCmd.Flags().Bool("oneline", false, "This is a shorthand for \"--pretty=oneline --abbrev-commit\" used together")
	revListCmd.Flags().Bool("parents", false, "Print also the parents of the commit")
	revListCmd.Flags().BoolP("perl-regexp", "P", false, "Consider the limiting patterns to be Perl-compatible regular expressions")
	revListCmd.Flags().String("pretty", "", "Pretty-print the contents of the commit logs in a given format")
	revListCmd.Flags().String("progress", "", "Show progress reports on stderr as objects are considered")
	revListCmd.Flags().Bool("quiet", false, "Don’t print anything to standard output")
	revListCmd.Flags().Bool("reflog", false, "Pretend as if all objects mentioned by reflogs are listed on the command line as <commit>")
	revListCmd.Flags().BoolP("regexp-ignore-case", "i", false, "Match the regular expression limiting patterns without regard to letter case")
	revListCmd.Flags().Bool("relative-date", false, "Synonym for --date=relative")
	revListCmd.Flags().String("remotes", "", "Pretend as if all the refs in refs/remotes are listed on the command line as <commit>")
	revListCmd.Flags().Bool("remove-empty", false, "Stop when a given path disappears from the tree")
	revListCmd.Flags().Bool("reverse", false, "Output the commits chosen to be shown in reverse order")
	revListCmd.Flags().Bool("right-only", false, "List only commits on the left side of a symmetric difference")
	revListCmd.Flags().String("show-linear-break", "", "Put a barrier in between linear branches")
	revListCmd.Flags().Bool("show-pulls", false, "Include all commits from the default mode, but also any merge commits")
	revListCmd.Flags().Bool("show-signature", false, "Check the validity of a signed commit object by passing the signature to gpg --verify")
	revListCmd.Flags().Bool("simplify-by-decoration", false, "Commits that are referred by some branch or tag are selected")
	revListCmd.Flags().Bool("simplify-merges", false, "Simplify each commit C to its replacement C' in the final history")
	revListCmd.Flags().String("since", "", "Show commits more recent than a specific date")
	revListCmd.Flags().String("since-as-filter", "", "Show all commits more recent than a specific date")
	revListCmd.Flags().Bool("single-worktree", false, "Examine the current working tree only")
	revListCmd.Flags().String("skip", "", "Skip number commits before starting to show the commit output")
	revListCmd.Flags().Bool("sparse", false, "All commits in the simplified history are shown")
	revListCmd.Flags().Bool("stdin", false, "Read <commit> additionally from the standard input")
	revListCmd.Flags().String("tags", "", "Pretend as if all the refs in refs/tags are listed on the command line as <commit>")
	revListCmd.Flags().Bool("timestamp", false, "Print the raw commit timestamp")
	revListCmd.Flags().Bool("topo-order", false, "Show no parents before all of its children are shown")
	revListCmd.Flags().Bool("unpacked", false, "Print the object IDs that are not in packs")
	revListCmd.Flags().String("until", "", "Show commits older than a specific date")
	revListCmd.Flags().Bool("use-bitmap-index", false, "Try to speed up the traversal using the pack bitmap index")
	revListCmd.Flags().BoolP("walk-reflogs", "g", false, "Walk reflog entries from the most recent one to older ones")
	rootCmd.AddCommand(revListCmd)

	revListCmd.Flag("ancestry-path").NoOptDefVal = " "
	revListCmd.Flag("ancestry-path").NoOptDefVal = " "
	revListCmd.Flag("branches").NoOptDefVal = " "
	revListCmd.Flag("disk-usage").NoOptDefVal = " "
	revListCmd.Flag("exclude-hidden").NoOptDefVal = " "
	revListCmd.Flag("no-walk").NoOptDefVal = " "
	revListCmd.Flag("pretty").NoOptDefVal = " "
	revListCmd.Flag("remotes").NoOptDefVal = " "
	revListCmd.Flag("show-linear-break").NoOptDefVal = " "
	revListCmd.Flag("tags").NoOptDefVal = " "

	// TODO still a lot of completions missing / incorrect
	carapace.Gen(revListCmd).FlagCompletion(carapace.ActionMap{
		"after":     time.ActionDate(),
		"author":    git.ActionAuthors(),
		"before":    time.ActionDate(),
		"branches":  git.ActionLocalBranches(),
		"committer": git.ActionCommitters(),
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
		"disk-usage": carapace.ActionValues("human"),
		"exclude": carapace.Batch(
			git.ActionLocalBranches(),
			git.ActionTags(),
			git.ActionRemotes(),
		).ToA(),
		"exclude-hidden": carapace.ActionValues("receive", "uploadpack"),
		"format":         carapace.ActionValues("oneline", "short", "medium", "full", "fuller", "reference", "email", "raw", "format:"),
		"missing": carapace.ActionValues(
			"error", "requests that rev-list stop with an error if a missing object is encountered",
			"allow-any", "allow object traversal to continue if a missing object is encountered",
			"allow-promisor", "is like allow-any, but will only allow object traversal to continue for EXPECTED promisor missing objects",
			"print", "is like allow-any, but will also print a list of the missing objects",
		),
		"no-walk":         carapace.ActionValues("sorted", "unsorted"),
		"pretty":          carapace.ActionValues("oneline", "short", "medium", "full", "fuller", "reference", "email", "raw", "format:"),
		"remotes":         git.ActionRemotes(),
		"since":           time.ActionDate(),
		"since-as-filter": time.ActionDate(),
		"tags":            git.ActionTags(),
		"until":           time.ActionDate(),
	})

	carapace.Gen(revListCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			prefix := ""
			if strings.HasPrefix(c.Value, "^") {
				prefix = "^"
			}

			c.Value = strings.TrimPrefix(c.Value, prefix)
			return git.ActionRefs(git.RefOption{}.Default()).Invoke(c).Prefix(prefix).ToA()
		}),
	)

	carapace.Gen(revListCmd).DashAnyCompletion(
		carapace.ActionFiles(),
	)
}
