package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pulseWorkCmd = &cobra.Command{
	Use:   "pulse:work [--stop-when-empty]",
	Short: "Process incoming Pulse data from the ingest stream",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pulseWorkCmd).Standalone()

	pulseWorkCmd.Flags().Bool("stop-when-empty", false, "Stop when the stream is empty")
	rootCmd.AddCommand(pulseWorkCmd)
}
