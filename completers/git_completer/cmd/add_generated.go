package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add file contents to the index",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	addCmd.Flags().BoolP("all", "A", false, "add changes from all tracked and untracked files")
	addCmd.Flags().String("chmod", "", "override the executable bit of the listed files")
	addCmd.Flags().BoolP("edit", "e", false, "edit current diff and apply")
	addCmd.Flags().BoolP("force", "f", false, "allow adding otherwise ignored files")
	addCmd.Flags().Bool("ignore-errors", false, "just skip files which cannot be added because of errors")
	addCmd.Flags().Bool("ignore-missing", false, "check if - even missing - files are ignored in dry run")
	addCmd.Flags().Bool("ignore-removal", false, "ignore paths removed in the working tree (same as --no-all)")
	addCmd.Flags().BoolP("interactive", "i", false, "interactive picking")
	addCmd.Flags().BoolP("dry-run", "n", false, "dry run")
	addCmd.Flags().BoolP("intent-to-add", "N", false, "record only the fact that the path will be added later")
	addCmd.Flags().Bool("pathspec-file-nul", false, "with --pathspec-from-file, pathspec elements are separated with NUL character")
	addCmd.Flags().String("pathspec-from-file", "", "read pathspec from file")
	addCmd.Flags().BoolP("patch", "p", false, "select hunks interactively")
	addCmd.Flags().Bool("refresh", false, "don't add, only refresh the index")
	addCmd.Flags().Bool("renormalize", false, "renormalize EOL of tracked files (implies -u)")
	addCmd.Flags().BoolP("update", "u", false, "update tracked files")
	addCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	rootCmd.AddCommand(addCmd)
}
