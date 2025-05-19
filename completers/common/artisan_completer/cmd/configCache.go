package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCacheCmd = &cobra.Command{
	Use:   "config:cache",
	Short: "Create a cache file for faster configuration loading",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCacheCmd).Standalone()

	rootCmd.AddCommand(configCacheCmd)
}
