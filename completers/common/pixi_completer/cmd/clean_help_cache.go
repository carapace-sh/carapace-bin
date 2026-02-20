package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clean_help_cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Clean the cache of your system which are touched by pixi",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clean_help_cacheCmd).Standalone()

	clean_helpCmd.AddCommand(clean_help_cacheCmd)
}
