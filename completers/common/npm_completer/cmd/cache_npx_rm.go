package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_npx_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove npx cache entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_npx_rmCmd).Standalone()
	cache_npxCmd.AddCommand(cache_npx_rmCmd)
}
