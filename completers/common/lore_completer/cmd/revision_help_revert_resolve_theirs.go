package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_revert_resolve_theirsCmd = &cobra.Command{
	Use:   "theirs",
	Short: "Resolve using the incoming changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_revert_resolve_theirsCmd).Standalone()

	revision_help_revert_resolveCmd.AddCommand(revision_help_revert_resolve_theirsCmd)
}
