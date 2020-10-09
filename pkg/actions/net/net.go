// package net contains network related actions
package net

import (
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"

	"github.com/mitchellh/go-homedir"

	"github.com/rsteube/carapace"
)

// ActionHosts completes known hosts
//   192.168.1.1
//   pihole
func ActionHosts() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		hosts := []string{}
		if file, err := homedir.Expand("~/.ssh/known_hosts"); err == nil {
			if content, err := ioutil.ReadFile(file); err == nil {
				r := regexp.MustCompile(`^(?P<host>[^ ,#]+)`)
				for _, entry := range strings.Split(string(content), "\n") {
					if r.MatchString(entry) {
						hosts = append(hosts, r.FindStringSubmatch(entry)[0])
					}
				}
			} else {
				return carapace.ActionMessage(err.Error())
			}
		}
		return carapace.ActionValues(hosts...)
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
//   lo (loopback)
//   wlp3s0 (wifi)
func ActionDevices(includedDevices IncludedDevices) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if _, err := exec.LookPath("nmcli"); err == nil {
			if output, err := exec.Command("nmcli", "--terse", "--fields", "device,type", "device", "status").Output(); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				lines := strings.Split(string(output), "\n")
				vals := make([]string, 0)
				for _, line := range lines[:len(lines)-1] {
					parts := strings.Split(line, ":")
					if includedDevices.Includes(parts[1]) {
						vals = append(vals, parts[0], parts[1])
					}
				}
				return carapace.ActionValuesDescribed(vals...)
			}
		} else if _, err := exec.LookPath("ifconfig"); err == nil {
			// fallback to basic ifconfig if nmcli is not available
			if output, err := exec.Command("ifconfig").Output(); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				interfaces := []string{}
				for _, line := range strings.Split(string(output), "\n") {
					if matches, _ := regexp.MatchString("^[0-9a-zA-Z]", line); matches {
						interfaces = append(interfaces, strings.Split(line, ":")[0])
					}
				}
				return carapace.ActionValues(interfaces...)
			}
		}
		return carapace.ActionMessage("neither nmcli nor ifconfig available")
	})
}

// ActionConnections completes stored network connections
//   somewifi (802-11-wireless abcdefgh-hijk-lmnop-qert-012345678902)
//   private (vpn abcdefgh-hijk-lmnop-qert-012345678901)
func ActionConnections() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("nmcli", "--terse", "connection", "show").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, (len(lines)-1)*2)
			for index, line := range lines[:len(lines)-1] {
				parts := strings.Split(line, ":")
				vals[index*2] = parts[0]
				vals[index*2+1] = parts[2] + " " + parts[1]
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}

// ActionBssids completes BSSID's of local wifi networks
//   AA:BB:CC:DD:EE:FF (somewifi)
//   AA:BB:CC:DD:EE:11 (anotherwifi)
func ActionBssids() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("nmcli", "--terse", "--fields", "bssid,ssid", "device", "wifi", "list").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, (len(lines)-1)*2)
			for index, line := range lines[:len(lines)-1] {
				vals[index*2] = line[:22]
				vals[index*2+1] = line[23:]
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}

// ActionSsids completes SSID's of local wifi networks
//   somewifi (AA:BB:CC:DD:EE:FF)
//   anotherwifi (AA:BB:CC:DD:EE:11)
func ActionSsids() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("nmcli", "--terse", "--fields", "bssid,ssid", "device", "wifi", "list").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				if ssid := line[23:]; ssid != "" {
					vals = append(vals, line[23:], line[:22])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}
