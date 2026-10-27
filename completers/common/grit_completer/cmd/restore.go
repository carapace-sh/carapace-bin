package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore working tree and/or index paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restoreCmd).Standalone()

	restoreCmd.Flags().String("conflict", "", "Conflict marker style (implies `--merge`)")
	restoreCmd.Flags().Bool("dry-run", false, "Report paths that would change without writing the worktree or index")
	restoreCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	restoreCmd.Flags().StringSlice("hunks", nil, "Discard only the listed unstaged hunk ids (`grit diff --unstaged --json`) or unstage staged ids (`--staged`)")
	restoreCmd.Flags().StringSlice("lines", nil, "Discard or unstage only new-side line ranges (`path:start-end` from `grit diff --json`)")
	restoreCmd.Flags().Bool("merge", false, "Rebuild conflict markers in the worktree from index stages")
	restoreCmd.Flags().Bool("ours", false, "For unmerged paths, write our side (index stage 2) into the worktree")
	restoreCmd.Flags().Bool("pathspec-file-nul", false, "`--pathspec-from-file` entries are NUL-separated")
	restoreCmd.Flags().String("pathspec-from-file", "", "Read pathspecs from a file (`-` for stdin) instead of or in addition to argv paths")
	restoreCmd.Flags().String("source", "", "Tree or commit to restore from")
	restoreCmd.Flags().BoolP("staged", "S", false, "Restore the index from HEAD (or `--source`)")
	restoreCmd.Flags().Bool("theirs", false, "For unmerged paths, write their side (index stage 3) into the worktree")
	restoreCmd.Flags().BoolP("worktree", "W", false, "Restore the working tree (default when `--staged` is absent)")
	rootCmd.AddCommand(restoreCmd)

	carapace.Gen(restoreCmd).FlagCompletion(carapace.ActionMap{
		"conflict": carapace.ActionValuesDescribed(
			"diff3", "Three-section markers including the base section",
			"merge", "Standard two-section conflict markers",
			"zdiff3", "Zealous diff3 (`zdiff3`)",
		),
		"source": git.ActionRefs(git.RefOption{}.Default()),
	})

	carapace.Gen(restoreCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
