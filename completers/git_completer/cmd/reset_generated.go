package cmd

import (
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset current HEAD to the specified state",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	resetCmd.Flags().Bool("hard", false, "reset HEAD, index and working tree")
	resetCmd.Flags().Bool("keep", false, "reset HEAD but keep local changes")
	resetCmd.Flags().Bool("merge", false, "reset HEAD, index and working tree")
	resetCmd.Flags().Bool("mixed", false, "reset HEAD and index")
	resetCmd.Flags().BoolP("intent-to-add", "N", false, "record only the fact that removed paths will be added later")
	resetCmd.Flags().Bool("pathspec-file-nul", false, "with --pathspec-from-file, pathspec elements are separated with NUL character")
	resetCmd.Flags().String("pathspec-from-file", "", "read pathspec from file")
	resetCmd.Flags().BoolP("patch", "p", false, "select hunks interactively")
	resetCmd.Flags().BoolP("quiet", "q", false, "be quiet, only report errors")
	resetCmd.Flags().String("recurse-submodules", "", "control recursive updating of submodules")
	resetCmd.Flags().Bool("soft", false, "reset only HEAD")
	rootCmd.AddCommand(resetCmd)
}
