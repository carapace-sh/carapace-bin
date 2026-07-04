package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "afplay",
	Short: "play audio files",
	Long:  "https://keith.github.io/xcode-manpages/afplay.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("debug", "d", false, "Debug print output")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().Bool("leaks", false, "Run leaks analysis")
	rootCmd.Flags().StringP("quality", "q", "", "Set quality for rate-scaled playback (0=low, 1=high)")
	rootCmd.Flags().StringP("rate", "r", "", "Play at playback rate")
	rootCmd.Flags().StringP("time", "t", "", "Play for TIME seconds")
	rootCmd.Flags().StringP("volume", "v", "", "Set the volume for playback of the file")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
