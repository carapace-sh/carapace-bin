package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_merge_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a merge process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_merge_startCmd).Standalone()

	branch_help_mergeCmd.AddCommand(branch_help_merge_startCmd)
}
