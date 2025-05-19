package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var maintenance_unregisterCmd = &cobra.Command{
	Use:   "unregister",
	Short: "Remove the current repository from background maintenance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(maintenance_unregisterCmd).Standalone()

	maintenanceCmd.AddCommand(maintenance_unregisterCmd)
}
