package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_maintenanceWindowCmd = &cobra.Command{
	Use:   "maintenance-window",
	Short: "Display commands for scheduling automatic maintenance on your database cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_maintenanceWindowCmd).Standalone()
	databasesCmd.AddCommand(databases_maintenanceWindowCmd)
}
