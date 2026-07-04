package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show system logging statistics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statsCmd).Standalone()

	statsCmd.Flags().String("archive", "", "Display statistics from the given log archive")
	statsCmd.Flags().String("count", "", "Limit output to <count> rows per section")
	statsCmd.Flags().String("end", "", "Display statistics for events up to the given end date")
	statsCmd.Flags().Bool("overview", false, "Display statistics for the entire log archive")
	statsCmd.Flags().Bool("pager", true, "Paginate output using less")
	statsCmd.Flags().Bool("per-book", false, "Display statistics per log book in the archive")
	statsCmd.Flags().Bool("per-file", false, "Display statistics per file in the archive")
	statsCmd.Flags().String("predicate", "", "Display statistics for events matching the given predicate")
	statsCmd.Flags().String("process", "", "Display statistics for a given process")
	statsCmd.Flags().String("sender", "", "Display statistics for a given sender")
	statsCmd.Flags().String("sort", "", "Sort output by events or bytes")
	statsCmd.Flags().String("start", "", "Display statistics for events from the given start date")
	statsCmd.Flags().String("style", "", "Control the output style (human, json)")
	rootCmd.AddCommand(statsCmd)

	carapace.Gen(statsCmd).FlagCompletion(carapace.ActionMap{
		"archive": carapace.ActionFiles(),
		"sort":    carapace.ActionValues("events", "bytes"),
		"style":   carapace.ActionValues("human", "json"),
	})
}
