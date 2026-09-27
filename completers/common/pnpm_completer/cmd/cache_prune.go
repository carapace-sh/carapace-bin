package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Deletes registry metadata cache directories that this version of pnpm can no longer read",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_pruneCmd).Standalone()

	cache_pruneCmd.Flags().Bool("dry-run", false, "Lists what would be deleted without removing anything")
	cache_pruneCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	cacheCmd.AddCommand(cache_pruneCmd)
}
