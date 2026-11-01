package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play back a terminal session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(playCmd).Standalone()

	playCmd.Flags().BoolP("help", "h", false, "Print help")
	playCmd.Flags().StringP("idle-time-limit", "i", "", "Limit idle time to a given number of seconds")
	playCmd.Flags().BoolP("loop", "l", false, "Loop playback continuously")
	playCmd.Flags().BoolP("pause-on-markers", "m", false, "Automatically pause on markers")
	playCmd.Flags().BoolP("resize", "r", false, "Auto-resize terminal to match original size")
	playCmd.Flags().StringP("speed", "s", "", "Set playback speed")
	rootCmd.AddCommand(playCmd)

	carapace.Gen(playCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			carapace.ActionValuesDescribed("-", "standard input"),
		).ToA(),
	)
}
