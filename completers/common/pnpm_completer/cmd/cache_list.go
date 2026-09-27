package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the available packages metadata cache. Supports filtering by glob",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_listCmd).Standalone()

	cache_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	cacheCmd.AddCommand(cache_listCmd)
}
