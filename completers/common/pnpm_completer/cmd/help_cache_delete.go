package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_cache_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes metadata cache for the specified package(s). Supports patterns",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_cache_deleteCmd).Standalone()

	help_cacheCmd.AddCommand(help_cache_deleteCmd)
}
