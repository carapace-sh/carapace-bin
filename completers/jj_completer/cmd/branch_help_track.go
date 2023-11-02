package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var branch_help_trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Start tracking given remote branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_trackCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_trackCmd)
}
