package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "traceroute",
	Short: "Trace route to a network host",
	Long:  "https://keith.github.io/xcode-manpages/traceroute.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "Perform AS lookups")
	rootCmd.Flags().BoolS("I", "I", false, "Use ICMP ECHO probes")
	rootCmd.Flags().BoolS("N", "N", false, "Do not probe path MTU")
	rootCmd.Flags().BoolS("S", "S", false, "Use TCP SYN probes")
	rootCmd.Flags().BoolS("a", "a", false, "Perform AS lookups")
	rootCmd.Flags().BoolS("d", "d", false, "Enable socket-level debugging")
	rootCmd.Flags().BoolS("e", "e", false, "Display extended info")
	rootCmd.Flags().BoolS("f", "f", false, "Show firewalls")
	rootCmd.Flags().BoolS("n", "n", false, "Do not resolve host names")
	rootCmd.Flags().BoolS("r", "r", false, "Do not use routing table")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose mode")
	rootCmd.Flags().BoolS("x", "x", false, "Do not compute checksums")

	rootCmd.Flags().StringS("g", "g", "", "Source route gateway")
	rootCmd.Flags().StringS("i", "i", "", "Specify interface")
	rootCmd.Flags().StringS("m", "m", "", "Maximum TTL")
	rootCmd.Flags().StringS("p", "p", "", "Destination port")
	rootCmd.Flags().StringS("q", "q", "", "Number of probes per TTL")
	rootCmd.Flags().StringS("s", "s", "", "Source address")
	rootCmd.Flags().StringS("t", "t", "", "Type of service")
	rootCmd.Flags().StringS("w", "w", "", "Wait time for response")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionValues())
}
