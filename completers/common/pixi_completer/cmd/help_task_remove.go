package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_task_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a command from the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_task_removeCmd).Standalone()

	help_taskCmd.AddCommand(help_task_removeCmd)
}
