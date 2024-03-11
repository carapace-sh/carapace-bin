package git

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionDeveloperInterfaces completes developer interfaces
//
//	format-bundle (The bundle file format)
//	format-chunk (Chunk-based file formats)
func ActionDeveloperInterfaces() carapace.Action {
	return carapace.ActionExecCommand("git", "help", "--developer-interfaces")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^   (?P<interfaces>[^ ]+) +(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("developer interfaces")
}
