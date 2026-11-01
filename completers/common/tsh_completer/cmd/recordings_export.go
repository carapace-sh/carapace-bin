package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
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

	carapace.Gen(recordings_exportCmd).FlagCompletion(carapace.ActionMap{
		"encoder": carapace.ActionValues("auto", "avi", "ffmpeg"),
		"out":     carapace.ActionFiles(),
	})

	carapace.Gen(recordings_exportCmd).PositionalCompletion(
		tsh.ActionRecordings(),
	)
}
