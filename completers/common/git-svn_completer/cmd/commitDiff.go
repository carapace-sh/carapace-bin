package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var commitDiffCmd = &cobra.Command{
	Use:   "commit-diff",
	Short: "Commit a diff between two trees",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commitDiffCmd).Standalone()

	commitDiffCmd.Flags().IntP("copy-similarity", "C", 0, "Copy similarity threshold")
	commitDiffCmd.Flags().BoolP("edit", "e", false, "Edit commit message")
	commitDiffCmd.Flags().StringP("file", "F", "", "Take commit message from file")
	commitDiffCmd.Flags().Bool("find-copies-harder", false, "Find copies harder")
	commitDiffCmd.Flags().IntS("l", "l", 0, "Rename limit")
	commitDiffCmd.Flags().StringP("message", "m", "", "Commit message")
	commitDiffCmd.Flags().StringP("revision", "r", "", "SVN revision (required)")
	commitDiffCmd.Flags().Bool("rmdir", false, "Remove empty directories")
	rootCmd.AddCommand(commitDiffCmd)

	carapace.Gen(commitDiffCmd).FlagCompletion(carapace.ActionMap{
		"file":     carapace.ActionFiles(),
		"message":  carapace.ActionValues(),
		"revision": carapace.ActionValues(),
	})

	carapace.Gen(commitDiffCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionValues(),
		carapace.ActionValues(),
	)
}
