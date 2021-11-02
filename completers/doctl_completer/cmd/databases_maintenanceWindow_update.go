package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_maintenanceWindow_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the maintenance window for a database cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_maintenanceWindow_updateCmd).Standalone()
	databases_maintenanceWindow_updateCmd.Flags().String("day", "", "The day of the week the maintenance window occurs (e.g. 'tuesday') (required)")
	databases_maintenanceWindow_updateCmd.Flags().String("hour", "", "The hour in UTC when maintenance updates will be applied, in 24 hour format (e.g. '16:00') (required)")
	databases_maintenanceWindowCmd.AddCommand(databases_maintenanceWindow_updateCmd)

	carapace.Gen(databases_maintenanceWindow_updateCmd).FlagCompletion(carapace.ActionMap{
		"day": carapace.ActionValues("monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"),
	})
}
