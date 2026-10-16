package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tasks_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display information about a task",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_infoCmd).Standalone()

	tasksCmd.AddCommand(tasks_infoCmd)

	carapace.Gen(tasks_infoCmd).PositionalCompletion(
		action.ActionTasks(),
	)
}
