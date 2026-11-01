package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_cache_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the available packages metadata cache. Supports filtering by glob",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_cache_listCmd).Standalone()

	help_cacheCmd.AddCommand(help_cache_listCmd)
}
