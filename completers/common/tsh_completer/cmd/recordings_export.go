package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var recordings_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export recorded desktop sessions to video.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recordings_exportCmd).Standalone()

	recordings_exportCmd.Flags().String("encoder", "auto", "The video encoder to use.")
	recordings_exportCmd.Flags().Bool("no-quiet", false, "Quiet mode.")
	recordings_exportCmd.Flags().String("out", "", "Override output file name.")
	recordings_exportCmd.Flags().BoolP("quiet", "q", false, "Quiet mode.")
	recordings_exportCmd.Flag("no-quiet").Hidden = true
	recordingsCmd.AddCommand(recordings_exportCmd)
}
