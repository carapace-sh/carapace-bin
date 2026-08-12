package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_merge_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the merge, resetting the current merge state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_merge_restartCmd).Standalone()

	help_branch_mergeCmd.AddCommand(help_branch_merge_restartCmd)
}
