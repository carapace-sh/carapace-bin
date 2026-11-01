package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var inheritCmd = &cobra.Command{
	Use:     "inherit [-rS] property dataset...",
	Short:   "clear property to inherit from parent",
	GroupID: "property",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inheritCmd).Standalone()

	inheritCmd.Flags().BoolS("S", "S", false, "revert to received value")
	inheritCmd.Flags().BoolS("r", "r", false, "apply recursively")

	rootCmd.AddCommand(inheritCmd)

	carapace.Gen(inheritCmd).PositionalCompletion(
		zfs.ActionProperties(),
	)

	carapace.Gen(inheritCmd).PositionalAnyCompletion(
		zfs.ActionDatasets(zfs.DatasetOpts{Filesystem: true, Volume: true, Snapshot: true}),
	)
}
