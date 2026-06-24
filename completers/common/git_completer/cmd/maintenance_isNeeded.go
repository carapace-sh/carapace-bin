package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var maintenance_isNeededCmd = &cobra.Command{
	Use:   "is-needed",
	Short: "Check if maintenance needs to be run",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(maintenance_isNeededCmd).Standalone()

	maintenance_isNeededCmd.Flags().Bool("auto", false, "check if required thresholds are met without actually running maintenance")
	maintenance_isNeededCmd.Flags().StringSlice("task", nil, "check a specific task")
	maintenanceCmd.AddCommand(maintenance_isNeededCmd)

	carapace.Gen(maintenance_isNeededCmd).FlagCompletion(carapace.ActionMap{
		"task": git.ActionMaintenanceTasks(),
	})
}
