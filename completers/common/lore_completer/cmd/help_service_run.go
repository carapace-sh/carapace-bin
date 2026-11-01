package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_service_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run this process as the service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_service_runCmd).Standalone()

	help_serviceCmd.AddCommand(help_service_runCmd)
}
