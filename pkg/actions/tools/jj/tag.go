package jj

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

// ActionTags completes tags
//
//	v0.0.4 (added prefix to functions)
//	v0.0.5 (fish: fix test)
func ActionTags() carapace.Action {
	return actionExecJJ("log", "--no-graph", "--revisions", "tags()", "--template", `tags ++ "\t" ++ description.first_line() ++ "\n"`)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			splitted := strings.SplitN(line, "\t", 2)
			for _, tag := range strings.Split(splitted[0], " ") {
				vals = append(vals, tag, splitted[1])
			}
		}
		return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Tag)
	}).Tag("tags")
}
