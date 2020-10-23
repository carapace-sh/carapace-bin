package cmd

import (
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove files from the working tree and from the index",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rmCmd.Flags().Bool("cached", false, "only remove from the index")
	rmCmd.Flags().BoolP("force", "f", false, "override the up-to-date check")
	rmCmd.Flags().Bool("ignore-unmatch", false, "exit with a zero status even if nothing matched")
	rmCmd.Flags().BoolP("dry-run", "n", false, "dry run")
	rmCmd.Flags().Bool("pathspec-file-nul", false, "with --pathspec-from-file, pathspec elements are separated with NUL character")
	rmCmd.Flags().String("pathspec-from-file", "", "read pathspec from file")
	rmCmd.Flags().BoolP("quiet", "q", false, "do not list removed files")
	rmCmd.Flags().BoolS("r", "r", false, "allow recursive removal")
	rootCmd.AddCommand(rmCmd)
}
