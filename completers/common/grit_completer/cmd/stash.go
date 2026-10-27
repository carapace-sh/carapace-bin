package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/grit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stashCmd = &cobra.Command{
	Use:   "stash",
	Short: "Save or restore shelved work",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stashCmd).Standalone()

	stashCmd.PersistentFlags().Bool("dry-run", false, "Preview stash without updating refs, index, or worktree")
	stashCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stashCmd.PersistentFlags().StringSlice("hunks", nil, "Stash only unstaged hunks from `grit diff --json` (comma-separated, repeatable)")
	stashCmd.PersistentFlags().BoolP("include-untracked", "u", false, "Include untracked files (bare `grit stash` / `grit stash push`)")
	stashCmd.PersistentFlags().BoolP("keep-index", "k", false, "After stashing, leave the index and worktree matching staged content")
	stashCmd.PersistentFlags().StringP("limit", "n", "", "Cap path lists in `--dry-run` preview output")
	stashCmd.PersistentFlags().StringSlice("lines", nil, "Stash only unstaged line ranges (`path:start-end`, same as `grit add --lines`)")
	stashCmd.PersistentFlags().StringP("message", "m", "", "Stash message (bare `grit stash` / `grit stash push`)")
	stashCmd.Flags().Bool("pathspec-file-nul", false, "`--pathspec-from-file` entries are NUL-separated")
	stashCmd.Flags().String("pathspec-from-file", "", "Read pathspecs from a file (`-` for stdin) instead of or in addition to argv paths")
	stashCmd.PersistentFlags().BoolP("staged", "S", false, "Stash only index changes vs `HEAD` (unstaged edits on other paths remain)")
	rootCmd.AddCommand(stashCmd)

	carapace.Gen(stashCmd).FlagCompletion(carapace.ActionMap{
		"hunks": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionHunkIds("unstaged")
		}),
		"lines": action.ActionLineSelections("unstaged"),
	})
}
