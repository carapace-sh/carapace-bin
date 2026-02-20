package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var task_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(task_help_listCmd).Standalone()

	task_helpCmd.AddCommand(task_help_listCmd)
}
