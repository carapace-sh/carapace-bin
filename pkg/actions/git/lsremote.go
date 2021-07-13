package git

import (
	"strings"

	"github.com/rsteube/carapace"
)

type LsRemoteRefOption struct {
	Branches bool
	Tags     bool
}

func ActionLsRemoteRefs(url string, opts LsRemoteRefOption) carapace.Action {
	return carapace.ActionExecCommand("git", "ls-remote", "--refs", "--tags", "--heads", url)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			fields := strings.Fields(line)
			if opts.Branches && strings.HasPrefix(fields[1], "refs/heads/") {
				vals = append(vals, strings.TrimPrefix(fields[1], "refs/heads/"), fields[0])
			} else if opts.Tags && strings.HasPrefix(fields[1], "refs/tags/") {
				vals = append(vals, strings.TrimPrefix(fields[1], "refs/tags/"), fields[0])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
