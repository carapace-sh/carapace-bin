package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:     "diff [-FHth] snapshot [snapshot|filesystem]",
	Short:   "show differences between snapshots",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()

	diffCmd.Flags().BoolS("F", "F", false, "display file type")
	diffCmd.Flags().BoolS("H", "H", false, "scripting mode")
	diffCmd.Flags().BoolS("h", "h", false, "display path names without escaping")
	diffCmd.Flags().BoolS("t", "t", false, "display inode change time")

	rootCmd.AddCommand(diffCmd)

	carapace.Gen(diffCmd).PositionalCompletion(
		zfs.ActionSnapshots(),
		zfs.ActionDatasets(zfs.DatasetOpts{Snapshot: true, Filesystem: true}),
	)
}
