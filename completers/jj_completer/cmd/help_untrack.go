package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_untrackCmd = &cobra.Command{
	Use:   "untrack",
	Short: "Stop tracking specified paths in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_untrackCmd).Standalone()

	helpCmd.AddCommand(help_untrackCmd)
}
