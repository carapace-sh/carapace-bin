package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove cached files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_rmCmd).Standalone()
	cache_rmCmd.Flags().BoolP("all", "a", false, "remove all cached files")
	cacheCmd.AddCommand(cache_rmCmd)
}
