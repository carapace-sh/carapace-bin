package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var serviceWatchdogsCmd = &cobra.Command{
	Use:   "service-watchdogs",
	Short: "Get/set service watchdog state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceWatchdogsCmd).Standalone()

	rootCmd.AddCommand(serviceWatchdogsCmd)

	carapace.Gen(serviceWatchdogsCmd).PositionalCompletion(
		carapace.ActionValues("yes", "no"),
	)
}
