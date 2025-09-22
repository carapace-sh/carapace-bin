package git

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

// ActionTags completes tags
//
//	v0.0.1
//	v0.0.2
func ActionTags() carapace.Action {
	return carapace.ActionExecCommand("git", "tag", "--format", "%(refname)\n%(subject)")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for index, line := range lines[:len(lines)-1] {
			if index%2 == 0 {
				vals = append(vals, strings.TrimPrefix(line, "refs/tags/"), lines[index+1])
			}
		}
		return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Tag)
	}).Tag("tags").UidF(Uid("tag")).QueryF(Uid("tag"))
}
