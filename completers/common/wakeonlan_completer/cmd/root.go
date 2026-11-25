package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wakeonlan",
	Short: "Send magic packets to wake-on-lan enabled computers",
	Long:  "https://github.com/jpoliv/wakeonlan",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("delay", "d", "", "delay between packets in seconds")
	rootCmd.Flags().BoolP("dry-run", "n", false, "print commands without executing them")
	rootCmd.Flags().StringP("file", "f", "", "read hardware addresses from file")
	rootCmd.Flags().BoolP("help", "h", false, "display help")
	rootCmd.Flags().StringP("ip", "i", "", "destination IP address")
	rootCmd.Flags().StringP("port", "p", "", "destination port")
	rootCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	rootCmd.Flags().BoolP("version", "v", false, "display version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
		"ip":   net.ActionHosts(),
		"port": net.ActionPorts(),
	})
}
