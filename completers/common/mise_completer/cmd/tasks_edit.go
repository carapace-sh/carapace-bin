package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tasks_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a task",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_editCmd).Standalone()

	tasksCmd.AddCommand(tasks_editCmd)

	carapace.Gen(tasks_editCmd).PositionalCompletion(
		action.ActionTasks(),
	)
}
