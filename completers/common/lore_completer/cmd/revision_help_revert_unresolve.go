package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_revert_unresolveCmd = &cobra.Command{
	Use:   "unresolve",
	Short: "Marks the revert unresolved",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_revert_unresolveCmd).Standalone()

	revision_help_revertCmd.AddCommand(revision_help_revert_unresolveCmd)
}
