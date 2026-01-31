package jj

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

// ActionHeadCommits completes head commits
//
//	@  (HEAD)
//	@- (commit message)
func ActionHeadCommits(limit int) carapace.Action {
	return actionExecJJ("log", "--no-graph", "--template", `description.first_line() ++ "\n"`, "--revisions", "::@", "--limit", strconv.Itoa(limit))(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for index, line := range lines[:len(lines)-1] {
			vals = append(vals, "@"+strings.Repeat("-", index), line)
		}
		return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Head)
	}).Tag("head commits")
}

// ActionPrevCommits completes head commits
//
//	04 (build(deps): bump github.com/creack/pty from 1.1.18 to 1.1.20 (#8265))
//	03 (fix(release create): Handle latest flag value when updating the rel...)
func ActionPrevCommits(limit int) carapace.Action {
	return actionExecJJ("log", "--no-graph", "--template", `commit_id.short() ++ "\t" ++ description.first_line() ++ "\n"`, "--revisions", "::@", "--limit", strconv.Itoa(limit))(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		format := "%0" + strconv.Itoa(len(strconv.Itoa(limit-1))) + "d"

		vals := make([]string, 0)
		for index, line := range lines[1 : len(lines)-1] {
			splitted := strings.SplitN(line, "\t", 2)
			vals = append(vals, fmt.Sprintf(format, index), splitted[1])
		}
		return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Head)
	}).Tag("previous commits")
}

// ActionNextCommits completes next commits
func ActionNextCommits(limit int) carapace.Action {
	return actionExecJJ("log", "--no-graph", "--template", `commit_id.short() ++ "\t" ++ description.first_line() ++ "\n"`, "--revisions", "@-::", "--limit", strconv.Itoa(limit))(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		format := "%0" + strconv.Itoa(len(strconv.Itoa(limit-1))) + "d"

		vals := make([]string, 0)
		for index, line := range lines[1 : len(lines)-1] {
			splitted := strings.SplitN(line, "\t", 2)
			vals = append(vals, fmt.Sprintf(format, len(lines)-3-index), splitted[1])
		}
		return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Head)
	}).Tag("previous commits")
}

// ActionRecentCommits completes recent commits
//
//	3b61f0a729f7 (compat: added cobra bridge)
//	3c2e7535ce2f (elivsh: testing tag support)
func ActionRecentCommits(limit int) carapace.Action {
	return actionExecJJ("log", "--no-graph", "--template", `commit_id.short() ++ "\n" ++ description.first_line() ++ "\n"`, "--limit", strconv.Itoa(limit))(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValuesDescribed(lines[:len(lines)-1]...).Style(styles.Git.Commit)
	}).Tag("recent commits")
}
