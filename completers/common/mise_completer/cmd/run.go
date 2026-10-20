package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Run tasks and their dependencies",
	Aliases: []string{"r"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("affected", false, "Run matching tasks only in projects affected by Git changes")
	runCmd.Flags().String("affected-base", "", "Git base revision for `--affected`")
	runCmd.Flags().Bool("affected-committed", false, "Count committed changes between the base and head revisions")
	runCmd.Flags().Bool("affected-explain", false, "Explain why `--affected` selected each project and task")
	runCmd.Flags().String("affected-head", "", "Git head revision for `--affected`")
	runCmd.Flags().Bool("affected-json", false, "Print the affected projects and tasks as JSON without running tasks")
	runCmd.Flags().Bool("affected-uncommitted", false, "Count staged and unstaged changes to tracked files")
	runCmd.Flags().Bool("affected-untracked", false, "Count untracked files that Git does not ignore")
	runCmd.Flags().Bool("all", false, "Open the task picker with tasks from every project in the monorepo")
	runCmd.Flags().StringSlice("allow-env", nil, "Allow a specific env var through (implies `--deny-env` for everything else)")
	runCmd.Flags().StringSlice("allow-net", nil, "Allow network access only to HOST (not supported on any platform)")
	runCmd.Flags().StringSlice("allow-read", nil, "Allow reads from a specific path (implies `--deny-read` for everything else)")
	runCmd.Flags().StringSlice("allow-write", nil, "Allow writes to a specific path (implies `--deny-write` for everything else)")
	runCmd.Flags().StringP("cd", "C", "", "Change to this directory before running the tasks")
	runCmd.Flags().BoolP("continue-on-error", "c", false, "Continue running tasks even if one fails")
	runCmd.Flags().Bool("deny-all", false, "Block reads, writes, network, and env vars")
	runCmd.Flags().Bool("deny-env", false, "Block env var inheritance except PATH, HOME, USER, SHELL, TERM, COLORTERM, and LANG")
	runCmd.Flags().Bool("deny-net", false, "Block all network access")
	runCmd.Flags().Bool("deny-read", false, "Block filesystem reads (system libraries and tool directories stay readable)")
	runCmd.Flags().Bool("deny-write", false, "Block filesystem writes except temporary and device paths (/tmp, /dev)")
	runCmd.Flags().BoolP("dry-run", "n", false, "Print the tasks in the order they would run, without running them")
	runCmd.Flags().BoolP("force", "f", false, "Run the tasks even if their outputs are up to date")
	runCmd.Flags().Bool("fresh-env", false, "Bypass the environment cache and recompute the environment")
	runCmd.Flags().StringP("jobs", "j", "", "Number of tasks to run in parallel (default: the `jobs` setting)")
	runCmd.Flags().Bool("no-cache", false, "Do not use the cache for remote tasks")
	runCmd.Flags().Bool("no-deps", false, "Do not run `[deps]` providers that have `auto = true` before the tasks")
	runCmd.Flags().Bool("no-timing", false, "Hide the elapsed time printed after each task completes")
	runCmd.Flags().Bool("no-timings", false, "Hide the elapsed time printed after each task completes")
	runCmd.Flags().StringP("output", "o", "", "How task output is displayed")
	runCmd.Flags().BoolP("quiet", "q", false, "Do not show extra output")
	runCmd.Flags().BoolP("raw", "r", false, "Connect the tasks directly to stdin, stdout, and stderr instead of reading their output by line")
	runCmd.Flags().StringSlice("secrets", nil, "[experimental] Give the named tasks these secrets (comma-separated); their dependencies and subtasks do not get them")
	runCmd.Flags().Bool("secrets-all", false, "[experimental] Give the named tasks every secret the project can inject; their dependencies and subtasks do not get them")
	runCmd.Flags().StringP("shell", "s", "", "Shell used to run TOML tasks that do not set their own `shell`")
	runCmd.Flags().BoolP("silent", "S", false, "Do not show any output except for errors")
	runCmd.Flags().Bool("skip-deps", false, "Run only the named tasks, not the tasks they depend on")
	runCmd.Flags().Bool("skip-tools", false, "Do not install missing tools before running the tasks")
	runCmd.Flags().String("task-cache", "read-write", "[experimental] Set artifact cache access for this run")
	runCmd.Flags().Bool("task-cache-explain", false, "[experimental] Explain the inputs that produced each task's artifact cache key")
	runCmd.Flags().Bool("task-cache-explain-json", false, "[experimental] Print cache-key input details as JSON Lines without running tasks (requires `--dry-run`)")
	runCmd.Flags().Bool("task-cache-stats", false, "[experimental] Report artifact cache hits, restored bytes, and time saved")
	runCmd.Flags().String("timeout", "", "Stop the run if it takes longer than this, e.g. 30s or 5m")
	runCmd.Flags().Bool("timing", false, "Show the elapsed time after each task completes")
	runCmd.Flags().Bool("timings", false, "Show the elapsed time after each task completes")
	runCmd.Flags().StringSliceP("tool", "t", nil, "Extra tools to use with the tasks, e.g. node@24 python@3.13")
	runCmd.Flag("no-timing").Hidden = true
	runCmd.Flag("timing").Hidden = true
	runCmd.Flag("timings").Hidden = true
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"allow-read":  carapace.ActionFiles(),
		"allow-write": carapace.ActionFiles(),
		"cd":          carapace.ActionDirectories(),
		"secrets":     mise.ActionSecrets(),
		"task-cache":  action.TaskCacheModes(),
		"tool":        mise.ActionToolVersions(),
	})

	carapace.Gen(runCmd).PositionalCompletion(
		mise.ActionTasks(),
	)
}
