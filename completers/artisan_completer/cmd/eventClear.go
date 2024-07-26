package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eventClearCmd = &cobra.Command{
	Use:   "event:clear",
	Short: "Clear all cached events and listeners",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eventClearCmd).Standalone()

	rootCmd.AddCommand(eventClearCmd)
}
