package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list cache entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_lsCmd).Standalone()
	cacheCmd.AddCommand(cache_lsCmd)
}
