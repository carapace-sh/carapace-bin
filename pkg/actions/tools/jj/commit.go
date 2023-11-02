package jj

import (
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
)

// ActionRecentCommits completes
//
//	3b61f0a729f7 (compat: added cobra bridge)
//	3c2e7535ce2f (elivsh: testing tag support)
func ActionRecentCommits(limit int) carapace.Action {
	return carapace.ActionExecCommand("jj", "log", "--no-graph", "--template", `commit_id.short() ++ "\n" ++ description.first_line() ++ "\n"`, "--limit", strconv.Itoa(limit))(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValuesDescribed(lines[:len(lines)-1]...).Style(styles.Git.Commit)
	}).Tag("recent commits")
}
