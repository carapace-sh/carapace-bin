package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_revert_help_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the revert, resetting the current revert state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_revert_help_restartCmd).Standalone()

	revision_revert_helpCmd.AddCommand(revision_revert_help_restartCmd)
}
