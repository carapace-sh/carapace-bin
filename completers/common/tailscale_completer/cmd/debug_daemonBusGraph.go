package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_daemonBusGraphCmd = &cobra.Command{
	Use:   "daemon-bus-graph",
	Short: "Print graph for the tailscaled bus",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_daemonBusGraphCmd).Standalone()

	debug_daemonBusGraphCmd.Flags().String("format", "", "output format")
	debugCmd.AddCommand(debug_daemonBusGraphCmd)

	carapace.Gen(debug_daemonBusGraphCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "dot"),
	})
}
