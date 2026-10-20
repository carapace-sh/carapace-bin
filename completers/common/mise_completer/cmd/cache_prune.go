package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var cache_pruneCmd = &cobra.Command{
	Use:     "prune",
	Short:   "Remove cache files that have not been accessed recently",
	Aliases: []string{"p"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_pruneCmd).Standalone()

	cache_pruneCmd.Flags().Bool("dry-run", false, "Show what would be pruned without deleting anything")
	cache_pruneCmd.Flags().BoolP("help", "h", false, "Print help")
	cache_pruneCmd.Flags().CountP("verbose", "v", "Show pruned files")
	cacheCmd.AddCommand(cache_pruneCmd)

	carapace.Gen(cache_pruneCmd).PositionalAnyCompletion(
		mise.ActionTools(),
	)
}
