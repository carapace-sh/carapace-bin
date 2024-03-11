package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var maintenance_registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Initialize Git config values",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(maintenance_registerCmd).Standalone()

	maintenanceCmd.AddCommand(maintenance_registerCmd)
}
