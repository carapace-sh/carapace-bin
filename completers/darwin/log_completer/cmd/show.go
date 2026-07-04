package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "View/search system logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().Bool("backtrace", true, "Control whether backtraces are shown")
	showCmd.Flags().String("color", "", "Control color output (auto, always, none)")
	showCmd.Flags().Bool("debug", false, "Control whether Debug events are shown")
	showCmd.Flags().String("end", "", "Display events up to the given end date")
	showCmd.Flags().Bool("info", false, "Control whether Info events are shown")
	showCmd.Flags().String("last", "", "Display recent events up to the given limit")
	showCmd.Flags().Bool("loss", false, "Control whether message loss events are shown")
	showCmd.Flags().Bool("mach-continuous-time", false, "Print mach continuous time timestamps")
	showCmd.Flags().Bool("pager", true, "Paginate output using less")
	showCmd.Flags().String("predicate", "", "Filter events using the given predicate")
	showCmd.Flags().String("process", "", "Filter events using the specified process")
	showCmd.Flags().Bool("signpost", false, "Control whether signposts are shown")
	showCmd.Flags().String("start", "", "Display events from the given start date")
	showCmd.Flags().String("style", "", "Output format (default, syslog, json, ndjson, compact)")
	showCmd.Flags().String("timezone", "", "Use the given timezone when displaying event timestamps")
	showCmd.Flags().Bool("unreliable", false, "Annotate output with whether the log was emitted unreliably")
	rootCmd.AddCommand(showCmd)

	carapace.Gen(showCmd).FlagCompletion(carapace.ActionMap{
		"color":    carapace.ActionValues("auto", "always", "none"),
		"last":     carapace.ActionValues(),
		"style":    carapace.ActionValues("default", "syslog", "json", "ndjson", "compact"),
		"timezone": carapace.ActionValues("local"),
	})

	carapace.Gen(showCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
