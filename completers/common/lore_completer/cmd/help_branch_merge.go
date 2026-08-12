package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge two branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_mergeCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_mergeCmd)
}
