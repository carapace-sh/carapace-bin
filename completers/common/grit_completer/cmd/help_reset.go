package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Move the branch tip and optionally reset the index or worktree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_resetCmd).Standalone()

	helpCmd.AddCommand(help_resetCmd)
}
