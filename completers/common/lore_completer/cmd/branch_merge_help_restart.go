package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_merge_help_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the merge, resetting the current merge state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_merge_help_restartCmd).Standalone()

	branch_merge_helpCmd.AddCommand(branch_merge_help_restartCmd)
}
