package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_cleanCmd).Standalone()

	cacheCmd.AddCommand(cache_cleanCmd)
}
