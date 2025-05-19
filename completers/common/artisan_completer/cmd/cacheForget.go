package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cacheForgetCmd = &cobra.Command{
	Use:   "cache:forget <key> [<store>]",
	Short: "Remove an item from the cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cacheForgetCmd).Standalone()

	rootCmd.AddCommand(cacheForgetCmd)
}
