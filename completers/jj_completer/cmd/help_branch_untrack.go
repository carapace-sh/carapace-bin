package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_branch_untrackCmd = &cobra.Command{
	Use:   "untrack",
	Short: "Stop tracking given remote branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_untrackCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_untrackCmd)
}
