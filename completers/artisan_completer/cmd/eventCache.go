package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eventCacheCmd = &cobra.Command{
	Use:   "event:cache",
	Short: "Discover and cache the application's events and listeners",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eventCacheCmd).Standalone()

	rootCmd.AddCommand(eventCacheCmd)
}
