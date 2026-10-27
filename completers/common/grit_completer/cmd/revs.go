package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var revsCmd = &cobra.Command{
	Use:   "revs",
	Short: "List or count commits matching revision ranges",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revsCmd).Standalone()

	revsCmd.Flags().StringS("G", "G", "", "Keep commits whose added/removed lines match this pattern (`git log -G`)")
	revsCmd.Flags().StringS("S", "S", "", "Keep commits where occurrence count of this string changed (`git log -S`)")
	revsCmd.Flags().String("after", "", "Alias for `--since`")
	revsCmd.Flags().Bool("all", false, "Include every ref tip")
	revsCmd.Flags().Bool("all-match", false, "Require every `--grep` pattern to match")
	revsCmd.Flags().StringSlice("author", nil, "Limit to commits whose author matches this pattern (repeatable; OR together)")
	revsCmd.Flags().String("before", "", "Alias for `--until`, unless the value is an `lw1.` log cursor (deprecated)")
	revsCmd.Flags().String("branches", "", "Include branch tips, or those matching an optional glob")
	revsCmd.Flags().StringSlice("committer", nil, "Limit to commits whose committer matches this pattern (repeatable; OR together)")
	revsCmd.Flags().Bool("count", false, "Print only the number of reachable commits")
	revsCmd.Flags().Bool("first-parent", false, "Follow only the first parent at merges")
	revsCmd.Flags().BoolP("fixed-strings", "F", false, "Treat `--author`, `--committer`, and `--grep` patterns as fixed strings")
	revsCmd.Flags().StringSlice("grep", nil, "Limit to commits whose message matches this pattern (repeatable; OR unless `--all-match`)")
	revsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	revsCmd.Flags().Bool("invert-grep", false, "Invert `--grep` matching (commits that do not match the grep group)")
	revsCmd.Flags().Bool("left-right", false, "With `--count`, print left/right counts for a symmetric range (`A...B`)")
	revsCmd.Flags().StringP("limit", "n", "", "Maximum number of rows to list")
	revsCmd.Flags().String("max-count", "", "Maximum number of rows to list")
	revsCmd.Flags().Bool("merges", false, "Include only merge commits")
	revsCmd.Flags().Bool("no-mailmap", false, "Do not apply `.mailmap` when matching `--author` or `--committer`")
	revsCmd.Flags().Bool("no-merges", false, "Exclude merge commits")
	revsCmd.Flags().Bool("pickaxe-regex", false, "Treat `-S` as a BRE instead of a literal string (`--pickaxe-regex`)")
	revsCmd.Flags().BoolP("regexp-ignore-case", "i", false, "Case-insensitive pattern matching for `--author`, `--committer`, and `--grep`")
	revsCmd.Flags().String("remotes", "", "Include remote-tracking tips, or those matching an optional glob")
	revsCmd.Flags().Bool("reverse", false, "Reverse commit order (oldest first)")
	revsCmd.Flags().String("since", "", "Exclude commits older than this date (Git `--since`)")
	revsCmd.Flags().String("skip", "0", "Skip this many commits after ordering")
	revsCmd.Flags().String("tags", "", "Include tag tips, or those matching an optional glob")
	revsCmd.Flags().String("until", "", "Exclude commits newer than this date (Git `--until`)")
	rootCmd.AddCommand(revsCmd)

	carapace.Gen(revsCmd).PositionalAnyCompletion(
		git.ActionRefRanges(git.RefOption{}.Default()),
	)
}
