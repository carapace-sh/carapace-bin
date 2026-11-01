package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_merge_resolve_mineCmd = &cobra.Command{
	Use:   "mine",
	Short: "Resolve using my changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_merge_resolve_mineCmd).Standalone()

	help_branch_merge_resolveCmd.AddCommand(help_branch_merge_resolve_mineCmd)
}
