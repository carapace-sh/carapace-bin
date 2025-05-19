package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeCacheTableCmd = &cobra.Command{
	Use:     "make:cache-table",
	Short:   "Create a migration for the cache database table",
	Aliases: []string{"cache:table"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeCacheTableCmd).Standalone()

	rootCmd.AddCommand(makeCacheTableCmd)
}
