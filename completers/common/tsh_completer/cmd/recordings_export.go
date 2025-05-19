package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var recordings_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export recorded desktop sesions to video.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recordings_exportCmd).Standalone()

	recordings_exportCmd.Flags().String("out", "", "Override output file name")
	recordingsCmd.AddCommand(recordings_exportCmd)
}
