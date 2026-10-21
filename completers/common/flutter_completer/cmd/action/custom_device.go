package action

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionCustomDevices() carapace.Action {
	return carapace.ActionExecCommand("flutter", "--suppress-analytics", "custom-devices", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^\s*id: (.*?), label: (.*?), enabled: (?:true|false)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		if len(vals) == 0 {
			if strings.Contains(string(output), "No custom devices found") {
				return carapace.ActionValues()
			}
			return carapace.ActionMessage(strings.TrimSpace(string(output)))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
