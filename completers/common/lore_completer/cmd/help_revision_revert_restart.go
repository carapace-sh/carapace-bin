package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_revert_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the revert, resetting the current revert state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_revert_restartCmd).Standalone()

	help_revision_revertCmd.AddCommand(help_revision_revert_restartCmd)
}
