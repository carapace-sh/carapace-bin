package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tasks_depsCmd = &cobra.Command{
	Use:   "deps",
	Short: "Display task dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_depsCmd).Standalone()

	tasksCmd.AddCommand(tasks_depsCmd)

	carapace.Gen(tasks_depsCmd).PositionalCompletion(
		action.ActionTasks(),
	)
}
