package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a recording to an asciinema server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uploadCmd).Standalone()

	uploadCmd.Flags().String("audio-url", "", "Audio URL for synchronized playback")
	uploadCmd.Flags().String("description", "", "Description of the recording (Markdown supported)")
	uploadCmd.Flags().BoolP("help", "h", false, "Print help")
	uploadCmd.Flags().String("server-url", "", "asciinema server URL")
	uploadCmd.Flags().StringP("title", "t", "", "Title of the recording")
	uploadCmd.Flags().String("visibility", "", "Recording visibility level")
	rootCmd.AddCommand(uploadCmd)

	carapace.Gen(uploadCmd).FlagCompletion(carapace.ActionMap{
		"visibility": carapace.ActionValues("private", "public", "unlisted"),
	})

	carapace.Gen(uploadCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
