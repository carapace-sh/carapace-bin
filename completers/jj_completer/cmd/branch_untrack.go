package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var branch_untrackCmd = &cobra.Command{
	Use:   "untrack",
	Short: "Stop tracking given remote branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_untrackCmd).Standalone()

	branch_untrackCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branchCmd.AddCommand(branch_untrackCmd)
}
