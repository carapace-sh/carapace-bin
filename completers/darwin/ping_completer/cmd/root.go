package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ping",
	Short: "Send ICMP ECHO_REQUEST packets",
	Long:  "https://keith.github.io/xcode-manpages/ping.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolS("Adaptive|-A", "Adaptive|-A", false, "")
	rootCmd.Flags().BoolS("Audible|-a", "Audible|-a", false, "")
	rootCmd.Flags().StringS("Count|-c", "Count|-c", "", "")
	rootCmd.Flags().BoolS("Flood|-f", "Flood|-f", false, "")
	rootCmd.Flags().StringS("Iface|-I", "Iface|-I", "", "")
	rootCmd.Flags().StringS("Interval|-i", "Interval|-i", "", "")
	rootCmd.Flags().BoolS("NoDNS|-n", "NoDNS|-n", false, "")
	rootCmd.Flags().StringS("Packetsize|-s", "Packetsize|-s", "", "")
	rootCmd.Flags().BoolS("Quiet|-q", "Quiet|-q", false, "")
	rootCmd.Flags().BoolS("Route|-R", "Route|-R", false, "")
	rootCmd.Flags().StringS("Src|-S", "Src|-S", "", "")
	rootCmd.Flags().StringS("TTL|-t", "TTL|-t", "", "")
	rootCmd.Flags().StringS("Timeout|-W", "Timeout|-W", "", "")
	rootCmd.Flags().BoolS("Verbose|-v", "Verbose|-v", false, "")
	rootCmd.Flags().StringS("Waittime|-w", "Waittime|-w", "", "")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
