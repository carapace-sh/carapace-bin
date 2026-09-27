package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes metadata cache for the specified package(s). Supports patterns",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_deleteCmd).Standalone()

	cache_deleteCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	cacheCmd.AddCommand(cache_deleteCmd)
}
