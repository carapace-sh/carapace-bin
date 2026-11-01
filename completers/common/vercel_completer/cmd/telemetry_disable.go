package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var telemetry_disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disables telemetry collection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(telemetry_disableCmd).Standalone()

	telemetryCmd.AddCommand(telemetry_disableCmd)
}
