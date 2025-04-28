package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "timeout",
	Short: "run a command with a time limit",
	Long:  "https://man7.org/linux/man-pages/man1/timeout.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().BoolP("foreground", "f", false, "when not running timeout directly from a shell prompt, allow COMMAND to read from the TTY and get TTY signals")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().StringP("kill-after", "k", "", "also send a KILL signal if COMMAND is still running this long after the initial signal was sent")
	rootCmd.Flags().BoolP("preserve-status", "p", false, "exit with the same status as COMMAND, even when the command times out")
	rootCmd.Flags().StringP("signal", "s", "", "specify the signal to be sent on timeout")
	rootCmd.Flags().BoolP("verbose", "v", false, "diagnose to stderr any signal sent upon timeout")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"signal": ps.ActionKillSignals(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin().Shift(1),
	)
}
