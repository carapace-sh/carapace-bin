package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_merge_unresolveCmd = &cobra.Command{
	Use:   "unresolve",
	Short: "Marks the merge unresolved",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_merge_unresolveCmd).Standalone()

	help_branch_mergeCmd.AddCommand(help_branch_merge_unresolveCmd)
}
