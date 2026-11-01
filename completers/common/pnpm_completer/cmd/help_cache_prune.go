package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_cache_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Deletes registry metadata cache directories that this version of pnpm can no longer read",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_cache_pruneCmd).Standalone()

	help_cacheCmd.AddCommand(help_cache_pruneCmd)
}
