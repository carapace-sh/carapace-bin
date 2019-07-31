package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var eventsCmd = &cobra.Command{
	Use:     "events [-vHfc] [pool]",
	Short:   "display pool events",
	GroupID: "monitoring",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eventsCmd).Standalone()

	eventsCmd.Flags().BoolS("H", "H", false, "scripting mode")
	eventsCmd.Flags().BoolS("c", "c", false, "clear all previous events")
	eventsCmd.Flags().BoolS("f", "f", false, "follow new events")
	eventsCmd.Flags().BoolS("v", "v", false, "verbose")

	rootCmd.AddCommand(eventsCmd)

	carapace.Gen(eventsCmd).PositionalCompletion(
		zfs.ActionPools(),
	)
}
