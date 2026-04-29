package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run [task]",
	Aliases: []string{"r"},
	Short:   "Run a mise task",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()
	runCmd.Flags().BoolP("force", "f", false, "Force run even if dependencies are not needed")
	runCmd.Flags().StringP("jobs", "j", "", "Number of tasks to run in parallel")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).PositionalCompletion(
		mise.ActionTasks(),
	)
}
