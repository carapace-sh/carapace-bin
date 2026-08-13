package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var telemetry_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Shows whether telemetry collection is enabled or disabled",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(telemetry_statusCmd).Standalone()

	telemetryCmd.AddCommand(telemetry_statusCmd)
}
