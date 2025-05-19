package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/turbo"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run tasks across projects in your monorepo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	addRunFlags(runCmd)
	rootCmd.AddCommand(runCmd)
}

func addRunFlags(cmd *cobra.Command) {
	cmd.Flags().String("cache-dir", "", "Override the filesystem cache directory")
	cmd.Flags().String("cache-workers", "", "Set the number of concurrent cache operations (default 10) [default: 10]")
	cmd.Flags().String("concurrency", "", "Limit the concurrency of task execution. Use 1 for serial (i.e. one-at-a-time) execution")
	cmd.Flags().Bool("continue", false, "Continue execution even if a task exits with an error or non-zero exit code")
	cmd.Flags().String("dry-run", "", "[possible values: text, json]")
	cmd.Flags().String("filter", "", "Use the given selector to specify package(s) to act as entry points")
	cmd.Flags().Bool("force", false, "Ignore the existing cache (to force execution)")
	cmd.Flags().String("global-deps", "", "Specify glob of global filesystem dependencies to be hashed")
	cmd.Flags().String("graph", "", "Generate a graph of the task execution and output to a file when a filename is specified")
	cmd.Flags().String("ignore", "", "Files to ignore when calculating changed files (i.e. --since). Supports globs")
	cmd.Flags().Bool("include-dependencies", false, "Include the dependencies of tasks in execution")
	cmd.Flags().String("log-prefix", "", "Use \"none\" to remove prefixes from task logs")
	cmd.Flags().Bool("no-cache", false, "Avoid saving task results to the cache. Useful for development/watch tasks")
	cmd.Flags().Bool("no-daemon", false, "Run without using turbo's daemon process")
	cmd.Flags().Bool("no-deps", false, "Exclude dependent task consumers from execution")
	cmd.Flags().String("output-logs", "", "Set type of process output logging.")
	cmd.Flags().Bool("parallel", false, "Execute all tasks in parallel")
	cmd.Flags().String("profile", "", "File to write turbo's performance profile output into")
	cmd.Flags().Bool("remote-only", false, "Ignore the local filesystem cache for all tasks")
	cmd.Flags().String("scope", "", "Specify package(s) to act as entry points for task execution. Supports globs")
	cmd.Flags().String("since", "", "Limit/Set scope to changed packages since a mergebase")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"cache-dir": carapace.ActionDirectories(),
		"dry-run":   carapace.ActionValues("text", "json"),
		"filter": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if index := strings.LastIndexAny(c.Value, "[]"); index != -1 && []rune(c.Value)[index] == '[' {
				return git.ActionRefs(git.RefOption{}.Default()).Invoke(c).Prefix(c.Value[:index+1]).Suffix("]").ToA().NoSpace()
			}
			if index := strings.LastIndexAny(c.Value, "{}"); index != -1 && []rune(c.Value)[index] == '{' {
				prefix := c.Value[:index+1]
				c.Value = c.Value[index+1:]
				return carapace.ActionDirectories().Invoke(c).Prefix(prefix).ToA().NoSpace()
			}
			return npm.ActionDependencies().NoSpace()
		}),
		"graph":      carapace.ActionFiles(),
		"log-prefix": carapace.ActionValues("none"),
		"output-logs": carapace.ActionValuesDescribed(
			"errors-only", "sho only errors",
			"full", "show all output",
			"hash-only", "show only turbo-computed task hashes",
			"new-only", "show only new output with only hashes for cached tasks",
			"none", "hide process output",
		),
		"profile": carapace.ActionFiles(),
	})

	carapace.Gen(runCmd).PositionalAnyCompletion(
		turbo.ActionPipelineTasks().FilterArgs(),
	)
}
