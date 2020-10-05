package fs

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionMounts() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("mount").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			re := regexp.MustCompile(`^(?P<target>\S+) on (?P<mountt>\S+) type (?P<type>\S+) (?P<mode>.+)`)
			mounts := make([]string, 0)
			for _, line := range strings.Split(string(output), "\n") {
				if re.MatchString(line) {
					matches := re.FindStringSubmatch(line)
					mounts = append(mounts, matches[2], matches[1])
				}
			}
			return carapace.ActionValuesDescribed(mounts...)
		}
	})
}
