package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var tasks_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open a task in your editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_editCmd).Standalone()

	tasks_editCmd.Flags().BoolP("help", "h", false, "Print help")
	tasks_editCmd.Flags().BoolP("path", "p", false, "Print the task's file path instead of opening it")
	tasksCmd.AddCommand(tasks_editCmd)

	carapace.Gen(tasks_editCmd).PositionalCompletion(
		mise.ActionTasks(),
	)
}
