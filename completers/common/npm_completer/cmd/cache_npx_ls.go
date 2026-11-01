package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_npx_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list npx cache entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_npx_lsCmd).Standalone()
	cache_npxCmd.AddCommand(cache_npx_lsCmd)
}
