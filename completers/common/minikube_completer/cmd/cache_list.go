package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available images from the local cache.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_listCmd).Standalone()
	cache_listCmd.Flags().String("format", "{{.CacheImage}}", "Go template format string for the cache list output.")
	cacheCmd.AddCommand(cache_listCmd)
}
