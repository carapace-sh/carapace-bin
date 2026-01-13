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

	rootCmd.PersistentFlags().BoolP("color", "c", false, "Colored output. Enabled by default. Set flag to false or use NO_COLOR=1 to disable.")
	rootCmd.PersistentFlags().String("completion", "", "Generates shell completion script.")
	rootCmd.PersistentFlags().StringP("concurrency", "C", "", "Limit number of tasks to run concurrently.")
	rootCmd.PersistentFlags().StringP("dir", "d", "", "Sets the directory in which Task will execute and look for a Taskfile.")
	rootCmd.PersistentFlags().Bool("disable-fuzzy", false, "Disables fuzzy matching for task names.")
	rootCmd.PersistentFlags().BoolP("dry", "n", false, "Compiles and prints tasks in the order that they would be run, without executing them.")
	rootCmd.PersistentFlags().BoolP("exit-code", "x", false, "Pass-through the exit code of the task command.")
	rootCmd.PersistentFlags().Bool("experiments", false, "Lists all the available experiments and whether or not they are enabled.")
	rootCmd.PersistentFlags().BoolP("failfast", "F", false, "When running tasks in parallel, stop all tasks if one fails.")
	rootCmd.PersistentFlags().BoolP("force", "f", false, "Forces execution even when the task is up-to-date.")
	rootCmd.PersistentFlags().BoolP("global", "g", false, "Runs global Taskfile, from $HOME/{T,t}askfile.{yml,yaml}.")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Shows Task usage.")
	rootCmd.PersistentFlags().BoolP("init", "i", false, "Creates a new Taskfile.yml in the current folder.")
	rootCmd.PersistentFlags().Bool("insecure", false, "Forces Task to download Taskfiles over insecure connections.")
	rootCmd.PersistentFlags().StringP("interval", "I", "", "Interval to watch for changes.")
	rootCmd.PersistentFlags().BoolP("json", "j", false, "Formats task list as JSON.")
	rootCmd.PersistentFlags().BoolP("list", "l", false, "Lists tasks with description of current Taskfile.")
	rootCmd.PersistentFlags().BoolP("list-all", "a", false, "Lists tasks with or without a description.")
	rootCmd.PersistentFlags().Bool("nested", false, "Nest namespaces when listing tasks as JSON")
	rootCmd.PersistentFlags().Bool("no-status", false, "Ignore status when listing tasks as JSON")
	rootCmd.PersistentFlags().StringP("output", "o", "", "Sets output style: [interleaved|group|prefixed].")
	rootCmd.PersistentFlags().String("output-group-begin", "", "Message template to print before a task's grouped output.")
	rootCmd.PersistentFlags().String("output-group-end", "", "Message template to print after a task's grouped output.")
	rootCmd.PersistentFlags().Bool("output-group-error-only", false, "Swallow output from successful tasks.")
	rootCmd.PersistentFlags().BoolP("parallel", "p", false, "Executes tasks provided on command line in parallel.")
	rootCmd.PersistentFlags().BoolP("silent", "s", false, "Disables echoing.")
	rootCmd.PersistentFlags().String("sort", "", "Changes the order of the tasks when listed. [default|alphanumeric|none].")
	rootCmd.PersistentFlags().Bool("status", false, "Exits with non-zero exit code if any of the given tasks is not up-to-date.")
	rootCmd.PersistentFlags().Bool("summary", false, "Show summary about a task.")
	rootCmd.PersistentFlags().StringP("taskfile", "t", "", "Choose which Taskfile to run. Defaults to \"Taskfile.yml\".")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enables verbose mode.")
	rootCmd.PersistentFlags().Bool("version", false, "Show Task version.")
	rootCmd.PersistentFlags().BoolP("watch", "w", false, "Enables watch of the given task.")
	rootCmd.PersistentFlags().BoolP("yes", "y", false, "Assume \"yes\" as answer to all prompts.")

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
