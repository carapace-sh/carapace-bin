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
			"OFF", "logging disabled",
			"ERR", "log only critical errors",
			"WARN", "log warning messages",
			"INFO", "log various informational messages",
			"DEBUG", "verbose logging for debugging purposes",
			"TRACE", "even more verbose logging than DEBUG",
		),
		carapace.ActionValues("domains"),
		ActionDomains().UniqueList(","),
	)
}

func ActionDomains() carapace.Action {
	return carapace.ActionValues(
		"PLATFORM", "RFKILL", "ETHER", "WIFI", "WIFI_SCAN", "BT", "MB",
		"DHCP4", "DHCP6", "DHCP", "PPP", "IP4", "IP6", "IP", "AUTOIP4",
		"DNS", "VPN", "VPN_PLUGIN", "SHARING", "SUPPLICANT", "AGENTS",
		"SETTINGS", "SUSPEND", "CORE", "DEVICE", "OLPC", "WIMAX",
		"INFINIBAND", "FIREWALL", "ADSL", "BOND", "VLAN", "BRIDGE",
		"DBUS_PROPS", "TEAM", "CONCHECK", "DCB", "DISPATCH", "AUDIT",
		"SYSTEMD", "PROXY",
		"NONE", "ALL", "DEFAULT", "HW",
	)
}
