package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/restic_completer/cmd/action"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Scan the repository and show basic statistics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statsCmd).Standalone()
	statsCmd.Flags().StringArrayP("host", "H", nil, "only consider snapshots with the given `host` (can be specified multiple times)")
	statsCmd.Flags().String("mode", "restore-size", "counting mode: restore-size (default), files-by-contents, blobs-per-file or raw-data")
	statsCmd.Flags().StringArray("path", nil, "only consider snapshots which include this (absolute) `path` (can be specified multiple times)")
	statsCmd.Flags().StringSlice("tag", nil, "only consider snapshots which include this `taglist` in the format `tag[,tag,...]` (can be specified multiple times)")
	rootCmd.AddCommand(statsCmd)

	carapace.Gen(statsCmd).FlagCompletion(carapace.ActionMap{
		"host": action.ActionSnapshotHosts(statsCmd),
		"mode": carapace.ActionValues("restore-size", "files-by-contents", "blobs-per-file", "raw-data"),
		"path": carapace.ActionFiles(),
		"tag":  action.ActionSnapshotTags(statsCmd).UniqueList(","),
	})

	carapace.Gen(statsCmd).PositionalAnyCompletion(
		action.ActionSnapshotIDs(statsCmd).FilterArgs(),
	)
}
