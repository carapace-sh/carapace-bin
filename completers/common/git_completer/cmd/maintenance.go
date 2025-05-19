package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var maintenanceCmd = &cobra.Command{
	Use:     "maintenance",
	Short:   "Run tasks to optimize Git repository data",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(maintenanceCmd).Standalone()

	rootCmd.AddCommand(maintenanceCmd)
}
