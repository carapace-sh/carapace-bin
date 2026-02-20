package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var task_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a command from the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(task_help_removeCmd).Standalone()

	task_helpCmd.AddCommand(task_help_removeCmd)
}
