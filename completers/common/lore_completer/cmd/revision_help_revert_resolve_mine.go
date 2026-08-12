package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_revert_resolve_mineCmd = &cobra.Command{
	Use:   "mine",
	Short: "Resolve using my changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_revert_resolve_mineCmd).Standalone()

	revision_help_revert_resolveCmd.AddCommand(revision_help_revert_resolve_mineCmd)
}
