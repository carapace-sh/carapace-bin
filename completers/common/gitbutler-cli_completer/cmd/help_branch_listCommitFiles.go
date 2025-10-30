package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_listCommitFilesCmd = &cobra.Command{
	Use:   "list-commit-files",
	Short: "List all changes of a commit, along with their patches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_listCommitFilesCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_listCommitFilesCmd)
}
