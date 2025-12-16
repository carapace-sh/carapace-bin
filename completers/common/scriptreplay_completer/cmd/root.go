package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scriptreplay",
	Short: "play back typescripts, using timing information",
	Long:  "https://man7.org/linux/man-pages/man1/scriptreplay.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("cr-mode", "c", "", "CR char mode")
	rootCmd.Flags().StringP("divisor", "d", "", "speed up or slow down execution with time divisor")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().StringP("log-in", "I", "", "script stdin log file")
	rootCmd.Flags().StringP("log-io", "B", "", "script stdin and stdout log file")
	rootCmd.Flags().StringP("log-out", "O", "", "script stdout log file (default)")
	rootCmd.Flags().StringP("log-timing", "T", "", "alias to -t")
	rootCmd.Flags().StringP("maxdelay", "m", "", "wait at most this many seconds between updates")
	rootCmd.Flags().StringP("stream", "x", "", "stream type")
	rootCmd.Flags().Bool("summary", false, "display overview about recorded session and exit")
	rootCmd.Flags().StringP("timing", "t", "", "script timing log file")
	rootCmd.Flags().StringP("typescript", "s", "", "deprecated alias to -O")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cr-mode":    carapace.ActionValues("auto", "always", "never"),
		"log-in":     carapace.ActionFiles(),
		"log-io":     carapace.ActionFiles(),
		"log-out":    carapace.ActionFiles(),
		"log-timing": carapace.ActionFiles(),
		"stream":     carapace.ActionValues("out", "in", "signal", "info"),
		"timing":     carapace.ActionFiles(),
		"typescript": carapace.ActionFiles(),
	})
}
