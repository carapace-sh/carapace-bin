package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_revert_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the revert, resetting the current revert state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_revert_restartCmd).Standalone()

	revision_revert_restartCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_revert_restartCmd.Flags().String("targets", "", "Path to a targets file")
	revision_revertCmd.AddCommand(revision_revert_restartCmd)

	carapace.Gen(revision_revert_restartCmd).FlagCompletion(carapace.ActionMap{
		"targets": carapace.ActionFiles(),
	})

	carapace.Gen(revision_revert_restartCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
