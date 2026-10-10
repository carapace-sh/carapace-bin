package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tasks_runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Run tasks",
	Aliases: []string{"r"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_runCmd).Standalone()

	tasksCmd.AddCommand(tasks_runCmd)

	carapace.Gen(tasks_runCmd).PositionalCompletion(
		action.ActionTasks(),
	)
}
