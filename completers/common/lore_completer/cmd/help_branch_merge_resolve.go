package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_merge_resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolves the merge",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_merge_resolveCmd).Standalone()

	help_branch_mergeCmd.AddCommand(help_branch_merge_resolveCmd)
}
