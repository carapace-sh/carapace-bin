package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/saw"
	"github.com/spf13/cobra"
)

var streamsCmd = &cobra.Command{
	Use:   "streams <log group>",
	Short: "List streams in log group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(streamsCmd).Standalone()

	streamsCmd.Flags().Bool("descending", false, "order streams descending")
	streamsCmd.Flags().String("orderBy", "", "order streams by LogStreamName or LastEventTime")
	streamsCmd.Flags().String("prefix", "", "stream prefix filter")
	rootCmd.AddCommand(streamsCmd)

	carapace.Gen(streamsCmd).FlagCompletion(carapace.ActionMap{
		"orderBy": carapace.ActionValues("LogStreamName", "LastEventTime"),
	})

	carapace.Gen(streamsCmd).PositionalCompletion(
		saw.ActionGroups(),
	)
}
