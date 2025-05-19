package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var maintenance_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Halt the background maintenance schedule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(maintenance_stopCmd).Standalone()

	maintenanceCmd.AddCommand(maintenance_stopCmd)
}
