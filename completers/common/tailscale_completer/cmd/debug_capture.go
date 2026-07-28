package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_captureCmd = &cobra.Command{
	Use:   "capture",
	Short: "Stream pcaps for debugging",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_captureCmd).Standalone()

	debug_captureCmd.Flags().StringP("o", "o", "", "path to stream the pcap (or - for stdout); leave empty to start wireshark")
	debugCmd.AddCommand(debug_captureCmd)

	carapace.Gen(debug_captureCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionFiles(),
	})
}
