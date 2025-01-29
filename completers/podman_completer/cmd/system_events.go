package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_eventsCmd = &cobra.Command{
	Use:   "events [options]",
	Short: "Show podman system events",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_eventsCmd).Standalone()

	system_eventsCmd.Flags().StringSliceP("filter", "f", []string{}, "filter output")
	system_eventsCmd.Flags().String("format", "", "format the output using a Go template")
	system_eventsCmd.Flags().Bool("no-trunc", false, "do not truncate the output")
	system_eventsCmd.Flags().String("since", "", "show all events created since timestamp")
	system_eventsCmd.Flags().Bool("stream", false, "stream events and do not exit when returning the last known event")
	system_eventsCmd.Flags().String("until", "", "show all events until timestamp")
	systemCmd.AddCommand(system_eventsCmd)
}
