package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove the shared cache files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_cleanCmd).Standalone()

	cache_cleanCmd.Flags().Bool("all", false, "Remove both the global cache files and the local cache files of the current project")
	cache_cleanCmd.Flags().Bool("mirror", false, "Remove the global cache files instead of the local cache files")
	cacheCmd.AddCommand(cache_cleanCmd)
}
