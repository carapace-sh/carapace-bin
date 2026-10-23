package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cacheCatCmd = &cobra.Command{
	Use:   "cache-cat <paths>",
	Short: "Print a zig-cache manifest file as zon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cacheCatCmd).Standalone()

	rootCmd.AddCommand(cacheCatCmd)

	cacheCatCmd.Flags().BoolP("help", "h", false, "Print help")

	carapace.Gen(cacheCatCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
