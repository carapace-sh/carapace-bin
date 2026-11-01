package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:     "clear pool [device...]",
	Short:   "clear device errors",
	GroupID: "fault",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clearCmd).Standalone()

	clearCmd.Flags().Bool("power", false, "power on the device's slot before clearing errors")

	rootCmd.AddCommand(clearCmd)

	carapace.Gen(clearCmd).PositionalCompletion(
		zfs.ActionPools(),
	)

	carapace.Gen(clearCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return zfs.ActionPoolDevices(c.Args[0])
			}
			return carapace.ActionValues()
		}),
	)
}
