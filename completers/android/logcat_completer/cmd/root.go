package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "logcat",
	Short: "Dump and stream system log messages",
	Long:  "https://cs.android.com/android/platform/superproject/main/+/main:system/logging/logcat/logcat.cpp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("binary", "B", false, "output the log in binary")
	rootCmd.Flags().StringSliceP("buffer", "b", nil, "request alternate ring `BUFFER`S (multiple or comma separated allowed)")
	rootCmd.Flags().StringP("buffer-size", "G", "", "set size of a ring buffer in logd (`SIZE` may be suffixed with K or M)")
	rootCmd.Flags().BoolP("clear", "c", false, "clear (flush) the entire log and exit")
	rootCmd.Flags().BoolP("dividers", "D", false, "print dividers between each log buffer")
	rootCmd.Flags().BoolP("dump", "d", false, "dump the log and then exit (don't block)")
	rootCmd.Flags().StringP("file", "f", "", "log to `FILE` instead of stdout")
	rootCmd.Flags().StringP("format", "v", "", "set log print `FORMAT`")
	rootCmd.Flags().BoolP("get-buffer-size", "g", false, "get size of the ring buffers within logd")
	rootCmd.Flags().BoolP("get-prune", "p", false, "get prune rules")
	rootCmd.Flags().String("id", "", "clears the associated files if the signature `ID` for logging to file changes")
	rootCmd.Flags().BoolP("last", "L", false, "dump logs from prior to last reboot from pstore")
	rootCmd.Flags().IntP("max-count", "m", 0, "exit after printing `COUNT` lines")
	rootCmd.Flags().String("pid", "", "only print logs from the given `PID`")
	rootCmd.Flags().Bool("print", false, "with --regex and --max-count, print all messages even if they do not match the regex")
	rootCmd.Flags().Bool("proto", false, "output the log in protobuffer")
	rootCmd.Flags().StringP("prune", "P", "", "set prune `RULES` using same format as listed by -p")
	rootCmd.Flags().StringP("regex", "e", "", "only print lines matching the given ECMAScript `EXPR`")
	rootCmd.Flags().IntP("rotate-count", "n", 0, "set max number of rotated logs (default 4)")
	rootCmd.Flags().IntP("rotate-kbytes", "r", 0, "rotate log every `N` KiB (requires -f)")
	rootCmd.Flags().BoolP("silent", "s", false, "set default filter to silent (like filterspec '*:S')")
	rootCmd.Flags().StringP("since", "T", "", "print most recent `COUNT` lines or since `TIME` (does not imply -d)")
	rootCmd.Flags().BoolP("statistics", "S", false, "output statistics")
	rootCmd.Flags().StringP("tail", "t", "", "print most recent `COUNT` lines or since `TIME` (implies -d)")
	rootCmd.Flags().String("uid", "", "only display log messages from `UIDS` in the comma-separated list")
	rootCmd.Flags().Bool("wrap", false, "sleep for 2 hours or until buffer is about to wrap")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"buffer": carapace.ActionValuesDescribed(
			"main", "main log buffer",
			"system", "system log buffer",
			"radio", "radio log buffer",
			"events", "event log buffer",
			"crash", "crash log buffer",
			"default", "default log buffer",
			"all", "all log buffers",
		).UniqueList(","),
		"buffer-size": carapace.ActionValues(),
		"file":        carapace.ActionFiles(),
		"format": carapace.ActionValuesDescribed(
			"brief", "show priority, tag, and PID",
			"long", "show all metadata fields",
			"process", "show PID only",
			"raw", "show the raw log message",
			"tag", "show the priority and tag only",
			"thread", "show priority, PID, and TID",
			"threadtime", "show date, time, priority, tag, PID, and TID (default)",
			"time", "show date, time, priority, tag, and PID",
			"color", "adverb: show each priority with a different color",
			"descriptive", "adverb: show event descriptions",
			"epoch", "adverb: show time as seconds since 1970-01-01",
			"monotonic", "adverb: show time as CPU seconds since boot",
			"printable", "adverb: ensure that any binary logging content is escaped",
			"uid", "adverb: show UID of logged process",
			"usec", "adverb: show time with microsecond precision",
			"UTC", "adverb: show time as UTC",
			"year", "adverb: add the year to the displayed time",
			"zone", "adverb: add the local timezone to the displayed time",
		).UniqueList(","),
		"regex": carapace.ActionValues(),
		"uid":   carapace.ActionValues(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("*").Suffix(":")
			default:
				return carapace.ActionValuesDescribed(
					"V", "Verbose",
					"D", "Debug",
					"I", "Info",
					"W", "Warn",
					"E", "Error",
					"F", "Fatal",
					"S", "Silent (suppress all output)",
				)
			}
		}),
	)
}
