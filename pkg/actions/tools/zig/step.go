package zig

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionSteps completes top level steps of the current build
//
//	install (Copy build artifacts to prefix path)
//	uninstall (Remove build artifacts from prefix path)
func ActionSteps() carapace.Action {
	return carapace.ActionExecCommand("zig", "build", "--list-steps")(func(output []byte) carapace.Action {
		re := regexp.MustCompile(`^  (\S+?)( \(default\))? +(\S.*)$`)
		values := make([]string, 0)
		for _, line := range strings.Split(string(output), "\n") {
			if match := re.FindStringSubmatch(line); match != nil {
				values = append(values, match[1], match[3])
			}
		}
		return carapace.ActionValuesDescribed(values...)
	}).Tag("steps").Uid("zig", "step")
}
