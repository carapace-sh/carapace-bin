package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a new commit to named virtual branch with all changes currently in the worktree or staging area assigned to it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_commitCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_commitCmd)
}
