package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_npx_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "show info about npx cache entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_npx_infoCmd).Standalone()
	cache_npxCmd.AddCommand(cache_npx_infoCmd)
}
