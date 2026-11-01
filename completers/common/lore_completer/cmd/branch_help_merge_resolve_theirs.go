package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_merge_resolve_theirsCmd = &cobra.Command{
	Use:   "theirs",
	Short: "Resolve using their changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_merge_resolve_theirsCmd).Standalone()

	branch_help_merge_resolveCmd.AddCommand(branch_help_merge_resolve_theirsCmd)
}
