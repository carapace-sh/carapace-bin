package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "netsh",
	Short: "network shell for configuring network components",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/netsh",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"add", "add a configuration entry",
			"advfirewall", "advanced firewall configuration",
			"bridge", "network bridge configuration",
			"delete", "delete a configuration entry",
			"dhcpclient", "DHCP client configuration",
			"dnsclient", "DNS client configuration",
			"dump", "display a configuration script",
			"exec", "run a script file",
			"firewall", "firewall configuration (legacy)",
			"http", "HTTP configuration",
			"interface", "interface configuration",
			"ipsec", "IPsec configuration",
			"lan", "wired LAN configuration",
			"namespace", "namespace configuration",
			"nap", "Network Access Protection configuration",
			"netio", "Network I/O configuration",
			"ras", "Remote Access configuration",
			"rpc", "RPC configuration",
			"set", "set configuration information",
			"show", "display information",
			"trace", "enable/disable tracing",
			"wcn", "Windows Connect Now configuration",
			"wfp", "Windows Filtering Platform configuration",
			"winhttp", "WinHTTP proxy configuration",
			"winsock", "Winsock configuration",
			"wlan", "wireless LAN configuration",
		),
	)
}
