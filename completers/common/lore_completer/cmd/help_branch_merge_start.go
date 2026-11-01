package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_merge_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a merge process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_merge_startCmd).Standalone()

	help_branch_mergeCmd.AddCommand(help_branch_merge_startCmd)
}
