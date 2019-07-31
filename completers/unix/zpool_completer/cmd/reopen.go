package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var reopenCmd = &cobra.Command{
	Use:     "reopen [-n] [pool]...",
	Short:   "reopen all vdevs",
	GroupID: "fault",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reopenCmd).Standalone()

	reopenCmd.Flags().BoolS("n", "n", false, "do not restart an in-progress scrub")

	rootCmd.AddCommand(reopenCmd)

	carapace.Gen(reopenCmd).PositionalAnyCompletion(
		zfs.ActionPools(),
	)
}
