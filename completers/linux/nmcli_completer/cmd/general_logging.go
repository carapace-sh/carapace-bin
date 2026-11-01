package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var general_loggingCmd = &cobra.Command{
	Use:   "logging",
	Short: "Get or change NetworkManager logging level and domains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(general_loggingCmd).Standalone()

	generalCmd.AddCommand(general_loggingCmd)

	carapace.Gen(general_loggingCmd).PositionalCompletion(
		carapace.ActionValues("level"),
		carapace.ActionValuesDescribed(
			"DEBUG", "verbose logging for debugging purposes",
			"ERR", "log only critical errors",
			"INFO", "log various informational messages",
			"OFF", "logging disabled",
			"TRACE", "even more verbose logging than DEBUG",
			"WARN", "log warning messages",
		),
		carapace.ActionValues("domains"),
		ActionDomains().UniqueList(","),
	)
}

func ActionDomains() carapace.Action {
	return carapace.ActionValues(
		"ADSL", "AGENTS", "ALL", "AUDIT", "AUTOIP4", "BOND", "BRIDGE",
		"BT", "CONCHECK", "CORE", "DBUS_PROPS", "DCB", "DEFAULT", "DEVICE", "DHCP",
		"DHCP4", "DHCP6", "DISPATCH", "DNS", "ETHER", "FIREWALL",
		"HW", "INFINIBAND", "IP", "IP4", "IP6", "MB",
		"NONE", "OLPC", "PLATFORM", "PPP", "PROXY", "RFKILL",
		"SETTINGS", "SHARING", "SUPPLICANT", "SUSPEND", "SYSTEMD", "TEAM",
		"VLAN", "VPN",
		"VPN_PLUGIN", "WIFI", "WIFI_SCAN", "WIMAX",
	)
}
