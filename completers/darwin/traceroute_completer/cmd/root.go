package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "traceroute",
	Short: "Trace route to host",
	Long:  "https://keith.github.io/xcode-manpages/traceroute.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolS("All|-a", "All|-a", false, "")
	rootCmd.Flags().BoolS("Debug|-d", "Debug|-d", false, "")
	rootCmd.Flags().BoolS("Extended|-e", "Extended|-e", false, "")
	rootCmd.Flags().BoolS("Firewall|-F", "Firewall|-F", false, "")
	rootCmd.Flags().StringS("Gateway|-g", "Gateway|-g", "", "")
	rootCmd.Flags().BoolS("ICMP|-I", "ICMP|-I", false, "")
	rootCmd.Flags().StringS("Iface|-i", "Iface|-i", "", "")
	rootCmd.Flags().StringS("Max ttl|-m", "Max ttl|-m", "", "")
	rootCmd.Flags().BoolS("Noprobe|-N", "Noprobe|-N", false, "")
	rootCmd.Flags().BoolS("Numeric|-n", "Numeric|-n", false, "")
	rootCmd.Flags().StringS("Port|-p", "Port|-p", "", "")
	rootCmd.Flags().StringS("Queries|-q", "Queries|-q", "", "")
	rootCmd.Flags().BoolS("Resolve|-r", "Resolve|-r", false, "")
	rootCmd.Flags().StringS("Source|-s", "Source|-s", "", "")
	rootCmd.Flags().StringS("Tos|-t", "Tos|-t", "", "")
	rootCmd.Flags().BoolS("Verbose|-v", "Verbose|-v", false, "")
	rootCmd.Flags().StringS("Waittime|-w", "Waittime|-w", "", "")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
