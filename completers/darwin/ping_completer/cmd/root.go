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

	rootCmd.Flags().BoolS("A", "A", false, "Adaptive ping (interval adapts to RTT)")
	rootCmd.Flags().BoolS("D", "D", false, "Print timestamp before each line")
	rootCmd.Flags().BoolS("R", "R", false, "Record route")
	rootCmd.Flags().BoolS("a", "a", false, "Audible ping (bell on each reply)")
	rootCmd.Flags().BoolS("f", "f", false, "Flood ping")
	rootCmd.Flags().BoolS("n", "n", false, "Do not resolve host names")
	rootCmd.Flags().BoolS("o", "o", false, "Exit after first reply")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet mode")
	rootCmd.Flags().BoolS("r", "r", false, "Bypass routing table")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose mode")

	rootCmd.Flags().StringS("I", "I", "", "Specify interface")
	rootCmd.Flags().StringS("S", "S", "", "Source address")
	rootCmd.Flags().StringS("W", "W", "", "Wait time for reply (milliseconds)")
	rootCmd.Flags().StringS("b", "b", "", "Bind interface")
	rootCmd.Flags().StringS("c", "c", "", "Number of packets to send")
	rootCmd.Flags().StringS("i", "i", "", "Wait interval between packets")
	rootCmd.Flags().StringS("m", "m", "", "TTL (time to live)")
	rootCmd.Flags().StringS("p", "p", "", "Pattern for packet data")
	rootCmd.Flags().StringS("s", "s", "", "Packet size")
	rootCmd.Flags().StringS("t", "t", "", "Timeout")
	rootCmd.Flags().StringS("w", "w", "", "Deadline timeout")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionValues())
}
