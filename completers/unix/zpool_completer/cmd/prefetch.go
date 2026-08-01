package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var prefetchCmd = &cobra.Command{
	Use:     "prefetch [-t type] pool",
	Short:   "prefetch pool data",
	GroupID: "maintenance",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prefetchCmd).Standalone()

	prefetchCmd.Flags().StringS("t", "t", "", "data types to prefetch")

	rootCmd.AddCommand(prefetchCmd)

	carapace.Gen(prefetchCmd).FlagCompletion(carapace.ActionMap{
		"t": carapace.ActionValues("ddt", "brt"),
	})

	carapace.Gen(prefetchCmd).PositionalCompletion(
		zfs.ActionPools(),
	)
}
