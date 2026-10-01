package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tasks_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_addCmd).Standalone()

	tasksCmd.AddCommand(tasks_addCmd)
}
