package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_merge_help_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a merge process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_merge_help_startCmd).Standalone()

	branch_merge_helpCmd.AddCommand(branch_merge_help_startCmd)
}
