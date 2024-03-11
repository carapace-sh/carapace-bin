package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stageCmd = &cobra.Command{
	Use:     "stage",
	Short:   "Add all yarn files to your vcs",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stageCmd).Standalone()

	stageCmd.Flags().BoolP("commit", "c", false, "Commit the staged files")
	stageCmd.Flags().BoolP("dry-run", "n", false, "Print the commit message and the list of modified files without staging / committing")
	stageCmd.Flags().BoolP("reset", "r", false, "Remove all files from the staging area")
	rootCmd.AddCommand(stageCmd)
}
