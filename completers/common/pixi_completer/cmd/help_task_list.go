package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_task_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_task_listCmd).Standalone()

	help_taskCmd.AddCommand(help_task_listCmd)
}
