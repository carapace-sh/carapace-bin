package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var holdCmd = &cobra.Command{
	Use:     "hold [-r] tag snapshot...",
	Short:   "add a hold on a snapshot",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(holdCmd).Standalone()

	holdCmd.Flags().BoolS("r", "r", false, "apply recursively")

	rootCmd.AddCommand(holdCmd)

	carapace.Gen(holdCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionValues()
			}
			return zfs.ActionSnapshots()
		}),
	)
}
