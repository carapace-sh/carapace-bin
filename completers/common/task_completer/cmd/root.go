package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/task"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "A task runner / simpler Make alternative written in Go",
	Long:  "https://taskfile.dev/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("color", "c", false, "Colored output. Enabled by default. Set flag to false or use NO_COLOR=1 to disable.")
	rootCmd.Flags().String("completion", "", "Generates shell completion script.")
	rootCmd.Flags().StringP("concurrency", "C", "", "Limit number of tasks to run concurrently.")
	rootCmd.Flags().StringP("dir", "d", "", "Sets the directory in which Task will execute and look for a Taskfile.")
	rootCmd.Flags().BoolP("dry", "n", false, "Compiles and prints tasks in the order that they would be run, without executing them.")
	rootCmd.Flags().BoolP("exit-code", "x", false, "Pass-through the exit code of the task command.")
	rootCmd.Flags().Bool("experiments", false, "Lists all the available experiments and whether or not they are enabled.")
	rootCmd.Flags().BoolP("force", "f", false, "Forces execution even when the task is up-to-date.")
	rootCmd.Flags().BoolP("global", "g", false, "Runs global Taskfile, from $HOME/{T,t}askfile.{yml,yaml}.")
	rootCmd.Flags().BoolP("help", "h", false, "Shows Task usage.")
	rootCmd.Flags().BoolP("init", "i", false, "Creates a new Taskfile.yml in the current folder.")
	rootCmd.Flags().Bool("insecure", false, "Forces Task to download Taskfiles over insecure connections.")
	rootCmd.Flags().StringP("interval", "I", "", "Interval to watch for changes.")
	rootCmd.Flags().BoolP("json", "j", false, "Formats task list as JSON.")
	rootCmd.Flags().BoolP("list", "l", false, "Lists tasks with description of current Taskfile.")
	rootCmd.Flags().BoolP("list-all", "a", false, "Lists tasks with or without a description.")
	rootCmd.Flags().Bool("nested", false, "Nest namespaces when listing tasks as JSON")
	rootCmd.Flags().Bool("no-status", false, "Ignore status when listing tasks as JSON")
	rootCmd.Flags().StringP("output", "o", "", "Sets output style: [interleaved|group|prefixed].")
	rootCmd.Flags().String("output-group-begin", "", "Message template to print before a task's grouped output.")
	rootCmd.Flags().String("output-group-end", "", "Message template to print after a task's grouped output.")
	rootCmd.Flags().Bool("output-group-error-only", false, "Swallow output from successful tasks.")
	rootCmd.Flags().BoolP("parallel", "p", false, "Executes tasks provided on command line in parallel.")
	rootCmd.Flags().BoolP("silent", "s", false, "Disables echoing.")
	rootCmd.Flags().String("sort", "", "Changes the order of the tasks when listed. [default|alphanumeric|none].")
	rootCmd.Flags().Bool("status", false, "Exits with non-zero exit code if any of the given tasks is not up-to-date.")
	rootCmd.Flags().Bool("summary", false, "Show summary about a task.")
	rootCmd.Flags().StringP("taskfile", "t", "", "Choose which Taskfile to run. Defaults to \"Taskfile.yml\".")
	rootCmd.Flags().BoolP("verbose", "v", false, "Enables verbose mode.")
	rootCmd.Flags().Bool("version", false, "Show Task version.")
	rootCmd.Flags().BoolP("watch", "w", false, "Enables watch of the given task.")
	rootCmd.Flags().BoolP("yes", "y", false, "Assume \"yes\" as answer to all prompts.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"completion": carapace.ActionValues("bash", "zsh", "fish", "powershell"),
		"dir":        carapace.ActionDirectories(),
		"output":     carapace.ActionValues("interleaved", "group", "prefixed"),
		"sort":       carapace.ActionValues("default", "alphanumeric", "none"),
		"taskfile":   carapace.ActionFiles(".yml", ".yaml"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return task.ActionTasks(rootCmd.Flag("taskfile").Value.String()).
				Chdir(rootCmd.Flag("dir").Value.String()).
				FilterArgs()
		}),
	)
}
