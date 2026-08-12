package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_merge_resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolves the merge",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_merge_resolveCmd).Standalone()

	branch_help_mergeCmd.AddCommand(branch_help_merge_resolveCmd)
}
