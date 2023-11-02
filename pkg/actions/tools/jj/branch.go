package jj

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
)

// ActionLocalBranches completes local branches
//
//	master (qvlkomxp 9a2e553b (empty) Merge pull request #939 from rsteube/docker-update-...)
//	mdbook-linkcheck (numonpmw ad5c8efd mdbook: enable web-links)
func ActionLocalBranches() carapace.Action {
	return carapace.ActionExecCommand("jj", "branch", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, strings.SplitN(line, ":", 2)...)
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

		vals := make([]string, 0)
		branch := ""
		for _, line := range lines[:len(lines)-1] {
			switch {
			case strings.HasPrefix(line, "  @"):
				splitted := strings.SplitN(line, ":", 2)
				switch remote {
				case "", strings.TrimPrefix(splitted[0], "  @"):
					vals = append(vals, branch+strings.TrimSpace(splitted[0]), splitted[1])
				}
			default:
				branch = strings.SplitN(line, ":", 2)[0]
			}
		}
		return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Branch)
	}).Tag("remote branches")
}
