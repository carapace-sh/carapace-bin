package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eventListCmd = &cobra.Command{
	Use:   "event:list [--event [EVENT]]",
	Short: "List the application's events and listeners",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eventListCmd).Standalone()

	eventListCmd.Flags().String("event", "", "Filter the events by name")
	rootCmd.AddCommand(eventListCmd)
}
