package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:     "destroy [-dfnpRrv] dataset",
	Short:   "destroy a dataset",
	GroupID: "dataset",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(destroyCmd).Standalone()

	destroyCmd.Flags().BoolS("R", "R", false, "recursively destroy dependents")
	destroyCmd.Flags().BoolS("d", "d", false, "defer snapshot destruction")
	destroyCmd.Flags().BoolS("f", "f", false, "force unmount")
	destroyCmd.Flags().BoolS("n", "n", false, "dry-run")
	destroyCmd.Flags().BoolS("p", "p", false, "machine-parsable verbose output")
	destroyCmd.Flags().BoolS("r", "r", false, "recursively destroy children")
	destroyCmd.Flags().BoolS("v", "v", false, "verbose")

	rootCmd.AddCommand(destroyCmd)

	carapace.Gen(destroyCmd).PositionalCompletion(
		zfs.ActionDatasets(zfs.DatasetOpts{Filesystem: true, Volume: true, Snapshot: true, Bookmark: true}),
	)
}
