package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var detachCmd = &cobra.Command{
	Use:     "detach pool device",
	Short:   "detach a device from a mirror",
	GroupID: "vdev",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detachCmd).Standalone()

	rootCmd.AddCommand(detachCmd)

	carapace.Gen(detachCmd).PositionalCompletion(
		zfs.ActionPools(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return zfs.ActionPoolDevices(c.Args[0])
		}),
	)
}
