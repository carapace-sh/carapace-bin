package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_revert_help_unresolveCmd = &cobra.Command{
	Use:   "unresolve",
	Short: "Marks the revert unresolved",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_revert_help_unresolveCmd).Standalone()

	revision_revert_helpCmd.AddCommand(revision_revert_help_unresolveCmd)
}
