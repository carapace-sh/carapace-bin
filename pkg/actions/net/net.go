// package net contains network related actions
package net

import (
	"os"
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/internal/actions/net/ssh"
	"github.com/carapace-sh/carapace/pkg/execlog"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionHosts completes known hosts
//
//	192.168.1.1
//	pihole
func ActionHosts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		batch := carapace.Batch()
		if file, err := c.Abs("~/.ssh/known_hosts"); err == nil {
			if content, err := os.ReadFile(file); err == nil {
				r := regexp.MustCompile(`^(?P<host>[^ ,#|]+)`)
				rIPv4 := regexp.MustCompile(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)
				rIPv6 := regexp.MustCompile(`^([A-Fa-f0-9]{0,4}:){5,7}[A-Fa-f0-9]{0,4}$`) // TODO likely wrong
				for _, entry := range strings.Split(string(content), "\n") {
					if r.MatchString(entry) {
						if host := r.FindStringSubmatch(entry)[0]; rIPv4.MatchString(host) {
							batch = append(batch, carapace.ActionStyledValues(host, style.Default).Tag("ipv4 addresses"))
						} else if rIPv6.MatchString(host) {
							batch = append(batch, carapace.ActionStyledValues(host, style.Bold).Tag("ipv6 addresses"))
						} else {
							batch = append(batch, carapace.ActionStyledValues(host, style.Blue).Tag("hostnames"))
						}
					}
				}
			} else {
				batch = append(batch, carapace.ActionMessage(err.Error()))
			}
		}
		batch = append(batch, ssh.ActionHosts().Style(style.Yellow).Suppress(`open .*/.ssh/config: no such file or directory`))
		return batch.ToA()
	}).Tag("hosts")
}

type IncludedDevices struct {
	Bluetooth   bool
	Bond        bool
	BondSlave   bool
	Bridge      bool
	BridgeSlave bool
	Cdma        bool
	Ethernet    bool
	Gsm         bool
	Infiniband  bool
	Loopback    bool
	OlpcMesh    bool
	Pppoe       bool
	Team        bool
	TeamSlave   bool
	Vlan        bool
	Vpn         bool
	Wifi        bool
	Wimax       bool
}

func (i IncludedDevices) Includes(deviceType string) bool {
	include := map[string]bool{
		"ethernet":     i.Ethernet,
		"wifi":         i.Wifi,
		"wimax":        i.Wimax,
		"pppoe":        i.Pppoe,
		"gsm":          i.Gsm,
		"cdma":         i.Cdma,
		"infiniband":   i.Infiniband,
		"bluetooth":    i.Bluetooth,
		"vlan":         i.Vlan,
		"bond":         i.Bond,
		"bond-slave":   i.BondSlave,
		"team":         i.Team,
		"team-slave":   i.TeamSlave,
		"bridge":       i.Bridge,
		"bridge-slave": i.BridgeSlave,
		"vpn":          i.Vpn,
		"olpc-mesh":    i.OlpcMesh,
		"loopback":     i.Loopback,
	}
	if value, exists := include[deviceType]; exists {
		return value
	} else {
		return false
	}
}

var AllDevices = IncludedDevices{
	Bluetooth:   true,
	Bond:        true,
	BondSlave:   true,
	Bridge:      true,
	BridgeSlave: true,
	Cdma:        true,
	Ethernet:    true,
	Gsm:         true,
	Infiniband:  true,
	Loopback:    true,
	OlpcMesh:    true,
	Pppoe:       true,
	Team:        true,
	TeamSlave:   true,
	Vlan:        true,
	Vpn:         true,
	Wifi:        true,
	Wimax:       true,
}

// ActionDevices completes network devices
//
//	lo (loopback)
//	wlp3s0 (wifi)
func ActionDevices(includedDevices IncludedDevices) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if _, err := execlog.LookPath("nmcli"); err == nil {
			return carapace.ActionExecCommand("nmcli", "--terse", "--fields", "device,type", "device", "status")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				vals := make([]string, 0)
				for _, line := range lines[:len(lines)-1] {
					parts := strings.Split(line, ":")
					if includedDevices.Includes(parts[1]) {
						vals = append(vals, parts[0], parts[1])
					}
				}
				return carapace.ActionValuesDescribed(vals...)
			})
		} else if _, err := execlog.LookPath("ifconfig"); err == nil {
			// fallback to basic ifconfig if nmcli is not available
			return carapace.ActionExecCommand("ifconfig")(func(output []byte) carapace.Action {
				interfaces := []string{}
				r := regexp.MustCompile("^[0-9a-zA-Z]")
				for _, line := range strings.Split(string(output), "\n") {
					if matches := r.MatchString(line); matches {
						interfaces = append(interfaces, strings.Split(line, ":")[0])
					}
				}
				return carapace.ActionValues(interfaces...)
			})
		}
		return carapace.ActionMessage("neither nmcli nor ifconfig available")
	}).Tag("network devices")
}

// ActionConnections completes stored network connections
//
//	somewifi (802-11-wireless abcdefgh-hijk-lmnop-qert-012345678902)
//	private (vpn abcdefgh-hijk-lmnop-qert-012345678901)
func ActionConnections() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("nmcli", "--terse", "connection", "show")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, (len(lines)-1)*2)
			for index, line := range lines[:len(lines)-1] {
				parts := strings.Split(line, ":")
				vals[index*2] = parts[0]
				vals[index*2+1] = parts[2] + " " + parts[1]
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("network connections")
}

// ActionBssids completes BSSID's of local wifi networks
//
//	AA:BB:CC:DD:EE:FF (somewifi)
//	AA:BB:CC:DD:EE:11 (anotherwifi)
func ActionBssids() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("nmcli", "--terse", "--fields", "bssid,ssid,bars", "device", "wifi", "list")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				mac := strings.Replace(line[:22], `\:`, `:`, -1)
				splitted := strings.Split(line[23:], ":")
				if name := splitted[0]; name != "" {
					vals = append(vals, mac, splitted[0], styleForBars(splitted[1]))
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	}).Tag("bssids")
}

// ActionSsids completes SSID's of local wifi networks
//
//	somewifi (AA:BB:CC:DD:EE:FF)
//	anotherwifi (AA:BB:CC:DD:EE:11)
func ActionSsids() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("nmcli", "--terse", "--fields", "bssid,ssid,bars", "device", "wifi", "list")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				mac := strings.Replace(line[:22], `\:`, `:`, -1)
				splitted := strings.Split(line[23:], ":")
				if name := splitted[0]; name != "" {
					vals = append(vals, splitted[0], mac, styleForBars(splitted[1]))
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	}).Tag("ssids")
}

func styleForBars(s string) string {
	switch s {
	case "▂___":
		return style.Blue
	case "▂▄__":
		return style.Magenta
	case "▂▄▆_":
		return style.Yellow
	case "▂▄▆█":
		return style.Green
	default:
		return style.Default
	}
}
