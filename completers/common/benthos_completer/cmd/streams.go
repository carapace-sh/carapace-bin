package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var streamsCmd = &cobra.Command{
	Use:   "streams",
	Short: "Run Benthos in streams mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(streamsCmd).Standalone()

	streamsCmd.Flags().Bool("no-api", false, "Disable the HTTP API for streams mode")
	streamsCmd.Flags().Bool("prefix-stream-endpoints", false, "Whether HTTP endpoints registered by stream configs should be prefixed with the stream ID")
	rootCmd.AddCommand(streamsCmd)

	carapace.Gen(streamsCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
