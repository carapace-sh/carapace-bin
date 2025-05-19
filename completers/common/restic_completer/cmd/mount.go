package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/restic_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "Mount the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mountCmd).Standalone()
	mountCmd.Flags().Bool("allow-other", false, "allow other users to access the data in the mounted directory")
	mountCmd.Flags().StringArrayP("host", "H", nil, "only consider snapshots for this host (can be specified multiple times)")
	mountCmd.Flags().Bool("no-default-permissions", false, "for 'allow-other', ignore Unix permissions and allow users to read all snapshot files")
	mountCmd.Flags().Bool("owner-root", false, "use 'root' as the owner of files and dirs")
	mountCmd.Flags().StringArray("path", nil, "only consider snapshots which include this (absolute) `path`")
	mountCmd.Flags().String("snapshot-template", "2006-01-02T15:04:05Z07:00", "set `template` to use for snapshot dirs")
	mountCmd.Flags().StringSlice("tag", nil, "only consider snapshots which include this `taglist`")
	rootCmd.AddCommand(mountCmd)

	carapace.Gen(mountCmd).FlagCompletion(carapace.ActionMap{
		"host": action.ActionSnapshotHosts(mountCmd),
		"path": carapace.ActionFiles(),
		"tag":  action.ActionSnapshotTags(mountCmd).UniqueList(","),
	})
}
