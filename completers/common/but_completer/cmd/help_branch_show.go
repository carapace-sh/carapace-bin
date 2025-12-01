package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show commits ahead of base for a specific branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_showCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_showCmd)
}
