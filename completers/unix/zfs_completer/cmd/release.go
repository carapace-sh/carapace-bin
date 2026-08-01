package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:     "release [-r] tag snapshot...",
	Short:   "release a hold on a snapshot",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(releaseCmd).Standalone()

	releaseCmd.Flags().BoolS("r", "r", false, "apply recursively")

	rootCmd.AddCommand(releaseCmd)

	carapace.Gen(releaseCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionValues()
			}
			return zfs.ActionSnapshots()
		}),
	)
}
