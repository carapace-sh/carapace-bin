package zig

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionTargets completes target triples
//
//	arc-linux-gnu
//	x86_64-linux-musl
func ActionTargets() carapace.Action {
	return carapace.ActionMultiParts("-", func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("zig", "targets")(func(output []byte) carapace.Action {
			lists := parseTargets(string(output))
			var values []string
			switch len(c.Parts) {
			case 0:
				values = lists["arch"]
			case 1:
				values = lists["os"]
			case 2:
				values = lists["abi"]
			}
			return carapace.ActionValues(values...).Invoke(c).ToA()
		})
	}).Tag("targets").Uid("zig", "target")
}

func parseTargets(output string) map[string][]string {
	re := regexp.MustCompile(`(?ms)^\s+\.(\w+) = \.\{$(.*?)^    \},$`)
	lists := make(map[string][]string)
	for _, match := range re.FindAllStringSubmatch(strings.TrimSpace(output), -1) {
		values := make([]string, 0)
		for _, entry := range regexp.MustCompile(`"([\w-]+)"`).FindAllStringSubmatch(match[2], -1) {
			values = append(values, entry[1])
		}
		lists[match[1]] = values
	}
	return lists
}
