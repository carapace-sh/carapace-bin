package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_revert_unresolveCmd = &cobra.Command{
	Use:   "unresolve",
	Short: "Marks the revert unresolved",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_revert_unresolveCmd).Standalone()

	help_revision_revertCmd.AddCommand(help_revision_revert_unresolveCmd)
}
