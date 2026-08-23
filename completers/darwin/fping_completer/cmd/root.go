package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fping",
	Short: "send ICMP ECHO_REQUEST packets to network hosts",
	Long:  "https://man.freebsd.org/cgi/man.cgi?fping",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("4", "4", false, "Use IPv4")
	rootCmd.Flags().BoolS("6", "6", false, "Use IPv6")
	rootCmd.Flags().BoolS("A", "A", false, "Display by address")
	rootCmd.Flags().StringS("B", "B", "", "Backoff factor")
	rootCmd.Flags().StringS("C", "C", "", "Vcount")
	rootCmd.Flags().BoolS("D", "D", false, "Timestamp")
	rootCmd.Flags().StringS("H", "H", "", "TTL")
	rootCmd.Flags().StringS("I", "I", "", "Interface")
	rootCmd.Flags().BoolS("M", "M", false, "Don't fragment")
	rootCmd.Flags().BoolS("N", "N", false, "Netdata")
	rootCmd.Flags().StringS("O", "O", "", "TOS")
	rootCmd.Flags().StringS("Q", "Q", "", "Squiet (seconds)")
	rootCmd.Flags().BoolS("R", "R", false, "Random")
	rootCmd.Flags().StringS("S", "S", "", "Source address")
	rootCmd.Flags().StringS("X", "X", "", "Fast reachable")
	rootCmd.Flags().BoolS("a", "a", false, "Show alive systems")
	rootCmd.Flags().StringS("b", "b", "", "Size of ping data")
	rootCmd.Flags().StringS("c", "c", "", "Count")
	rootCmd.Flags().BoolS("d", "d", false, "Use rDNS")
	rootCmd.Flags().BoolS("e", "e", false, "Elapsed time")
	rootCmd.Flags().StringS("f", "f", "", "File")
	rootCmd.Flags().StringS("g", "g", "", "Generate addr/mask")
	rootCmd.Flags().BoolS("h", "h", false, "Help")
	rootCmd.Flags().StringS("i", "i", "", "Interval (ms)")
	rootCmd.Flags().BoolS("l", "l", false, "Loop")
	rootCmd.Flags().BoolS("m", "m", false, "All targets")
	rootCmd.Flags().BoolS("n", "n", false, "Show name")
	rootCmd.Flags().BoolS("o", "o", false, "Outage")
	rootCmd.Flags().StringS("p", "p", "", "Period (ms)")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet")
	rootCmd.Flags().StringS("r", "r", "", "Retry")
	rootCmd.Flags().BoolS("s", "s", false, "Stats")
	rootCmd.Flags().StringS("t", "t", "", "Timeout (ms)")
	rootCmd.Flags().BoolS("u", "u", false, "Unreachable")
	rootCmd.Flags().BoolS("v", "v", false, "Version")
	rootCmd.Flags().StringS("x", "x", "", "Reachable")
}
