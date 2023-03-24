package git

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
)

// ActionHeadCommits completes recent head commits.
//
//	HEAD (commit message)
//	HEAD~1 (commit message)
func ActionHeadCommits(limit int) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"log", "--no-notes", "--pretty=tformat:%h   %<(64,trunc)%s", "--max-count", strconv.Itoa(limit)}
		return carapace.ActionExecCommand("git", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for index, line := range lines[:len(lines)-1] {
				if index == 0 {
					vals = append(vals, "HEAD", strings.TrimSpace(line[10:]))
				} else {
					vals = append(vals, "HEAD~"+fmt.Sprintf("%0"+strconv.Itoa(len(strconv.Itoa(limit-1)))+"d", index), strings.TrimSpace(line[10:]))
				}
			}
			return carapace.ActionValuesDescribed(vals...).Style(styles.Git.HeadCommit)
		})
	}).Tag("head commits")
}

// ActionRecentCommits completes recent commits
//
//	123456A ((refname) commit message)
//	123456B ((refname) commit message)
func ActionRecentCommits(limit int) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"log", "--exclude=refs/notes/*", "--all", "--no-merges", "--pretty=tformat:%h   %<(64,trunc)%s", "--max-count", strconv.Itoa(limit)}
		return carapace.ActionExecCommand("git", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				vals = append(vals, line[:7], strings.TrimSpace(line[10:]))
			}
			return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Commit)
		})
	}).Tag("recent commits")
}
