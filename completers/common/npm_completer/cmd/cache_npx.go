package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_npxCmd = &cobra.Command{
	Use:   "npx",
	Short: "manage npx cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_npxCmd).Standalone()
	cacheCmd.AddCommand(cache_npxCmd)
}
