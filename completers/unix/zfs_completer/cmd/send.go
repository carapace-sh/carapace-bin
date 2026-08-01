package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:     "send [-DLPVbcehnpsvw] [-R [-X dataset,...]] [[-I|-i] snapshot] snapshot",
	Short:   "generate a send stream",
	GroupID: "send",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sendCmd).Standalone()

	sendCmd.Flags().StringS("I", "I", "", "send all intermediary snapshots")
	sendCmd.Flags().BoolP("backup", "b", false, "send only received property values")
	sendCmd.Flags().BoolP("compressed", "c", false, "use compressed WRITE records")
	sendCmd.Flags().BoolP("dedup", "D", false, "deduplicate (deprecated, accepted for compatibility only)")
	sendCmd.Flags().BoolP("dryrun", "n", false, "dry-run")
	sendCmd.Flags().BoolP("embed", "e", false, "use WRITE_EMBEDDED records")
	sendCmd.Flags().StringArrayP("exclude", "X", nil, "exclude datasets from replication")
	sendCmd.Flags().BoolP("holds", "h", false, "include snapshot holds")
	sendCmd.Flags().StringS("i", "i", "", "send incremental from snapshot")
	sendCmd.Flags().BoolP("large-block", "L", false, "allow large blocks")
	sendCmd.Flags().BoolP("parsable", "P", false, "machine-parsable output")
	sendCmd.Flags().BoolP("proctitle", "V", false, "show data rate in process title")
	sendCmd.Flags().BoolP("props", "p", false, "include dataset properties")
	sendCmd.Flags().BoolP("raw", "w", false, "raw send for encrypted datasets")
	sendCmd.Flags().StringP("redact", "d", "", "redaction bookmark")
	sendCmd.Flags().BoolP("replicate", "R", false, "replicate")
	sendCmd.Flags().BoolP("saved", "S", false, "send saved (partially received) dataset state")
	sendCmd.Flags().BoolP("skip-missing", "s", false, "skip missing snapshots")
	sendCmd.Flags().StringS("t", "t", "", "resume send with token")
	sendCmd.Flags().BoolP("verbose", "v", false, "verbose")

	rootCmd.AddCommand(sendCmd)

	carapace.Gen(sendCmd).FlagCompletion(carapace.ActionMap{
		"I":       zfs.ActionSnapshots(),
		"exclude": zfs.ActionFilesystems().UniqueList(","),
		"i":       zfs.ActionDatasets(zfs.DatasetOpts{Snapshot: true, Bookmark: true}),
		"redact":  zfs.ActionBookmarks(),
	})

	carapace.Gen(sendCmd).PositionalCompletion(
		zfs.ActionDatasets(zfs.DatasetOpts{Snapshot: true, Filesystem: true, Volume: true}),
	)
}
