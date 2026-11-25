// package net contains network related actions
package net

import (
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-ssh/knownhosts"
	"github.com/carapace-sh/carapace/pkg/execlog"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionHosts completes known hosts
//
//	192.168.1.1
//	pihole
func ActionHosts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		opts := make([]knownhosts.Option, 0, 1)
		if home, ok := c.LookupEnv("HOME"); ok && home != "" {
			opts = append(opts, knownhosts.WithHomeDir(home))
		} else if home, err := os.UserHomeDir(); err == nil && home != "" {
			opts = append(opts, knownhosts.WithHomeDir(home))
		}

		known, _ := knownhosts.IndexDefault(opts...)
		hosts := known.HostCandidates()
		if len(hosts) == 0 {
			return carapace.ActionValues().Tag("known hosts")
		}

		ipv4 := make([]string, 0)
		ipv6 := make([]string, 0)
		names := make([]string, 0)
		for _, host := range hosts {
			ip := net.ParseIP(host)
			switch {
			case ip == nil:
				names = append(names, host)
			case ip.To4() != nil:
				ipv4 = append(ipv4, host)
			default:
				ipv6 = append(ipv6, host)
			}
		}

		batch := carapace.Batch()
		if len(ipv4) > 0 {
			batch = append(batch, carapace.ActionStyledValues(ipv4...).Style(style.Default).Tag("ipv4 addresses"))
		}
		if len(ipv6) > 0 {
			batch = append(batch, carapace.ActionStyledValues(ipv6...).Style(style.Bold).Tag("ipv6 addresses"))
		}
		if len(names) > 0 {
			batch = append(batch, carapace.ActionStyledValues(names...).Style(style.Blue).Tag("hostnames"))
		}
		return batch.ToA()
	}).Tag("known hosts")
}

// ActionUserHosts completes known hosts with a user prefix
//
//	user@192.168.1.1
//	user@pihole
func ActionUserHosts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if prefix, _, ok := strings.Cut(c.Value, "@"); ok {
			return ActionHosts().Prefix(prefix + "@")
		}
		return ActionHosts()
	})
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

func (i IncludedDevices) Default() IncludedDevices {
	return AllDevices
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
				for line := range strings.SplitSeq(string(output), "\n") {
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
				mac := strings.ReplaceAll(line[:22], `\:`, `:`)
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
				mac := strings.ReplaceAll(line[:22], `\:`, `:`)
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
