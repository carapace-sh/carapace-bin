package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Run a task",
	Aliases: []string{"r"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().BoolP("dry-run", "n", false, "Show task commands without executing")
	runCmd.Flags().BoolP("force", "f", false, "Force task execution even if outputs are up to date")
	runCmd.Flags().StringP("jobs", "j", "", "Number of jobs to run in parallel")
	runCmd.Flags().Bool("prefix", false, "Print [task_name] prefix on task stdout/stderr")
	runCmd.Flags().BoolP("quiet", "q", false, "Suppress non-error task output")
	runCmd.Flags().Bool("silent", false, "Suppress all output")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).PositionalCompletion(
		action.ActionTasks(),
	)
}
