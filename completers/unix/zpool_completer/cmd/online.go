package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var onlineCmd = &cobra.Command{
	Use:     "online [--power] [-e] pool device...",
	Short:   "bring a device online",
	GroupID: "fault",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(onlineCmd).Standalone()

	onlineCmd.Flags().BoolS("e", "e", false, "expand device to use all available space")
	onlineCmd.Flags().Bool("power", false, "power on the device's slot in the storage enclosure and wait for it to appear (Linux only)")

	rootCmd.AddCommand(onlineCmd)

	carapace.Gen(onlineCmd).PositionalCompletion(
		zfs.ActionPools(),
	)

	carapace.Gen(onlineCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return zfs.ActionPoolDevices(c.Args[0])
			}
			return carapace.ActionValues()
		}),
	)
}
