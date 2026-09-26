package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Prune dangling cache entries and cached environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_pruneCmd).Standalone()

	cache_pruneCmd.Flags().Bool("ci", false, "Optimize the cache for persistence in a continuous integration environment, like GitHub Actions")
	cache_pruneCmd.Flags().Bool("force", false, "Force removal of the cache, ignoring in-use checks")
	cacheCmd.AddCommand(cache_pruneCmd)
}
