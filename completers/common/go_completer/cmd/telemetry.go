package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var telemetryCmd = &cobra.Command{
	Use:   "telemetry",
	Short: "manage telemetry data and settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(telemetryCmd).Standalone()

	rootCmd.AddCommand(telemetryCmd)

	carapace.Gen(telemetryCmd).PositionalCompletion(
		carapace.ActionValues("off", "local", "on").StyleF(style.ForKeyword),
	)
}
