package man

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPages completes man pages
//
//	git (the stupid content tracker)
//	git-add (Add file contents to the index)
func ActionPages() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("man", "-k", c.Value)(func(output []byte) carapace.Action {
			r := regexp.MustCompile(`^(?P<name>.*?) ?\(\d+\).* +- (?P<description>.*)$`)

			vals := make([]string, 0)
			for _, line := range strings.Split(string(output), "\n") {
				if r.MatchString(line) {
					matches := r.FindStringSubmatch(line)
					vals = append(vals, matches[1], matches[2])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
