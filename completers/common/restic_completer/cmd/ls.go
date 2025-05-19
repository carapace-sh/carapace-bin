package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/restic_completer/cmd/action"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List files in a snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsCmd).Standalone()
	lsCmd.Flags().StringArrayP("host", "H", nil, "only consider snapshots for this `host`, when no snapshot ID is given (can be specified multiple times)")
	lsCmd.Flags().BoolP("long", "l", false, "use a long listing format showing size and mode")
	lsCmd.Flags().StringArray("path", nil, "only consider snapshots which include this (absolute) `path`, when no snapshot ID is given")
	lsCmd.Flags().Bool("recursive", false, "include files in subfolders of the listed directories")
	lsCmd.Flags().StringSlice("tag", nil, "only consider snapshots which include this `taglist`, when no snapshot ID is given")
	rootCmd.AddCommand(lsCmd)

	carapace.Gen(lsCmd).FlagCompletion(carapace.ActionMap{
		"host": action.ActionSnapshotHosts(lsCmd),
		"path": carapace.ActionFiles(),
		"tag":  action.ActionSnapshotTags(lsCmd).UniqueList(","),
	})

	carapace.Gen(lsCmd).PositionalCompletion(
		action.ActionSnapshotIDs(lsCmd),
	)

	carapace.Gen(lsCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionSnapshotPaths(lsCmd, c.Args[0]).Invoke(c).ToMultiPartsA("/")
		}),
	)
}
