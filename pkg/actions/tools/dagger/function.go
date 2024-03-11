package dagger

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionFunctions completes functions
//
//	container-echo (dagger call container-echo --string-arg yo stdout)
//	grep-dir (dagger call grep-dir --directory-arg . --pattern GrepDir)
func ActionFunctions() carapace.Action {
	return carapace.ActionExecCommand("dagger", "functions", "--silent")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^(?P<name>[^ ]+) +example usage: "(?P<example>.*)"$`)

		vals := make([]string, 0)
		for _, line := range lines[1:] {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
