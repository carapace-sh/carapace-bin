package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cachePruneStaleTagsCmd = &cobra.Command{
	Use:   "cache:prune-stale-tags [<store>]",
	Short: "Prune stale cache tags from the cache (Redis only)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cachePruneStaleTagsCmd).Standalone()

	rootCmd.AddCommand(cachePruneStaleTagsCmd)
}
