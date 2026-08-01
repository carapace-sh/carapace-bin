package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:     "rename [-f] [-p] [-u] [-r] filesystem|volume|snapshot target",
	Short:   "rename a dataset",
	GroupID: "dataset",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(renameCmd).Standalone()

	renameCmd.Flags().BoolS("f", "f", false, "force unmount")
	renameCmd.Flags().BoolS("p", "p", false, "create parent datasets")
	renameCmd.Flags().BoolS("r", "r", false, "recursively rename snapshots")
	renameCmd.Flags().BoolS("u", "u", false, "do not remount")

	rootCmd.AddCommand(renameCmd)

	carapace.Gen(renameCmd).PositionalCompletion(
		zfs.ActionDatasets(zfs.DatasetOpts{Filesystem: true, Volume: true, Snapshot: true}),
	)
}
