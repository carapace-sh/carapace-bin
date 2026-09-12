package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Prunes old cache files in mise",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_pruneCmd).Standalone()

	cacheCmd.AddCommand(cache_pruneCmd)
}
