package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scoutSyncIndexSettingsCmd = &cobra.Command{
	Use:   "scout:sync-index-settings",
	Short: "Sync your configured index settings with your search engine (Meilisearch)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scoutSyncIndexSettingsCmd).Standalone()

	rootCmd.AddCommand(scoutSyncIndexSettingsCmd)
}
