package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uptime",
	Short: "Tell how long the system has been running",
	Long:  "https://man7.org/linux/man-pages/man1/uptime.1.html",
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
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("pretty", "p", false, "show uptime in pretty format")
	rootCmd.Flags().BoolP("raw", "r", false, "show uptime values in raw format")
	rootCmd.Flags().BoolP("since", "s", false, "system up since")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")
}
