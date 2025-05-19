package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/restic_completer/cmd/action"
	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Print a backed-up file to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpCmd).Standalone()
	dumpCmd.Flags().StringP("archive", "a", "tar", "set archive `format` as \"tar\" or \"zip\"")
	dumpCmd.Flags().StringArrayP("host", "H", nil, "only consider snapshots for this host when the snapshot ID is \"latest\" (can be specified multiple times)")
	dumpCmd.Flags().StringArray("path", nil, "only consider snapshots which include this (absolute) `path` for snapshot ID \"latest\"")
	dumpCmd.Flags().StringSlice("tag", nil, "only consider snapshots which include this `taglist` for snapshot ID \"latest\"")
	rootCmd.AddCommand(dumpCmd)

	carapace.Gen(dumpCmd).FlagCompletion(carapace.ActionMap{
		"archive": carapace.ActionValues("tar", "zip"),
		"host":    action.ActionSnapshotHosts(dumpCmd),
		"path":    carapace.ActionFiles(),
		"tag":     action.ActionSnapshotTags(dumpCmd).UniqueList(","),
	})

	carapace.Gen(dumpCmd).PositionalCompletion(
		carapace.Batch(
			action.ActionSnapshotIDs(dumpCmd),
			carapace.ActionValues("latest"),
		).ToA(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionSnapshotPaths(dumpCmd, c.Args[0]).Invoke(c).ToMultiPartsA("/")
		}),
	)
}
