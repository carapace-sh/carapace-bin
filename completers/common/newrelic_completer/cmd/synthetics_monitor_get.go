package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_monitor_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a New Relic Synthetics monitor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_monitor_getCmd).Standalone()
	synthetics_monitor_getCmd.Flags().String("monitorId", "", "A New Relic Synthetics monitor ID")
	synthetics_monitorCmd.AddCommand(synthetics_monitor_getCmd)
}
