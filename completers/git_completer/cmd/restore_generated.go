package cmd

import (
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore working tree files",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	restoreCmd.Flags().BoolP("ours", "2", false, "checkout our version for unmerged files")
	restoreCmd.Flags().BoolP("theirs", "3", false, "checkout their version for unmerged files")
	restoreCmd.Flags().String("conflict", "", "conflict style (merge or diff3)")
	restoreCmd.Flags().Bool("ignore-skip-worktree-bits", false, "do not limit pathspecs to sparse entries only")
	restoreCmd.Flags().Bool("ignore-unmerged", false, "ignore unmerged entries")
	restoreCmd.Flags().BoolP("merge", "m", false, "perform a 3-way merge with the new branch")
	restoreCmd.Flags().Bool("overlay", false, "use overlay mode")
	restoreCmd.Flags().Bool("pathspec-file-nul", false, "with --pathspec-from-file, pathspec elements are separated with NUL character")
	restoreCmd.Flags().String("pathspec-from-file", "", "read pathspec from file")
	restoreCmd.Flags().BoolP("patch", "p", false, "select hunks interactively")
	restoreCmd.Flags().Bool("progress", false, "force progress reporting")
	restoreCmd.Flags().BoolP("quiet", "q", false, "suppress progress reporting")
	restoreCmd.Flags().String("recurse-submodules", "", "control recursive updating of submodules")
	restoreCmd.Flags().BoolP("source", "s", false, "<tree-ish>    which tree-ish to checkout from")
	restoreCmd.Flags().BoolP("staged", "S", false, "restore the index")
	restoreCmd.Flags().BoolP("worktree", "W", false, "restore the working tree (default)")
	rootCmd.AddCommand(restoreCmd)
}
