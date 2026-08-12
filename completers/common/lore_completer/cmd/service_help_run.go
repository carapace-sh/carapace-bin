package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_help_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run this process as the service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_help_runCmd).Standalone()

	service_helpCmd.AddCommand(service_help_runCmd)
}
