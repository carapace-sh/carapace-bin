package jj

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

// ActionLocalBranches completes local branches
//
//	master (qvlkomxp 9a2e553b (empty) Merge pull request #939 from rsteube/docker-update-...)
//	mdbook-linkcheck (numonpmw ad5c8efd mdbook: enable web-links)
func ActionLocalBranches() carapace.Action {
	return carapace.ActionExecCommand("jj", "branch", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		rLocal := regexp.MustCompile(`^(?P<branch>[^@: ]+): (?P<changeid>[^ ]+) (?P<commitid>[^ ]+) (?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			if matches := rLocal.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[4])
			}
		}
		return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Branch)
	}).Tag("local branches")
}

// ActionRemoteBranches completes remote branches
//
//	master@git (qvlkomxp 9a2e553b (empty) Merge pull request #939 from rsteube/docker-update-...)
//	master@origin (qvlkomxp 9a2e553b (empty) Merge pull request #939 from rsteube/docker-update-...)
func ActionRemoteBranches(remote string) carapace.Action {
	return carapace.ActionExecCommand("jj", "branch", "list", "--all")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		rLocal := regexp.MustCompile(`^(?P<branch>[^@: ]+): (?P<changeid>[^ ]+) (?P<commitid>[^ ]+) (?P<description>.*)$`)
		rRemote := regexp.MustCompile(`^(?P<branch>[^@: ]+)@(?P<remote>[^: ]+): (?P<changeid>[^ ]+) (?P<commitid>[^ ]+) (?P<description>.*)$`)
		rTracking := regexp.MustCompile(`^  @(?P<remote>[^: ]+): (?P<changeid>[^ ]+) (?P<commitid>[^ ]+) (?P<description>.*)$`)

		vals := make([]string, 0)
		branch := ""
		for _, line := range lines[:len(lines)-1] {
			switch {
			case strings.HasPrefix(line, "  @"):
				if matches := rTracking.FindStringSubmatch(line); matches != nil {
					vals = append(vals, fmt.Sprintf("%v@%v", branch, matches[1]), matches[4])
				}
			default:
				if matches := rRemote.FindStringSubmatch(line); matches != nil {
					vals = append(vals, fmt.Sprintf("%v@%v", matches[1], matches[2]), matches[5])
				} else if matches := rLocal.FindStringSubmatch(line); matches != nil {
					branch = matches[1]
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Branch)
	}).Tag("remote branches")
}
