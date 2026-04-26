package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var telemetryCmd = &cobra.Command{
	Use:    "telemetry",
	Short:  "Information about telemetry in gh",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(telemetryCmd).Standalone()

	rootCmd.AddCommand(telemetryCmd)
}
