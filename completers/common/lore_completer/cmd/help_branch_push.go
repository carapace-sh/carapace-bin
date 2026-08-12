package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push commits to remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_pushCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_pushCmd)
}
