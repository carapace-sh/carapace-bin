package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "networksetup",
	Short: "configuration tool for network settings in the System Preferences",
	Long:  "https://keith.github.io/xcode-manpages/networksetup.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"-listallnetworkservices", "List all network services",
			"-listallhardwareports", "List all hardware ports",
			"-listnetworkserviceorder", "List network services in order",
			"-getnetworkserviceenabled", "Get the enabled status of a network service",
			"-setnetworkserviceenabled", "Enable or disable a network service",
			"-getdnsservers", "Get DNS servers for a network service",
			"-setdnsservers", "Set DNS servers for a network service",
			"-getinfo", "Get information about a network service",
			"-setmanual", "Set manual IP configuration",
			"-setdhcp", "Set DHCP configuration",
			"-setbootp", "Set BOOTP configuration",
			"-setmanualwithdhcprouter", "Set manual IP with DHCP router",
			"-getproxy", "Get proxy settings for a network service",
			"-setwebproxy", "Set web proxy for a network service",
			"-setsecurewebproxy", "Set secure web proxy for a network service",
			"-setftpproxy", "Set FTP proxy for a network service",
			"-setsocksfirewallproxy", "Set SOCKS firewall proxy for a network service",
			"-setproxybypassdomains", "Set proxy bypass domains for a network service",
			"-setv6off", "Disable IPv6 for a network service",
			"-setv6automatic", "Set IPv6 to automatic for a network service",
			"-setv6manual", "Set IPv6 manual configuration for a network service",
			"-listWirelessPorts", "List wireless ports",
			"-setairportnetwork", "Set the AirPort network",
			"-getairportnetwork", "Get the current AirPort network",
		),
	)
}
