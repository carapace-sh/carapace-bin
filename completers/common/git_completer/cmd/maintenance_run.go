package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var maintenance_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start running maintenance on the current repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(maintenance_runCmd).Standalone()

	maintenance_runCmd.Flags().Bool("auto", false, "run tasks based on the state of the repository")
	maintenance_runCmd.Flags().Bool("quiet", false, "don't report progress or other information to stderr")
	maintenance_runCmd.Flags().Bool("schedule", false, "run tasks based on frequency")
	maintenance_runCmd.Flags().StringSlice("task", nil, "run a specific task")
	maintenanceCmd.AddCommand(maintenance_runCmd)

	carapace.Gen(maintenance_runCmd).FlagCompletion(carapace.ActionMap{
		"task": git.ActionMaintenanceTasks(),
	})
}
