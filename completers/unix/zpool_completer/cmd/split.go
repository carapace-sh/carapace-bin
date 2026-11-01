package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:     "split [-gLlnP] [-o property=value]... [-R root] pool newpool [device...]",
	Short:   "split a mirror into a new pool",
	GroupID: "vdev",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(splitCmd).Standalone()

	splitCmd.Flags().BoolS("L", "L", false, "resolve symbolic links")
	splitCmd.Flags().BoolS("P", "P", false, "display full paths")
	splitCmd.Flags().StringS("R", "R", "", "alternate root directory")
	splitCmd.Flags().BoolS("g", "g", false, "display vdev GUIDs")
	splitCmd.Flags().BoolS("l", "l", false, "load encryption keys")
	splitCmd.Flags().BoolS("n", "n", false, "dry-run")
	splitCmd.Flags().StringArrayS("o", "o", nil, "set property")

	rootCmd.AddCommand(splitCmd)

	carapace.Gen(splitCmd).FlagCompletion(carapace.ActionMap{
		"R": carapace.ActionDirectories(),
		"o": zfs.ActionPoolPropertyAssignments(),
	})

	carapace.Gen(splitCmd).PositionalCompletion(
		zfs.ActionPools(),
	)

	carapace.Gen(splitCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 1 {
				return zfs.ActionPoolDevices(c.Args[0])
			}
			return carapace.ActionValues()
		}),
	)
}
