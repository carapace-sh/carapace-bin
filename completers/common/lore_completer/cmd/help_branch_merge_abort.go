package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_merge_abortCmd = &cobra.Command{
	Use:   "abort",
	Short: "Abort a merge process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_merge_abortCmd).Standalone()

	help_branch_mergeCmd.AddCommand(help_branch_merge_abortCmd)
}
