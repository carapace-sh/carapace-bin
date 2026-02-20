package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_clean_cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Clean the cache of your system which are touched by pixi",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_clean_cacheCmd).Standalone()

	help_cleanCmd.AddCommand(help_clean_cacheCmd)
}
