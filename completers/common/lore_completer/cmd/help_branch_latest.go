package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_latestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Branch latest related commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_latestCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_latestCmd)
}
