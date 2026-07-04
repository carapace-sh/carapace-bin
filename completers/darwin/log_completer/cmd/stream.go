package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Watch live system logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(streamCmd).Standalone()

	streamCmd.Flags().Bool("backtrace", false, "Control whether backtraces are shown")
	streamCmd.Flags().String("color", "", "Control color output (auto, always, none)")
	streamCmd.Flags().Bool("ignore-dropped", false, "Don't output anything for dropped messages")
	streamCmd.Flags().String("level", "", "Include events at, and below, the given level")
	streamCmd.Flags().Bool("mach-continuous-time", false, "Print mach continuous time timestamps")
	streamCmd.Flags().String("predicate", "", "Filter events using the given predicate")
	streamCmd.Flags().String("process", "", "Stream events from the specified process")
	streamCmd.Flags().Bool("source", false, "Annotate output with source file and line-number")
	streamCmd.Flags().String("style", "", "Output format (default, syslog, json, ndjson, compact)")
	streamCmd.Flags().String("timeout", "", "Terminate streaming after timeout has elapsed")
	streamCmd.Flags().String("type", "", "Limit streaming to a given event type")
	streamCmd.Flags().Bool("unreliable", false, "Annotate output with whether the log was emitted unreliably")
	streamCmd.Flags().String("user", "", "Stream events from the specified user")
	rootCmd.AddCommand(streamCmd)

	carapace.Gen(streamCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("auto", "always", "none"),
		"level": carapace.ActionValues("default", "info", "debug"),
		"style": carapace.ActionValues("default", "syslog", "json", "ndjson", "compact"),
		"type":  carapace.ActionValues("activity", "log", "trace"),
	})
}
