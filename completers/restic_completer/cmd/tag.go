package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/restic_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Modify tags on snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tagCmd).Standalone()
	tagCmd.Flags().StringSlice("add", nil, "`tags` which will be added to the existing tags in the format `tag[,tag,...]` (can be given multiple times)")
	tagCmd.Flags().StringArrayP("host", "H", nil, "only consider snapshots for this `host`, when no snapshot ID is given (can be specified multiple times)")
	tagCmd.Flags().StringArray("path", nil, "only consider snapshots which include this (absolute) `path`, when no snapshot-ID is given")
	tagCmd.Flags().StringSlice("remove", nil, "`tags` which will be removed from the existing tags in the format `tag[,tag,...]` (can be given multiple times)")
	tagCmd.Flags().StringSlice("set", nil, "`tags` which will replace the existing tags in the format `tag[,tag,...]` (can be given multiple times)")
	tagCmd.Flags().StringSlice("tag", nil, "only consider snapshots which include this `taglist`, when no snapshot-ID is given")
	rootCmd.AddCommand(tagCmd)

	carapace.Gen(tagCmd).FlagCompletion(carapace.ActionMap{
		"add":    action.ActionSnapshotTags(tagCmd).UniqueList(","),
		"host":   action.ActionSnapshotHosts(tagCmd),
		"remove": action.ActionSnapshotTags(tagCmd).UniqueList(","),
		"set":    action.ActionSnapshotTags(tagCmd).UniqueList(","),
		"tag":    action.ActionSnapshotTags(tagCmd).UniqueList(","),
	})

	carapace.Gen(tagCmd).PositionalAnyCompletion(
		action.ActionSnapshotIDs(tagCmd).FilterArgs(),
	)
}
