package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var replaceCmd = &cobra.Command{
	Use:     "replace [-fsw] [-o property=value] pool device [new_device]",
	Short:   "replace a device in a pool",
	GroupID: "vdev",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(replaceCmd).Standalone()

	replaceCmd.Flags().BoolS("f", "f", false, "force replace")
	replaceCmd.Flags().StringArrayS("o", "o", nil, "set property")
	replaceCmd.Flags().BoolS("s", "s", false, "do not resilver")
	replaceCmd.Flags().BoolS("w", "w", false, "wait until resilvering completes")

	rootCmd.AddCommand(replaceCmd)

	carapace.Gen(replaceCmd).FlagCompletion(carapace.ActionMap{
		"o": zfs.ActionPoolPropertyAssignments(),
	})

	carapace.Gen(replaceCmd).PositionalCompletion(
		zfs.ActionPools(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return zfs.ActionPoolDevices(c.Args[0])
		}),
		carapace.ActionFiles(),
	)
}
