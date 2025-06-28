package git

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

// ActionHeadCommits completes recent head commits
//
//	HEAD (commit message)
//	HEAD~1 (commit message)
func ActionHeadCommits(limit int) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"log", "--no-notes", "--first-parent", "--pretty=tformat:%h   %<(64,trunc)%s", "--max-count", strconv.Itoa(limit)}
		return carapace.ActionExecCommand("git", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for index, line := range lines[:len(lines)-1] {
				switch index {
				case 0:
					vals = append(vals, "HEAD", strings.TrimSpace(line[10:]))
				default:
					vals = append(vals, "HEAD~"+fmt.Sprintf("%0"+strconv.Itoa(len(strconv.Itoa(limit-1)))+"d", index), strings.TrimSpace(line[10:]))
				}
			}
			return carapace.ActionValuesDescribed(vals...).Style(styles.Git.HeadCommit)
		})
	}).Tag("head commits").UidF(Uid("ref"))
}

// ActionRefCommits completes commits reachable by given ref
//
//	00 (commit message)
//	01 (commit message)
func ActionRefCommits(ref string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		limit := 100
		args := []string{"log", "--no-notes", "--first-parent", "--pretty=tformat:%h   %<(64,trunc)%s", "--max-count", strconv.Itoa(limit), ref}
		return carapace.ActionExecCommand("git", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for index, line := range lines[:len(lines)-1] {
				vals = append(vals, fmt.Sprintf("%0"+strconv.Itoa(len(strconv.Itoa(limit-1)))+"d", index), strings.TrimSpace(line[10:]))
			}
			return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Commit)
		})
	}).Tag("ref commits")
}

// ActionRefParents completes parents of given ref
//
//	1 (7aca3ebc)
//	2 (4f4f9e93)
func ActionRefParents(ref string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"log", "--no-notes", "--pretty=%P", "--max-count", "1", ref}
		return carapace.ActionExecCommand("git", args...)(func(output []byte) carapace.Action {
			vals := make([]string, 0)
			for index, field := range strings.Fields(string(output)) {
				vals = append(vals, strconv.Itoa(index+1), field)
			}
			return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Commit)
		})
	}).Tag("parent commits")
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
