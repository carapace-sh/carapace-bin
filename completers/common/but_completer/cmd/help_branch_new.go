package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new branch in the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_newCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_newCmd)
}
