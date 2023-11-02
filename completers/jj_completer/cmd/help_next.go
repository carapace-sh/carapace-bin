package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Move the current working copy commit to the next child revision in the repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_nextCmd).Standalone()

	helpCmd.AddCommand(help_nextCmd)
}
