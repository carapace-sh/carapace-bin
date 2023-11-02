package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var branch_help_untrackCmd = &cobra.Command{
	Use:   "untrack",
	Short: "Stop tracking given remote branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_untrackCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_untrackCmd)
}
