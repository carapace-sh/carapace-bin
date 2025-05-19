package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var replayCmd = &cobra.Command{
	Use:   "replay [OPTIONS] <CAST_FILE>",
	Short: "Replay an asciicast terminal session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(replayCmd).Standalone()

	replayCmd.Flags().Bool("cat", false, "Just emit raw escape sequences all at once, with no timing information")
	replayCmd.Flags().Bool("explain", false, "Explain what is being sent/received")
	replayCmd.Flags().Bool("explain-only", false, "Don't replay, just show the explanation")
	replayCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(replayCmd)

	carapace.Gen(replayCmd).PositionalCompletion(
		carapace.ActionFiles(".cast"),
	)
}
