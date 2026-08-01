package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:     "sync [pool...]",
	Short:   "flush all dirty data to disk",
	GroupID: "maintenance",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syncCmd).Standalone()

	rootCmd.AddCommand(syncCmd)

	carapace.Gen(syncCmd).PositionalAnyCompletion(
		zfs.ActionPools(),
	)
}
