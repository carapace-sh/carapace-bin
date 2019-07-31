package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var offlineCmd = &cobra.Command{
	Use:     "offline [--power|[-ft]] pool device...",
	Short:   "take a device offline",
	GroupID: "fault",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineCmd).Standalone()

	offlineCmd.Flags().BoolS("f", "f", false, "force fault")
	offlineCmd.Flags().Bool("power", false, "power off the device's slot in the storage enclosure (Linux only)")
	offlineCmd.Flags().BoolS("t", "t", false, "temporary offline until reboot")

	rootCmd.AddCommand(offlineCmd)

	carapace.Gen(offlineCmd).PositionalCompletion(
		zfs.ActionPools(),
	)

	carapace.Gen(offlineCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return zfs.ActionPoolDevices(c.Args[0])
			}
			return carapace.ActionValues()
		}),
	)
}
