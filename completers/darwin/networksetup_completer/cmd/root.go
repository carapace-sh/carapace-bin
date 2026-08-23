package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/darwin/networksetup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "networksetup",
	Short: "configuration tool for network settings",
	Long:  "https://keith.github.io/xcode-manpages/networksetup.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("detectnewhardware", false, "Detect new hardware ports")
	rootCmd.Flags().String("getadditionalroutes", "", "Get additional IPv4 routes for network service")
	rootCmd.Flags().String("getairportnetwork", "", "Get the current AirPort network for hardware port")
	rootCmd.Flags().String("getcomputername", "", "Get the computer name")
	rootCmd.Flags().String("getdnsservers", "", "Get DNS servers for a network service")
	rootCmd.Flags().String("getinfo", "", "Get information about a network service")
	rootCmd.Flags().String("getmacaddress", "", "Get the MAC address for a hardware port")
	rootCmd.Flags().String("getnetworkserviceenabled", "", "Get the enabled status of a network service")
	rootCmd.Flags().String("getproxybypassdomains", "", "Get proxy bypass domains for a network service")
	rootCmd.Flags().String("getsecurewebproxy", "", "Get secure web proxy for a network service")
	rootCmd.Flags().String("getsocksfirewallproxy", "", "Get SOCKS firewall proxy for a network service")
	rootCmd.Flags().String("getv6additionalroutes", "", "Get additional IPv6 routes for network service")
	rootCmd.Flags().String("getwebproxy", "", "Get web proxy for a network service")
	rootCmd.Flags().Bool("listWirelessPorts", false, "List wireless ports")
	rootCmd.Flags().Bool("listallhardwareports", false, "List all hardware ports")
	rootCmd.Flags().Bool("listallnetworkservices", false, "List all network services")
	rootCmd.Flags().Bool("listnetworkserviceorder", false, "List network services in order")
	rootCmd.Flags().String("setadditionalroutes", "", "Set additional IPv4 routes for a network service")
	rootCmd.Flags().String("setbootp", "", "Set BOOTP configuration for a network service")
	rootCmd.Flags().String("setcomputername", "", "Set the computer name")
	rootCmd.Flags().String("setdhcp", "", "Set DHCP configuration for a network service")
	rootCmd.Flags().String("setdnsservers", "", "Set DNS servers for a network service")
	rootCmd.Flags().String("setftpproxy", "", "Set FTP proxy for a network service")
	rootCmd.Flags().String("setmanual", "", "Set manual IP configuration for a network service")
	rootCmd.Flags().String("setmanualwithdhcprouter", "", "Set manual IP with DHCP router for a network service")
	rootCmd.Flags().String("setnetworkserviceenabled", "", "Enable or disable a network service (on/off)")
	rootCmd.Flags().String("setproxybypassdomains", "", "Set proxy bypass domains for a network service")
	rootCmd.Flags().String("setsecurewebproxy", "", "Set secure web proxy for a network service")
	rootCmd.Flags().String("setsocksfirewallproxy", "", "Set SOCKS firewall proxy for a network service")
	rootCmd.Flags().String("setv6automatic", "", "Set IPv6 to automatic for a network service")
	rootCmd.Flags().String("setv6linklocal", "", "Set IPv6 to link-local for a network service")
	rootCmd.Flags().String("setv6manual", "", "Set IPv6 manual configuration for a network service")
	rootCmd.Flags().String("setv6off", "", "Disable IPv6 for a network service")
	rootCmd.Flags().String("setwebproxy", "", "Set web proxy for a network service")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"getadditionalroutes":      action.ActionNetworkServices(),
		"getairportnetwork":        action.ActionNetworkServices(),
		"getdnsservers":            action.ActionNetworkServices(),
		"getinfo":                  action.ActionNetworkServices(),
		"getmacaddress":            action.ActionNetworkServices(),
		"getnetworkserviceenabled": action.ActionNetworkServices(),
		"getproxybypassdomains":    action.ActionNetworkServices(),
		"getsecurewebproxy":        action.ActionNetworkServices(),
		"getsocksfirewallproxy":    action.ActionNetworkServices(),
		"getv6additionalroutes":    action.ActionNetworkServices(),
		"getwebproxy":              action.ActionNetworkServices(),
		"setadditionalroutes":      action.ActionNetworkServices(),
		"setbootp":                 action.ActionNetworkServices(),
		"setdhcp":                  action.ActionNetworkServices(),
		"setdnsservers":            action.ActionNetworkServices(),
		"setftpproxy":              action.ActionNetworkServices(),
		"setmanual":                action.ActionNetworkServices(),
		"setmanualwithdhcprouter":  action.ActionNetworkServices(),
		"setnetworkserviceenabled": carapace.ActionValues("on", "off"),
		"setproxybypassdomains":    action.ActionNetworkServices(),
		"setsecurewebproxy":        action.ActionNetworkServices(),
		"setsocksfirewallproxy":    action.ActionNetworkServices(),
		"setv6automatic":           action.ActionNetworkServices(),
		"setv6linklocal":           action.ActionNetworkServices(),
		"setv6manual":              action.ActionNetworkServices(),
		"setv6off":                 action.ActionNetworkServices(),
		"setwebproxy":              action.ActionNetworkServices(),
	})
}
