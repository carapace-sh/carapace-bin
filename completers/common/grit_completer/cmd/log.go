package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show recent commits reachable from HEAD",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logCmd).Standalone()

	logCmd.Flags().StringS("G", "G", "", "Keep commits whose added/removed lines match this pattern (`git log -G`)")
	logCmd.Flags().StringS("S", "S", "", "Keep commits where occurrence count of this string changed (`git log -S`)")
	logCmd.Flags().String("after", "", "Alias for `--since`")
	logCmd.Flags().Bool("all", false, "Include every ref tip")
	logCmd.Flags().Bool("all-match", false, "Require every `--grep` pattern to match")
	logCmd.Flags().StringSlice("author", nil, "Limit to commits whose author matches this pattern (repeatable; OR together)")
	logCmd.Flags().String("before", "", "Alias for `--until`, unless the value is an `lw1.` log cursor (deprecated)")
	logCmd.Flags().String("branches", "", "Include branch tips, or those matching an optional glob")
	logCmd.Flags().StringSlice("committer", nil, "Limit to commits whose committer matches this pattern (repeatable; OR together)")
	logCmd.Flags().Bool("count", false, "With `--group-by`, omit per-commit lists in `--json` output")
	logCmd.Flags().String("cursor", "", "Continue a paged walk using the opaque `next` token from a previous page")
	logCmd.Flags().Bool("decorate", false, "Show branch and tag names that point at each commit")
	logCmd.Flags().Bool("first-parent", false, "Follow only the first parent at merges")
	logCmd.Flags().BoolP("fixed-strings", "F", false, "Treat `--author`, `--committer`, and `--grep` patterns as fixed strings")
	logCmd.Flags().BoolP("function-context", "W", false, "Expand hunks to enclosing function boundaries (Git `-W` / `--function-context`)")
	logCmd.Flags().Bool("graph", false, "Draw an ASCII commit graph in human output (implies topological order)")
	logCmd.Flags().Bool("graph-unicode", false, "Use Unicode box-drawing characters in `--graph` output (human only)")
	logCmd.Flags().StringSlice("grep", nil, "Limit to commits whose message matches this pattern (repeatable; OR unless `--all-match`)")
	logCmd.Flags().String("group-by", "", "Aggregate commits (`author` matches `git shortlog -sne` ordering)")
	logCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	logCmd.Flags().Bool("invert-grep", false, "Invert `--grep` matching (commits that do not match the grep group)")
	logCmd.Flags().StringP("limit", "n", "", "Maximum number of rows to list")
	logCmd.Flags().String("max-count", "", "Maximum number of rows to list")
	logCmd.Flags().Bool("merges", false, "Include only merge commits")
	logCmd.Flags().Bool("name-only", false, "Paths only, one per line")
	logCmd.Flags().Bool("name-status", false, "Status letter and path(s) per line")
	logCmd.Flags().Bool("no-mailmap", false, "Do not apply `.mailmap` when matching `--author` or `--committer`")
	logCmd.Flags().Bool("no-merges", false, "Exclude merge commits")
	logCmd.Flags().Bool("numstat", false, "Machine-readable `insertions<TAB>deletions<TAB>path` lines")
	logCmd.Flags().Bool("oneline", false, "One line per commit: short oid and subject (refs when `--decorate`)")
	logCmd.Flags().BoolP("patch", "p", false, "Unified patch for each commit (`git log -p`)")
	logCmd.Flags().Bool("patch-text", false, "Git-apply-compatible unified patch (`git diff --binary` shape); not combinable with `--json`")
	logCmd.Flags().Bool("pickaxe-regex", false, "Treat `-S` as a BRE instead of a literal string (`--pickaxe-regex`)")
	logCmd.Flags().BoolP("regexp-ignore-case", "i", false, "Case-insensitive pattern matching for `--author`, `--committer`, and `--grep`")
	logCmd.Flags().String("remotes", "", "Include remote-tracking tips, or those matching an optional glob")
	logCmd.Flags().Bool("reverse", false, "Reverse commit order (oldest first)")
	logCmd.Flags().Bool("shortstat", false, "Only the summary line (`N files changed, …`)")
	logCmd.Flags().String("since", "", "Exclude commits older than this date (Git `--since`)")
	logCmd.Flags().String("skip", "0", "Skip this many commits after ordering")
	logCmd.Flags().Bool("stat", false, "Summary with change bars per file (like `git diff --stat`)")
	logCmd.Flags().String("tags", "", "Include tag tips, or those matching an optional glob")
	logCmd.Flags().StringP("unified", "U", "", "Lines of context around each change (Git `-U`)")
	logCmd.Flags().String("until", "", "Exclude commits newer than this date (Git `--until`)")
	rootCmd.AddCommand(logCmd)

	carapace.Gen(logCmd).FlagCompletion(carapace.ActionMap{
		"group-by": carapace.ActionValues("author"),
	})

	carapace.Gen(logCmd).PositionalAnyCompletion(
		git.ActionRefRanges(git.RefOption{}.Default()),
	)
}
