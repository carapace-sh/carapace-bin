package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clearCacheCmd = &cobra.Command{
	Use:     "clear-cache",
	Short:   "Clears composer's internal package cache",
	Aliases: []string{"cc", "clearcache"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clearCacheCmd).Standalone()

	clearCacheCmd.Flags().Bool("gc", false, "Only run garbage collection, not a full cache clear")

	rootCmd.AddCommand(clearCacheCmd)
}
