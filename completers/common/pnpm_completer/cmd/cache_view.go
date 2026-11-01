package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Views information from the specified package's cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_viewCmd).Standalone()

	cache_viewCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	cacheCmd.AddCommand(cache_viewCmd)
}
