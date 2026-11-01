package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:     "watch",
	Short:   "Run tasks and rerun them when files change",
	Aliases: []string{"w"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(watchCmd).Standalone()

	watchCmd.Flags().Bool("bell", false, "Ring the terminal bell on command completion")
	watchCmd.Flags().StringP("clear", "c", "", "Clear the screen before each run")
	watchCmd.Flags().String("color", "auto", "When to use terminal colours")
	watchCmd.Flags().String("colour", "auto", "When to use terminal colours")
	watchCmd.Flags().StringP("debounce", "d", "50ms", "How long to wait for more changes before running (default: 50ms)")
	watchCmd.Flags().String("delay-run", "", "Wait this long after a change before running the tasks, e.g. 5s")
	watchCmd.Flags().String("emit-events-to", "none", "Configure event emission")
	watchCmd.Flags().StringSliceP("exts", "e", nil, "Only react to files with these extensions, e.g. js,ts")
	watchCmd.Flags().StringSliceP("filter", "f", nil, "Only react to files that match this glob")
	watchCmd.Flags().StringSlice("filter-file", nil, "Files to load filters from")
	watchCmd.Flags().StringSliceP("filter-prog", "J", nil, "[experimental] Filter programs.")
	watchCmd.Flags().String("force-poll", "", "Poll for changes instead of using native file watching")
	watchCmd.Flags().StringSlice("fs-events", []string{"create", "remove", "rename", "modify", "metadata"}, "Filesystem events to filter to")
	watchCmd.Flags().StringSliceP("glob", "g", nil, "Globs to watch instead of the tasks' sources")
	watchCmd.Flags().BoolP("help", "h", false, "Print help")
	watchCmd.Flags().StringSliceP("ignore", "i", nil, "Ignore files that match this glob")
	watchCmd.Flags().StringSlice("ignore-file", nil, "Read ignore patterns from a file, one per line")
	watchCmd.Flags().Bool("ignore-nothing", false, "Ignore nothing")
	watchCmd.Flags().Bool("manual", false, "Show the manual page")
	watchCmd.Flags().StringSlice("map-signal", nil, "Translate a signal that watchexec receives into one it sends to the tasks")
	watchCmd.Flags().BoolS("n", "n", false, "Shorthand for '--shell=none'")
	watchCmd.Flags().Bool("no-default-ignore", false, "Do not apply watchexec's built-in ignores")
	watchCmd.Flags().Bool("no-discover-ignore", false, "Do not look for ignore files at all")
	watchCmd.Flags().Bool("no-global-ignore", false, "Do not apply global ignore files, such as `~/.gitignore` and `~/.config/watchexec/ignore`")
	watchCmd.Flags().Bool("no-meta", false, "Don't emit fs events for metadata changes")
	watchCmd.Flags().Bool("no-project-ignore", false, "Do not apply ignore files in the project, such as `.gitignore` and `.ignore`")
	watchCmd.Flags().Bool("no-vcs-ignore", false, "Do not apply Git and other version control ignore files")
	watchCmd.Flags().BoolP("notify", "N", false, "Alert when commands start and end")
	watchCmd.Flags().StringP("on-busy-update", "o", "do-nothing", "What to do when files change while the tasks are running")
	watchCmd.Flags().Bool("only-emit-events", false, "Only emit events to stdout, run no commands.")
	watchCmd.Flags().String("poll", "", "Poll for changes instead of using native file watching")
	watchCmd.Flags().BoolP("postpone", "p", false, "Wait for the first change before running the tasks")
	watchCmd.Flags().Bool("print-events", false, "Print the events that trigger each run")
	watchCmd.Flags().String("project-origin", "", "Set the project origin")
	watchCmd.Flags().BoolP("restart", "r", false, "Restart the tasks if they are still running when files change")
	watchCmd.Flags().String("shell", "", "Use a different shell")
	watchCmd.Flags().StringP("signal", "s", "", "Signal to send to the running tasks when files change")
	watchCmd.Flags().Bool("skip-deps", false, "Run and watch only the named tasks, not the tasks they depend on")
	watchCmd.Flags().Bool("stdin-quit", false, "Exit when stdin closes")
	watchCmd.Flags().String("stop-signal", "", "Signal that stops the tasks before a restart (default on Unix: SIGTERM)")
	watchCmd.Flags().String("stop-timeout", "10s", "How long to wait for the tasks to stop before killing them (default: 10s)")
	watchCmd.Flags().StringSliceP("task-flag", "t", nil, "Tasks to run")
	watchCmd.Flags().Bool("timings", false, "Print how long the command took to run")
	watchCmd.Flags().StringSliceP("watch", "w", nil, "Watch a file or directory")
	watchCmd.Flags().StringP("watch-file", "F", "", "Read paths to watch from a file, one per line")
	watchCmd.Flags().StringSliceP("watch-non-recursive", "W", nil, "Watch a directory without recursing into it")
	watchCmd.Flags().StringSlice("watchexec-env", nil, "Add env vars to the command")
	watchCmd.Flags().Bool("watchexec-quiet", false, "Don't print starting and stopping messages")
	watchCmd.Flags().String("workdir", "", "Has no effect: mise runs each task from its own `dir`")
	watchCmd.Flags().String("wrap-process", "", "How to wrap the task process: group, session, or none")
	watchCmd.Flag("clear").NoOptDefVal = " "
	watchCmd.Flag("colour").Hidden = true
	watchCmd.Flag("force-poll").NoOptDefVal = " "
	watchCmd.Flag("force-poll").Hidden = true
	watchCmd.Flag("poll").NoOptDefVal = " "
	watchCmd.Flag("task-flag").Hidden = true
	watchCmd.Flag("workdir").Hidden = true
	rootCmd.AddCommand(watchCmd)

	carapace.Gen(watchCmd).FlagCompletion(carapace.ActionMap{
		"clear":               carapace.ActionValues("clear", "reset"),
		"color":               carapace.ActionValues("always", "auto", "never"),
		"emit-events-to":      carapace.ActionValues("environment", "file", "json-file", "json-stdio", "none", "stdio"),
		"filter-file":         carapace.ActionFiles(),
		"fs-events":           carapace.ActionValues("access", "create", "metadata", "modify", "remove", "rename"),
		"ignore-file":         carapace.ActionFiles(),
		"on-busy-update":      carapace.ActionValues("do-nothing", "queue", "restart", "signal"),
		"project-origin":      carapace.ActionDirectories(),
		"watch":               carapace.ActionFiles(),
		"watch-file":          carapace.ActionFiles(),
		"watch-non-recursive": carapace.ActionFiles(),
		"workdir":             carapace.ActionDirectories(),
		"wrap-process":        carapace.ActionValues("group", "none", "session"),
	})

	carapace.Gen(watchCmd).PositionalCompletion(
		mise.ActionTasks(),
	)
}
