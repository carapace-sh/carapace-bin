package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_merge_resolve_theirsCmd = &cobra.Command{
	Use:   "theirs",
	Short: "Resolve using their changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_merge_resolve_theirsCmd).Standalone()

	help_branch_merge_resolveCmd.AddCommand(help_branch_merge_resolve_theirsCmd)
}
