package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_help_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Deletes registry metadata cache directories that this version of pnpm can no longer read",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_help_pruneCmd).Standalone()

	cache_helpCmd.AddCommand(cache_help_pruneCmd)
}
