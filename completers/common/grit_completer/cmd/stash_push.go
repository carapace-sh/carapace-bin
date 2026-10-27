package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stash_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Save local changes to the stash and reset the working tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_pushCmd).Standalone()

	stash_pushCmd.Flags().Bool("dry-run", false, "Preview stash without updating refs, index, or worktree")
	stash_pushCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stash_pushCmd.Flags().StringSlice("hunks", nil, "Stash only unstaged hunks from `grit diff --json`")
	stash_pushCmd.Flags().BoolP("include-untracked", "u", false, "Include untracked files in the stash")
	stash_pushCmd.Flags().BoolP("keep-index", "k", false, "Leave staged content in the index and worktree after stashing")
	stash_pushCmd.Flags().StringSlice("lines", nil, "Stash only unstaged line ranges (`path:start-end`)")
	stash_pushCmd.Flags().StringP("message", "m", "", "Stash message")
	stash_pushCmd.Flags().Bool("pathspec-file-nul", false, "`--pathspec-from-file` entries are NUL-separated")
	stash_pushCmd.Flags().String("pathspec-from-file", "", "Read pathspecs from a file (`-` for stdin) instead of or in addition to argv paths")
	stash_pushCmd.Flags().BoolP("staged", "S", false, "Stash only index changes vs `HEAD`")
	stashCmd.AddCommand(stash_pushCmd)

	carapace.Gen(stash_pushCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
