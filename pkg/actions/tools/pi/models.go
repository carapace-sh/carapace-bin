package pi

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionModels completes models
//
//	claude-sonnet-4-5
//	gpt-4o
func ActionModels() carapace.Action {
	return carapace.ActionExecCommand("pi", "--list-models")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		re := regexp.MustCompile(`\s{2,}`)
		for _, line := range lines[1:] {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			if splitted := re.Split(line, -1); len(splitted) >= 2 {
				vals = append(vals, splitted[1], splitted[0])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}