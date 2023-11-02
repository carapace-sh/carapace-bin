package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_branch_trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Start tracking given remote branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_trackCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_trackCmd)
}
