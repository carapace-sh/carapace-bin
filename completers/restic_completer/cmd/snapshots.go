package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/restic_completer/cmd/action"
	"github.com/spf13/cobra"
)

var snapshotsCmd = &cobra.Command{
	Use:   "snapshots",
	Short: "List all snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshotsCmd).Standalone()
	snapshotsCmd.Flags().BoolP("compact", "c", false, "use compact output format")
	snapshotsCmd.Flags().StringP("group-by", "g", "", "string for grouping snapshots by host,paths,tags")
	snapshotsCmd.Flags().StringArrayP("host", "H", nil, "only consider snapshots for this `host` (can be specified multiple times)")
	snapshotsCmd.Flags().Bool("last", false, "only show the last snapshot for each host and path")
	snapshotsCmd.Flags().Int("latest", 0, "only show the last `n` snapshots for each host and path")
	snapshotsCmd.Flags().StringArray("path", nil, "only consider snapshots for this `path` (can be specified multiple times)")
	snapshotsCmd.Flags().StringSlice("tag", nil, "only consider snapshots which include this `taglist` in the format `tag[,tag,...]` (can be specified multiple times)")
	rootCmd.AddCommand(snapshotsCmd)

	carapace.Gen(snapshotsCmd).FlagCompletion(carapace.ActionMap{
		"group-by": carapace.ActionValues("host", "paths", "tags").UniqueList(","),
		"host":     action.ActionSnapshotHosts(snapshotsCmd),
		"path":     carapace.ActionFiles(),
		"tag":      action.ActionSnapshotTags(snapshotsCmd).UniqueList(","),
	})

	carapace.Gen(snapshotsCmd).PositionalAnyCompletion(
		action.ActionSnapshotIDs(snapshotsCmd).FilterArgs(),
	)
}
