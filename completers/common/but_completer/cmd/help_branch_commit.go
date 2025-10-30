package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a new commit to named virtual branch with all changes currently in the worktree or staging area assigned to it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_commitCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_commitCmd)
}
