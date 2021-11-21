// package tshark contains wireshark related actions
package tshark

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

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
