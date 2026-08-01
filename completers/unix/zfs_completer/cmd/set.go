package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:     "set property=value dataset...",
	Short:   "set property on datasets",
	GroupID: "property",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()

	setCmd.Flags().BoolS("u", "u", false, "update property without mounting or sharing")

	rootCmd.AddCommand(setCmd)

	carapace.Gen(setCmd).PositionalCompletion(
		zfs.ActionPropertyAssignments(),
	)

	carapace.Gen(setCmd).PositionalAnyCompletion(
		carapace.Batch(
			zfs.ActionPropertyAssignments(),
			zfs.ActionDatasets(zfs.DatasetOpts{Filesystem: true, Volume: true, Snapshot: true}),
		).ToA(),
	)
}
