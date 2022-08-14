// package tshark contains wireshark related actions
package tshark

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionSelectors completes selectors
//
//	acdr.media_type (AC DR Media Type)
//	bacapp.vendor_identifier (BACapp Vendor Identifier
func ActionSelectors() carapace.Action {
	return carapace.ActionExecCommand("sh", "-c", "tshark -d . 2>&1 ; true")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^\t(?P<name>[^ ]+) \((?P<description>.*)\)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				vals = append(vals, r.FindStringSubmatch(line)[1:]...)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionStatistics completes statistics
//
//	ancp,tree
//	endpoints,ipv6
func ActionStatistics() carapace.Action {
	return carapace.ActionExecCommand("sh", "-c", "tshark -z help 2>&1")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			if strings.HasPrefix(line, "     ") {
				vals = append(vals, strings.TrimSpace(line))
			}
		}
		return carapace.ActionValues(vals...)
	})
}

// ActionProtocols completes protocols
//
//	aarp (Appletalk Address Resolution Protocol)
//	cmd (Cisco MetaData)
func ActionProtocols() carapace.Action {
	return carapace.ActionExecCommand("sh", "-c", "tshark -d ethertype==0x0800 2>&1 ; true")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^\t(?P<name>[^ ]+) \((?P<description>.*)\)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				vals = append(vals, r.FindStringSubmatch(line)[1:]...)
			}
		}
		return carapace.ActionValuesDescribed(vals...)

	})
}

// ActionFileTypes completes file types
//
//	5views (InfoVista 5View capture)
//	nokiapcap (Nokia tcpdump - pcap)
func ActionFileTypes() carapace.Action {
	return carapace.ActionExecCommand("sh", "-c", "tshark -F 2>&1 ; true")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^    (?P<name>[^ ]+) - (?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				vals = append(vals, r.FindStringSubmatch(line)[1:]...)
			}
		}
		return carapace.ActionValuesDescribed(vals...)

	})
}

// ActionReadFormats completes read file formats
//
//	AIX iptrace (Magic-value-based)
//	Ixia IxVeriWave .vwr Raw Capture (Heuristics-bas
func ActionReadFormats() carapace.Action {
	return carapace.ActionExecCommand("sh", "-c", "tshark -X read_format: 2>&1 ; true")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^    (?P<name>[^-]+) - (?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines[4:] {
			if r.MatchString(line) {
				vals = append(vals, r.FindStringSubmatch(line)[1:]...)
			}
		}
		return carapace.ActionValuesDescribed(vals...)

	})
}

// ActionLinkTypes completes link types for given interface
//
//	E10MB
//	BLUETOOTH_HCI_H4_WITH_PHDR
func ActionLinkTypes(iface string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"-L"}
		if iface != "" {
			args = append(args, "--interface", iface)
		}

		return carapace.ActionExecCommand("tshark", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^  (?P<name>[^ ]+) \((?P<description>.*)\)$`)

			vals := make([]string, 0)
			for _, line := range lines {
				if r.MatchString(line) {
					vals = append(vals, r.FindStringSubmatch(line)[1:]...)
				}
			}
			return carapace.ActionValuesDescribed(vals...)

		})
	})
}

// ActionInterfaces completes interfaces
//
//	ciscodump
//	dpauxmon
func ActionInterfaces() carapace.Action {
	return carapace.ActionExecCommand("tshark", "--list-interfaces")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		r := regexp.MustCompile(`^\d+\. (?P<name>[^ ]+)( (?P<description>.+))?$`)
		for _, line := range lines {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				if len(matches) == 3 {
					vals = append(vals, matches[1], matches[2])
				} else {
					vals = append(vals, matches[1], "")
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
