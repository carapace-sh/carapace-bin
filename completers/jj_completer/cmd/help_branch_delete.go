package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_branch_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing branch and propagate the deletion to remotes on the next push",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_deleteCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_deleteCmd)
}
