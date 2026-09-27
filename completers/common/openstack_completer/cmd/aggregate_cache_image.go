package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aggregate_cache_imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Request image caching for aggregate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aggregate_cache_imageCmd).Standalone()

	aggregate_cacheCmd.AddCommand(aggregate_cache_imageCmd)
}
