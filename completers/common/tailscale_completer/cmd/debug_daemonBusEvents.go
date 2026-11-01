package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_daemonBusEventsCmd = &cobra.Command{
	Use:   "daemon-bus-events",
	Short: "Watch events on the tailscaled bus",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_daemonBusEventsCmd).Standalone()

	debugCmd.AddCommand(debug_daemonBusEventsCmd)
}
