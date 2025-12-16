package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "w",
	Short: "Show who is logged on and what they are doing",
	Long:  "https://man7.org/linux/man-pages/man1/w.1.html",
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

	rootCmd.Flags().BoolP("container", "c", false, "show container uptime")
	rootCmd.Flags().BoolP("from", "f", false, "show remote hostname field")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("ip-addr", "i", false, "display IP address instead of hostname (if possible)")
	rootCmd.Flags().BoolP("no-current", "u", false, "ignore current process username")
	rootCmd.Flags().BoolP("no-header", "h", false, "do not print header")
	rootCmd.Flags().BoolP("old-style", "o", false, "old style output")
	rootCmd.Flags().BoolP("pids", "p", false, "show the PID(s) of processes in WHAT")
	rootCmd.Flags().BoolP("short", "s", false, "short format")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")

	carapace.Gen(rootCmd).PositionalCompletion(
		os.ActionUsers(),
	)
}
