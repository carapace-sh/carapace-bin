package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_listCommitFilesCmd = &cobra.Command{
	Use:   "list-commit-files",
	Short: "List all changes of a commit, along with their patches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_listCommitFilesCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_listCommitFilesCmd)
}
