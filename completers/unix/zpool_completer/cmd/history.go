package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:     "history [-il] [pool...]",
	Short:   "display pool command history",
	GroupID: "monitoring",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(historyCmd).Standalone()

	historyCmd.Flags().BoolS("i", "i", false, "display internal events")
	historyCmd.Flags().BoolS("l", "l", false, "long format")

	rootCmd.AddCommand(historyCmd)

	carapace.Gen(historyCmd).PositionalAnyCompletion(
		zfs.ActionPools(),
	)
}
