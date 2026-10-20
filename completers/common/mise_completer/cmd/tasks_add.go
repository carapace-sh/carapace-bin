package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var tasks_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_addCmd).Standalone()

	tasks_addCmd.Flags().StringSliceP("alias", "a", nil, "Other names for the task")
	tasks_addCmd.Flags().StringSliceP("depends", "d", nil, "Tasks to run before this one")
	tasks_addCmd.Flags().StringSlice("depends-post", nil, "Tasks to run after this one")
	tasks_addCmd.Flags().String("description", "", "Description of the task")
	tasks_addCmd.Flags().StringP("dir", "D", "", "Run the task in a specific directory")
	tasks_addCmd.Flags().BoolP("file", "f", false, "Create a file task instead of a TOML task")
	tasks_addCmd.Flags().BoolP("help", "h", false, "Print help")
	tasks_addCmd.Flags().BoolP("hide", "H", false, "Hide the task from `mise tasks` and completions")
	tasks_addCmd.Flags().StringSlice("outputs", nil, "Glob patterns of files this task creates, used to skip it when they are up to date")
	tasks_addCmd.Flags().BoolP("quiet", "q", false, "Do not print the command before running")
	tasks_addCmd.Flags().BoolP("raw", "r", false, "Connect the task directly to the terminal's stdin, stdout, and stderr")
	tasks_addCmd.Flags().String("run-windows", "", "Command to run on Windows (TOML tasks only; ignored with `--file`)")
	tasks_addCmd.Flags().String("shell", "", "Run the task in a specific shell")
	tasks_addCmd.Flags().Bool("silent", false, "Do not print the command or its output")
	tasks_addCmd.Flags().StringSliceP("sources", "s", nil, "Glob patterns of files this task uses as input")
	tasks_addCmd.Flags().StringSliceP("wait-for", "w", nil, "Wait for these tasks to finish if they are also being run")
	tasksCmd.AddCommand(tasks_addCmd)

	carapace.Gen(tasks_addCmd).FlagCompletion(carapace.ActionMap{
		"depends":      mise.ActionTasks(),
		"depends-post": mise.ActionTasks(),
		"dir":          carapace.ActionDirectories(),
		"wait-for":     mise.ActionTasks(),
	})

	carapace.Gen(tasks_addCmd).PositionalCompletion(
		mise.ActionTasks(),
	)
}
