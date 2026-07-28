package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Show Tailscale IP addresses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ipCmd).Standalone()

	ipCmd.Flags().Bool("1", false, "only print one IP address")
	ipCmd.Flags().Bool("4", false, "only print IPv4 address")
	ipCmd.Flags().Bool("6", false, "only print IPv6 address")
	ipCmd.Flags().String("assert", "", "assert that one of the node's IPs matches this IP address")
	rootCmd.AddCommand(ipCmd)
}
