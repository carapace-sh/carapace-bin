package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var routeCacheCmd = &cobra.Command{
	Use:   "route:cache",
	Short: "Create a route cache file for faster route registration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(routeCacheCmd).Standalone()

	rootCmd.AddCommand(routeCacheCmd)
}
