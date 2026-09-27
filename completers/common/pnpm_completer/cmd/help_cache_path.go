package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_cache_pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Prints the path to the cache directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_cache_pathCmd).Standalone()

	help_cacheCmd.AddCommand(help_cache_pathCmd)
}
