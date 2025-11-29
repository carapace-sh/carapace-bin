package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_uploadCmd = &cobra.Command{
	Use:     "upload [installable]",
	Short:   "upload specified or nix packages in current project to cache",
	Aliases: []string{"copy"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_uploadCmd).Standalone()

	cache_uploadCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	cache_uploadCmd.Flags().String("to", "", "URI of the cache to copy to")
	cacheCmd.AddCommand(cache_uploadCmd)

	carapace.Gen(cache_uploadCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})

	// TODO positional completion
}
