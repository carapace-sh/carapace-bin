package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "say",
	Short: "convert text to audible speech",
	Long:  "https://keith.github.io/xcode-manpages/say.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("format", "n", "", "Set the output audio file format")
	rootCmd.Flags().StringP("input", "f", "", "Convert the text file content to speech")
	rootCmd.Flags().StringP("network", "N", "", "Use the specified network service to download voices")
	rootCmd.Flags().StringP("output", "o", "", "Specify the path for an audio file to be written")
	rootCmd.Flags().StringP("progress", "p", "", "Show progress indicator while speaking")
	rootCmd.Flags().StringP("quality", "quality", "", "Set the audio quality")
	rootCmd.Flags().StringP("rate", "r", "", "Set the speech rate in words per minute")
	rootCmd.Flags().StringP("voice", "v", "", "Specify the voice to be used")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"input":  carapace.ActionFiles(),
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
	)
}
