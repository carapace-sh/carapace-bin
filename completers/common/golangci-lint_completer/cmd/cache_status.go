package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show cache status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_statusCmd).Standalone()

	cacheCmd.AddCommand(cache_statusCmd)
}
