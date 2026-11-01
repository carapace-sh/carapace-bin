package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_merge_intoCmd = &cobra.Command{
	Use:   "into",
	Short: "Merge into branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_merge_intoCmd).Standalone()

	help_branch_mergeCmd.AddCommand(help_branch_merge_intoCmd)
}
