package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_merge_abortCmd = &cobra.Command{
	Use:   "abort",
	Short: "Abort a merge process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_merge_abortCmd).Standalone()

	branch_help_mergeCmd.AddCommand(branch_help_merge_abortCmd)
}
