package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_revert_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the revert, resetting the current revert state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_revert_restartCmd).Standalone()

	revision_help_revertCmd.AddCommand(revision_help_revert_restartCmd)
}
