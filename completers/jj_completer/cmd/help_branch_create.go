package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_branch_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_createCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_createCmd)
}
