package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eventsCmd = &cobra.Command{
	Use:   "events [options]",
	Short: "Show podman system events",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eventsCmd).Standalone()

	eventsCmd.Flags().StringSliceP("filter", "f", []string{}, "filter output")
	eventsCmd.Flags().String("format", "", "format the output using a Go template")
	eventsCmd.Flags().Bool("no-trunc", false, "do not truncate the output")
	eventsCmd.Flags().String("since", "", "show all events created since timestamp")
	eventsCmd.Flags().Bool("stream", false, "stream events and do not exit when returning the last known event")
	eventsCmd.Flags().String("until", "", "show all events until timestamp")
	rootCmd.AddCommand(eventsCmd)
}
