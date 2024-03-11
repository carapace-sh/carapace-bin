package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_monitor_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a New Relic Synthetics Monitor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_monitor_searchCmd).Standalone()
	synthetics_monitor_searchCmd.Flags().StringP("name", "n", "", "search for results matching the given Synthetics monitor name")
	synthetics_monitorCmd.AddCommand(synthetics_monitor_searchCmd)
}
