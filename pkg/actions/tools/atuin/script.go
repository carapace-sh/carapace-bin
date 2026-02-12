package atuin

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionScripts completes scripts
//
//	one
//	two
func ActionScripts() carapace.Action {
	return carapace.ActionExecCommand("atuin", "scripts", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^- (?P<name>[^ ]+) (\[tags: (?P<tags>[^]]+)\])?$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[3])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("scripts")
}

// ActionTags completes tags
//
//	one
//	two
func ActionTags() carapace.Action {
	return carapace.ActionExecCommand("atuin", "scripts", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^- (?P<name>[^ ]+) (\[tags: (?P<tags>[^]]+)\])?$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				if matches[3] == "" {
					continue
				}
				for tag := range strings.SplitSeq(matches[3], ", ") {
					vals = append(vals, tag)
				}
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("tags")
}
