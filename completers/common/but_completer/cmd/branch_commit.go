package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a new commit to named virtual branch with all changes currently in the worktree or staging area assigned to it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_commitCmd).Standalone()

	branch_commitCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_commitCmd.Flags().StringP("message", "m", "", "The commit message")
	branch_commitCmd.MarkFlagRequired("message")
	branchCmd.AddCommand(branch_commitCmd)
}
