package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "reload cached images.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_reloadCmd).Standalone()

	cacheCmd.AddCommand(cache_reloadCmd)
}
