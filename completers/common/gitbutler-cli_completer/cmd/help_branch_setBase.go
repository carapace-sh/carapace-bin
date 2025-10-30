package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_setBaseCmd = &cobra.Command{
	Use:   "set-base",
	Short: "Switch to the GitButler workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_setBaseCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_setBaseCmd)
}
