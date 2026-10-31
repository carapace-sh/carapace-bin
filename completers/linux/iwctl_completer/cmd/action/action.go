package action

import (
	"path/filepath"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
)

// Command represents a command of an iwctl command family
type Command struct {
	Name        string
	Description string
	Entity      bool
}

// Commands maps an iwctl command family to its commands
var Commands = map[string][]Command{
	"ad-hoc": {
		{Name: "list", Description: "List devices in Ad-hoc mode"},
		{Name: "start", Description: "Start or join an existing Ad-Hoc network called \"network name\" with a passphrase", Entity: true},
		{Name: "start_open", Description: "Start or join an existing open Ad-Hoc network called \"network name\"", Entity: true},
		{Name: "stop", Description: "Leave an Ad-Hoc network", Entity: true},
	},
	"adapter": {
		{Name: "list", Description: "List adapters"},
		{Name: "show", Description: "Show adapter info", Entity: true},
		{Name: "set-property", Description: "Set property", Entity: true},
	},
	"ap": {
		{Name: "list", Description: "List devices in AP mode"},
		{Name: "start", Description: "Start an access point called \"network name\" with a passphrase", Entity: true},
		{Name: "start-profile", Description: "Start an access point based on a disk profile", Entity: true},
		{Name: "stop", Description: "Stop a started access point", Entity: true},
		{Name: "show", Description: "Show AP info", Entity: true},
		{Name: "scan", Description: "Start an AP scan", Entity: true},
		{Name: "get-networks", Description: "Get network list after scanning", Entity: true},
	},
	"debug": {
		{Name: "connect", Description: "Connect to a specific BSS", Entity: true},
		{Name: "roam", Description: "Roam to a BSS", Entity: true},
		{Name: "get-networks", Description: "Get networks", Entity: true},
		{Name: "autoconnect", Description: "Set AutoConnect property", Entity: true},
	},
	"device": {
		{Name: "list", Description: "List devices"},
		{Name: "show", Description: "Show device info", Entity: true},
		{Name: "set-property", Description: "Set property", Entity: true},
	},
	"dpp": {
		{Name: "list", Description: "List DPP-capable devices"},
		{Name: "start-enrollee", Description: "Starts a DPP Enrollee", Entity: true},
		{Name: "start-configurator", Description: "Starts a DPP Configurator", Entity: true},
		{Name: "stop", Description: "Aborts DPP operations", Entity: true},
		{Name: "show", Description: "Shows the DPP state", Entity: true},
	},
	"known-networks": {
		{Name: "list", Description: "List known networks"},
		{Name: "forget", Description: "Forget known network", Entity: true},
		{Name: "show", Description: "Show known network", Entity: true},
		{Name: "set-property", Description: "Set property", Entity: true},
	},
	"pkex": {
		{Name: "list", Description: "List shared code capable devices"},
		{Name: "stop", Description: "Aborts shared code operations", Entity: true},
		{Name: "show", Description: "Shows the shared code state", Entity: true},
		{Name: "enroll", Description: "Start a shared code enrollee", Entity: true},
		{Name: "configure", Description: "Start a shared code configurator", Entity: true},
	},
	"station": {
		{Name: "list", Description: "List devices in Station mode"},
		{Name: "connect", Description: "Connect to network", Entity: true},
		{Name: "connect-hidden", Description: "Connect to hidden network", Entity: true},
		{Name: "disconnect", Description: "Disconnect", Entity: true},
		{Name: "get-networks", Description: "Get networks", Entity: true},
		{Name: "get-hidden-access-points", Description: "Get hidden APs", Entity: true},
		{Name: "scan", Description: "Scan for networks", Entity: true},
		{Name: "show", Description: "Show station info", Entity: true},
		{Name: "get-bsses", Description: "Get BSS's for a network", Entity: true},
	},
	"wsc": {
		{Name: "list", Description: "List WSC-capable devices"},
		{Name: "push-button", Description: "PushButton mode", Entity: true},
		{Name: "start-user-pin", Description: "PIN mode", Entity: true},
		{Name: "start-pin", Description: "PIN mode with generated 8 digit PIN", Entity: true},
		{Name: "cancel", Description: "Aborts WSC operations", Entity: true},
	},
}

// FamilyDescriptions describes the iwctl command families
var FamilyDescriptions = map[string]string{
	"ad-hoc":         "Ad-Hoc",
	"adapter":        "Adapters",
	"ap":             "Access Point",
	"debug":          "Station Debug",
	"device":         "Devices",
	"dpp":            "Device Provisioning",
	"known-networks": "Known Networks",
	"pkex":           "Shared Code Device Provisioning (PKEX)",
	"station":        "Station",
	"wsc":            "WiFi Simple Configuration",
}

// ActionDevices completes wifi devices (iwd entity <wlan>)
func ActionDevices() carapace.Action {
	return net.ActionDevices(net.IncludedDevices{Wifi: true})
}

// ActionAdapters completes adapters (iwd entity <phy>)
func ActionAdapters() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		matches, _ := filepath.Glob("/sys/class/ieee80211/phy*")
		vals := make([]string, 0, len(matches))
		for _, match := range matches {
			vals = append(vals, filepath.Base(match))
		}
		return carapace.ActionValues(vals...)
	}).Tag("adapters")
}
