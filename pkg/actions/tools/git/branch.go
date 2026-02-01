package git

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

// ActionCurrentBranch completes the current branch
//
//	currentBranch
func ActionCurrentBranch() carapace.Action {
	return carapace.ActionExecCommand("git", "branch", "--show-current")(func(output []byte) carapace.Action {
		return carapace.ActionValues(strings.Split(string(output), "\n")[0])
	})
}

// ActionLocalBranches completes local branches
//
//	master (last commit msg)
//	another (last commit msg)
func ActionLocalBranches() carapace.Action {
	return carapace.ActionExecCommand("git", "branch", "--format", "%(refname:short)\n%(subject)")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		if strings.HasPrefix(lines[0], "(") {
			lines = lines[2:] // skip detached HEAD
		}
		return carapace.ActionValuesDescribed(lines[:len(lines)-1]...).Style(styles.Git.Branch)
	}).Tag("local branches").UidF(Uid("local-branch")).QueryF(Uid("local-branch"))
}

// ActionRemoteBranches completes remote branches
//
//	origin/master (last commit msg)
//	upstream/another (last commit msg)
func ActionRemoteBranches(remote string) carapace.Action {
	return carapace.ActionExecCommand("git", "branch", "--remote", "--format", "%(refname:short)\n%(subject)")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		prefix := ""
		if remote != "" {
			prefix = remote + "/"
		}

		vals := make([]string, 0)
		for index, line := range lines[:len(lines)-1] {
			if index%2 == 0 && strings.HasPrefix(line, prefix) {
				vals = append(vals, line, lines[index+1])
			}
		}
		return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Branch)
	}).Tag("remote branches").UidF(Uid("remote-branch")).QueryF(Uid("remote-branch"))
}

// ActionRemoteBranchNames is like ActionRemoteBranches but skips the remote prefix
//
//	master (last commit msg)
//	another (last commit msg)
func ActionRemoteBranchNames(remote string) carapace.Action {
	return carapace.ActionExecCommand("git", "branch", "--remote", "--format", "%(refname:short)\n%(subject)")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		prefix := ""
		if remote != "" {
			prefix = remote + "/"
		}

		vals := make([]string, 0)
		for index, line := range lines[:len(lines)-1] {
			if index%2 == 0 && strings.HasPrefix(line, prefix) {
				if _, branch, ok := strings.Cut(line, "/"); ok {
					vals = append(vals, branch, lines[index+1])
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Branch)
	}).Tag("remote branch names")
}
