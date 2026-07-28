package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_daemonBusQueuesCmd = &cobra.Command{
	Use:   "daemon-bus-queues",
	Short: "Print event bus queue depths per client",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_daemonBusQueuesCmd).Standalone()

	debugCmd.AddCommand(debug_daemonBusQueuesCmd)
}
