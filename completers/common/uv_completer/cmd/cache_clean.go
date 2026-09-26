package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var cache_cleanCmd = &cobra.Command{
	Use:     "clean",
	Short:   "Clear the cache, removing all entries or those linked to specific packages",
	Aliases: []string{"clear"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_cleanCmd).Standalone()

	cache_cleanCmd.Flags().Bool("force", false, "Force removal of the cache, ignoring in-use checks")
	cacheCmd.AddCommand(cache_cleanCmd)
	carapace.Gen(cache_cleanCmd).PositionalAnyCompletion(uv.ActionInstalledPackages())
}
