package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionFlagValues(flag string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("gcc", "--completion="+flag+"="+c.Value)(func(output []byte) carapace.Action {
			lines := strings.Split(strings.TrimRight(string(output), "\n"), "\n")
			vals := make([]string, 0)
			prefix := flag + "="
			for _, line := range lines {
				if line == "" {
					continue
				}
				if strings.HasPrefix(line, prefix) {
					vals = append(vals, strings.TrimPrefix(line, prefix))
				}
			}
			if len(vals) == 0 {
				return carapace.ActionValues()
			}
			return carapace.ActionValues(vals...)
		})
	})
}

// ActionDynamic completes all gcc options dynamically using gcc --completion=PREFIX
func ActionDynamic(prefix string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("gcc", "--completion="+prefix+c.Value)(func(output []byte) carapace.Action {
			lines := strings.Split(strings.TrimRight(string(output), "\n"), "\n")
			vals := make([]string, 0)
			seen := make(map[string]bool)
			for _, line := range lines {
				if line == "" || seen[line] {
					continue
				}
				seen[line] = true
				vals = append(vals, line)
			}
			if len(vals) == 0 {
				return carapace.ActionValues()
			}
			return carapace.ActionValues(vals...)
		})
	})
}