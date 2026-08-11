package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_revert_resolve_help_mineCmd = &cobra.Command{
	Use:   "mine",
	Short: "Resolve using my changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_revert_resolve_help_mineCmd).Standalone()

	revision_revert_resolve_helpCmd.AddCommand(revision_revert_resolve_help_mineCmd)
}
