package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/git_completer/cmd/common"
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

	common.AddCommitLimitingOptions(revListCmd)
	common.AddHistorySimplificationOptions(revListCmd)
	common.AddBisectionHelperOptions(revListCmd)
	common.AddCommitOrderingOptions(revListCmd)
	common.AddObjectTraversalOptions(revListCmd)
	common.AddCommitFormattingOptions(revListCmd)
	rootCmd.AddCommand(revListCmd)

	// TODO move
	revListCmd.Flag("pretty").NoOptDefVal = " "

	// TODO still a lot of completions missing / incorrect
	// TODO move
	carapace.Gen(revListCmd).FlagCompletion(carapace.ActionMap{
		"author":    git.ActionAuthors(),
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
