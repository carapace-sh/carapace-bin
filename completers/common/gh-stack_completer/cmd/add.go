package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [branch]",
	Short: "Add a new branch on top of the current stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().BoolP("all", "A", false, "Stage all changes including untracked files")
	addCmd.Flags().StringP("message", "m", "", "Create a commit with this message")
	addCmd.Flags().BoolP("update", "u", false, "Stage changes to tracked files only")
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).PositionalCompletion(
		git.ActionLocalBranches(),
	)
}
