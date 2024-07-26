package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var viewCacheCmd = &cobra.Command{
	Use:   "view:cache",
	Short: "Compile all of the application's Blade templates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(viewCacheCmd).Standalone()

	rootCmd.AddCommand(viewCacheCmd)
}
