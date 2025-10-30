package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all branches that aren't GitButler specific",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_listCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_listCmd)
}
