package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean <cache-types>...",
	Short: "remove or expire cached data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	rootCmd.AddCommand(cleanCmd)

	carapace.Gen(cleanCmd).PositionalAnyCompletion(
		carapace.ActionValuesDescribed(
			"all", "Delete all cached data from the repositories cache",
			"packages", "Delete packages from the repositories cache",
			"metadata", "Delete the metadata and dbcache from the repositories cache",
			"dbcache", "Delete dbcache from the repositories cache",
			"expire-cache", "Mark the repositories cache as expired",
		),
	)
}
