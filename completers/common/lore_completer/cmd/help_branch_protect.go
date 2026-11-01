package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_protectCmd = &cobra.Command{
	Use:   "protect",
	Short: "Protect a branch from direct pushes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_protectCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_protectCmd)
}
