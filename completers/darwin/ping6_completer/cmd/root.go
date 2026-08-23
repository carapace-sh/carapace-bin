package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ping6",
	Short: "send ICMPv6 ECHO_REQUEST packets to network hosts",
	Long:  "https://keith.github.io/xcode-manpages/ping6.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("B", "B", "", "Bound interface")
	rootCmd.Flags().BoolS("C", "C", false, "Apple connectivity test")
	rootCmd.Flags().BoolS("D", "D", false, "Dont fragment")
	rootCmd.Flags().StringS("G", "G", "", "Sweep max size")
	rootCmd.Flags().BoolS("H", "H", false, "Print headers")
	rootCmd.Flags().StringS("I", "I", "", "Interface")
	rootCmd.Flags().StringS("K", "K", "", "Network service type")
	rootCmd.Flags().BoolS("L", "L", false, "Suppress loopback")
	rootCmd.Flags().BoolS("N", "N", false, "Numeric output only")
	rootCmd.Flags().StringS("P", "P", "", "Policy")
	rootCmd.Flags().StringS("S", "S", "", "Source address")
	rootCmd.Flags().BoolS("T", "T", false, "Timestamp")
	rootCmd.Flags().StringS("a", "a", "", "Addr type")
	rootCmd.Flags().StringS("b", "b", "", "Buffer size")
	rootCmd.Flags().BoolS("c", "c", false, "Stop after count replies")
	rootCmd.Flags().BoolS("d", "d", false, "Set SO_DEBUG")
	rootCmd.Flags().BoolS("f", "f", false, "Flood ping")
	rootCmd.Flags().StringS("g", "g", "", "Gateway")
	rootCmd.Flags().StringS("h", "h", "", "Hop limit")
	rootCmd.Flags().StringS("i", "i", "", "Wait interval")
	rootCmd.Flags().StringS("k", "k", "", "Traffic class")
	rootCmd.Flags().StringS("l", "l", "", "Preload")
	rootCmd.Flags().BoolS("m", "m", false, "Join multicast group")
	rootCmd.Flags().BoolS("n", "n", false, "Numeric output")
	rootCmd.Flags().BoolS("o", "o", false, "Exit successfully after one reply")
	rootCmd.Flags().StringS("p", "p", "", "Pattern")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet output")
	rootCmd.Flags().StringS("s", "s", "", "Packet size")
	rootCmd.Flags().BoolS("t", "t", false, "TTL")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose output")
	rootCmd.Flags().BoolS("w", "w", false, "Timeout")
	rootCmd.Flags().StringS("z", "z", "", "Traffic class")
}
