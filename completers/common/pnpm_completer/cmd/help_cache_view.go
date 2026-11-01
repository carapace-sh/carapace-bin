package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_cache_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Views information from the specified package's cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_cache_viewCmd).Standalone()

	help_cacheCmd.AddCommand(help_cache_viewCmd)
}
