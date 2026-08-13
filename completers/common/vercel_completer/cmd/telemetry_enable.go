package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var telemetry_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enables telemetry collection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(telemetry_enableCmd).Standalone()

	telemetryCmd.AddCommand(telemetry_enableCmd)
}
