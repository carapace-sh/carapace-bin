package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scriptlive",
	Short: "re-run session typescripts, using timing information",
	Long:  "https://man7.org/linux/man-pages/man1/scriptlive.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("command", "c", "", "run command rather than interactive shell")
	rootCmd.Flags().StringP("divisor", "d", "", "speed up or slow down execution with time divisor")
	rootCmd.Flags().StringP("echo", "E", "", "echo input in session (auto, always or never)")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().StringP("log-in", "I", "", "script stdin log file")
	rootCmd.Flags().StringP("log-io", "B", "", "script stdin and stdout log file")
	rootCmd.Flags().StringP("log-timing", "T", "", "alias to -t")
	rootCmd.Flags().StringP("maxdelay", "m", "", "wait at most this many seconds between updates")
	rootCmd.Flags().StringP("timing", "t", "", "script timing log file")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"echo":       carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"log-in":     carapace.ActionFiles(),
		"log-io":     carapace.ActionFiles(),
		"log-timing": carapace.ActionFiles(),
		"timing":     carapace.ActionFiles(),
	})
}
