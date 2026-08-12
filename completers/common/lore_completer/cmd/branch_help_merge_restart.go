package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_merge_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the merge, resetting the current merge state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_merge_restartCmd).Standalone()

	branch_help_mergeCmd.AddCommand(branch_help_merge_restartCmd)
}
