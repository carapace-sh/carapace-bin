package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_revert_unresolveCmd = &cobra.Command{
	Use:   "unresolve",
	Short: "Marks the revert unresolved",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_revert_unresolveCmd).Standalone()

	revision_revert_unresolveCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_revert_unresolveCmd.Flags().String("targets", "", "Path to a targets file")
	revision_revertCmd.AddCommand(revision_revert_unresolveCmd)

	carapace.Gen(revision_revert_unresolveCmd).FlagCompletion(carapace.ActionMap{
		"targets": carapace.ActionFiles(),
	})

	carapace.Gen(revision_revert_unresolveCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
