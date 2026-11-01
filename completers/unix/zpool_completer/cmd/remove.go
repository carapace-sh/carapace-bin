package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove [-npw] pool device...",
	Short:   "remove a device from a pool",
	GroupID: "vdev",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().BoolS("n", "n", false, "dry-run")
	removeCmd.Flags().BoolS("p", "p", false, "display exact values")
	removeCmd.Flags().BoolS("s", "s", false, "stop removal")
	removeCmd.Flags().BoolS("w", "w", false, "wait until removal completes")

	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).PositionalCompletion(
		zfs.ActionPools(),
	)

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return zfs.ActionPoolVdevs(c.Args[0])
			}
			return carapace.ActionValues()
		}),
	)
}
