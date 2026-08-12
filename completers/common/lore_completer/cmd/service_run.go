package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run this process as the service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_runCmd).Standalone()

	service_runCmd.Flags().BoolP("help", "h", false, "Print help")
	serviceCmd.AddCommand(service_runCmd)
}
