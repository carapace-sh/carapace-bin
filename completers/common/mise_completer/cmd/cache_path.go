package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_pathCmd = &cobra.Command{
	Use:     "path",
	Short:   "Print the cache directory path",
	Aliases: []string{"dir"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_pathCmd).Standalone()

	cache_pathCmd.Flags().BoolP("help", "h", false, "Print help")
	cacheCmd.AddCommand(cache_pathCmd)
}
