package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/devenv_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tasks_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run tasks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_runCmd).Standalone()

	tasks_runCmd.Flags().StringArray("input", nil, "Set a task input value (repeatable, value parsed as JSON if valid, otherwise string)")
	tasks_runCmd.Flags().String("input-json", "", "Set task inputs from a JSON object string")
	tasks_runCmd.Flags().StringP("mode", "m", "", "The execution mode for tasks (affects dependency resolution)")
	tasks_runCmd.Flags().Bool("show-output", false, "Show task output for all tasks (equivalent to --verbose for tasks)")

	tasksCmd.AddCommand(tasks_runCmd)

	carapace.Gen(tasks_runCmd).FlagCompletion(carapace.ActionMap{
		"mode": actionExecutionModes(),
	})

	carapace.Gen(tasks_runCmd).PositionalAnyCompletion(
		action.ActionTasks(tasks_runCmd).FilterArgs(),
	)
}
