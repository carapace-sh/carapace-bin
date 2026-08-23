package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "traceroute6",
	Short: "print the route IPv6 packets will take to a network node",
	Long:  "https://keith.github.io/xcode-manpages/traceroute6.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("I", "I", false, "Use ICMP ECHO instead of UDP")
	rootCmd.Flags().BoolS("N", "N", false, "Numeric output")
	rootCmd.Flags().BoolS("T", "T", false, "Use TCP SYN")
	rootCmd.Flags().BoolS("U", "U", false, "Use UDP datagrams")
	rootCmd.Flags().BoolS("a", "a", false, "Resolve addresses")
	rootCmd.Flags().BoolS("d", "d", false, "Set SO_DEBUG")
	rootCmd.Flags().BoolS("e", "e", false, "Echo")
	rootCmd.Flags().StringS("f", "f", "", "First hop")
	rootCmd.Flags().StringS("g", "g", "", "Gateway")
	rootCmd.Flags().BoolS("l", "l", false, "Display TTL")
	rootCmd.Flags().StringS("m", "m", "", "Max hop limit")
	rootCmd.Flags().BoolS("n", "n", false, "Numeric output")
	rootCmd.Flags().StringS("p", "p", "", "Port")
	rootCmd.Flags().StringS("q", "q", "", "Number of probes")
	rootCmd.Flags().BoolS("r", "r", false, "Bypass routing tables")
	rootCmd.Flags().StringS("s", "s", "", "Source address")
	rootCmd.Flags().StringS("t", "t", "", "Traffic class")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")
	rootCmd.Flags().StringS("w", "w", "", "Wait time")
}
