package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/restic_completer/cmd/action"
	"github.com/spf13/cobra"
)

var forgetCmd = &cobra.Command{
	Use:   "forget",
	Short: "Remove snapshots from the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forgetCmd).Standalone()
	forgetCmd.Flags().IntP("keep-last", "l", 0, "keep the last `n` snapshots")
	forgetCmd.Flags().IntP("keep-hourly", "H", 0, "keep the last `n` hourly snapshots")
	forgetCmd.Flags().IntP("keep-daily", "d", 0, "keep the last `n` daily snapshots")
	forgetCmd.Flags().IntP("keep-weekly", "w", 0, "keep the last `n` weekly snapshots")
	forgetCmd.Flags().IntP("keep-monthly", "m", 0, "keep the last `n` monthly snapshots")
	forgetCmd.Flags().StringSlice("keep-tag", nil, "keep snapshots with this `taglist` (can be specified multiple times)")
	forgetCmd.Flags().IntP("keep-yearly", "y", 0, "keep the last `n` yearly snapshots")
	forgetCmd.Flags().BoolP("compact", "c", false, "use compact output format")
	forgetCmd.Flags().BoolP("dry-run", "n", false, "do not delete anything, just print what would be done")
	forgetCmd.Flags().StringP("group-by", "g", "host,paths", "string for grouping snapshots by host,paths,tags")
	forgetCmd.Flags().StringArray("host", nil, "only consider snapshots with the given `host` (can be specified multiple times)")
	forgetCmd.Flags().StringArray("hostname", nil, "only consider snapshots with the given `hostname` (can be specified multiple times)")
	forgetCmd.Flags().String("keep-within", "", "keep snapshots that are newer than `duration` (eg. 1y5m7d2h) relative to the latest snapshot")
	forgetCmd.Flags().String("keep-within-daily", "", "keep daily snapshots that are newer than `duration` (eg. 1y5m7d2h) relative to the latest snapshot")
	forgetCmd.Flags().String("keep-within-hourly", "", "keep hourly snapshots that are newer than `duration` (eg. 1y5m7d2h) relative to the latest snapshot")
	forgetCmd.Flags().String("keep-within-monthly", "", "keep monthly snapshots that are newer than `duration` (eg. 1y5m7d2h) relative to the latest snapshot")
	forgetCmd.Flags().String("keep-within-weekly", "", "keep weekly snapshots that are newer than `duration` (eg. 1y5m7d2h) relative to the latest snapshot")
	forgetCmd.Flags().String("keep-within-yearly", "", "keep yearly snapshots that are newer than `duration` (eg. 1y5m7d2h) relative to the latest snapshot")
	forgetCmd.Flags().String("max-repack-size", "", "maximum `size` to repack (allowed suffixes: k/K, m/M, g/G, t/T)")
	forgetCmd.Flags().String("max-unused", "5%", "tolerate given `limit` of unused data (absolute value in bytes with suffixes k/K, m/M, g/G, t/T, a value in % or the word 'unlimited')")
	forgetCmd.Flags().StringArray("path", nil, "only consider snapshots which include this (absolute) `path` (can be specified multiple times)")
	forgetCmd.Flags().Bool("prune", false, "automatically run the 'prune' command if snapshots have been removed")
	forgetCmd.Flags().Bool("repack-cacheable-only", false, "only repack packs which are cacheable")
	forgetCmd.Flags().StringSlice("tag", nil, "only consider snapshots which include this `taglist` in the format `tag[,tag,...]` (can be specified multiple times)")
	rootCmd.AddCommand(forgetCmd)

	carapace.Gen(forgetCmd).FlagCompletion(carapace.ActionMap{
		"group-by": carapace.ActionValues("host", "paths", "tags").UniqueList(","),
		"path":     carapace.ActionFiles(),
		"tag":      action.ActionSnapshotTags(forgetCmd).UniqueList(","),
	})

	carapace.Gen(forgetCmd).PositionalAnyCompletion(
		action.ActionSnapshotIDs(forgetCmd).FilterArgs(),
	)
}
