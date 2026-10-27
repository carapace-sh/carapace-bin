package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/grit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Stage changes. With no paths, stages everything",
	GroupID: "changes",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().Bool("dry-run", false, "Report paths that would be staged without writing the index or object database")
	addCmd.Flags().Bool("executable", false, "Mark staged regular files as executable (`100755`) in the index")
	addCmd.Flags().BoolP("force", "f", false, "Stage paths that would otherwise be rejected as ignored")
	addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	addCmd.Flags().StringSlice("hunks", nil, "Stage only the listed hunk ids from `grit diff --json` (comma-separated, repeatable)")
	addCmd.Flags().BoolP("intent-to-add", "N", false, "Record intent to add (empty index entries) instead of staging file content")
	addCmd.Flags().StringSlice("lines", nil, "Stage only new-side line ranges from `grit diff --json` (`path:start-end`, repeatable)")
	addCmd.Flags().Bool("non-executable", false, "Mark staged regular files as non-executable (`100644`) in the index")
	addCmd.Flags().String("patch-file", "", "Apply a unified diff to the index (`git apply --cached`). Use `-` for stdin")
	addCmd.Flags().Bool("pathspec-file-nul", false, "`--pathspec-from-file` entries are NUL-separated")
	addCmd.Flags().String("pathspec-from-file", "", "Read pathspecs from a file (`-` for stdin) instead of or in addition to argv paths")
	addCmd.Flags().BoolP("update", "u", false, "Stage tracked changes only (modified and deleted paths; no untracked files)")
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"hunks": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionHunkIds("unstaged")
		}),
		"lines": action.ActionLineSelections("unstaged"),
	})
}
