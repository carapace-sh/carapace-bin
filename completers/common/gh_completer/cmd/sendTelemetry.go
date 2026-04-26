package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sendTelemetryCmd = &cobra.Command{
	Use:    "send-telemetry",
	Short:  "Send telemetry event to GitHub",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sendTelemetryCmd).Standalone()

	rootCmd.AddCommand(sendTelemetryCmd)
}
