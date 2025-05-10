package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/restic_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a file, a directory or restic IDs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(findCmd).Standalone()
	findCmd.Flags().Bool("blob", false, "pattern is a blob-ID")
	findCmd.Flags().StringArrayP("host", "H", nil, "only consider snapshots for this `host`, when no snapshot ID is given (can be specified multiple times)")
	findCmd.Flags().BoolP("ignore-case", "i", false, "ignore case for pattern")
	findCmd.Flags().BoolP("long", "l", false, "use a long listing format showing size and mode")
	findCmd.Flags().StringP("newest", "N", "", "newest modification date/time")
	findCmd.Flags().StringP("oldest", "O", "", "oldest modification date/time")
	findCmd.Flags().Bool("pack", false, "pattern is a pack-ID")
	findCmd.Flags().StringArray("path", nil, "only consider snapshots which include this (absolute) `path`, when no snapshot-ID is given")
	findCmd.Flags().Bool("show-pack-id", false, "display the pack-ID the blobs belong to (with --blob or --tree)")
	findCmd.Flags().StringArrayP("snapshot", "s", nil, "snapshot `id` to search in (can be given multiple times)")
	findCmd.Flags().StringSlice("tag", nil, "only consider snapshots which include this `taglist`, when no snapshot-ID is given")
	findCmd.Flags().Bool("tree", false, "pattern is a tree-ID")
	rootCmd.AddCommand(findCmd)

	carapace.Gen(findCmd).FlagCompletion(carapace.ActionMap{
		"host":     action.ActionSnapshotHosts(findCmd).UniqueList(","),
		"newest":   time.ActionDateTime(time.DateTimeOpts{}),
		"oldest":   time.ActionDateTime(time.DateTimeOpts{}),
		"path":     carapace.ActionFiles(),
		"snapshot": action.ActionSnapshotIDs(findCmd),
		"tag":      action.ActionSnapshotTags(findCmd).UniqueList(","),
	})
}
