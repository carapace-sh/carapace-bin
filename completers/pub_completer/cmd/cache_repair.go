package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cache_repairCmd = &cobra.Command{
	Use:   "repair",
	Short: "Reinstall cached packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_repairCmd).Standalone()

	cache_repairCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	cacheCmd.AddCommand(cache_repairCmd)
}
