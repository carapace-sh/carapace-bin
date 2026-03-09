package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean [options] <cache_types>...",
	Short: "remove or invalidate cached data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	rootCmd.AddCommand(cleanCmd)

	carapace.Gen(cleanCmd).PositionalCompletion(
		carapace.ActionValues("all", "packages", "metadata", "dbcache", "expire-cache"),
	)
}
