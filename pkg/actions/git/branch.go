package git

import (
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionRemoteBranches completes remote branches
//   master (last commit msg)
//   remoteBranch (last commit msg)
func ActionRemoteBranches(remote string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if branches, err := branches(RefOption{RemoteBranches: true}); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)
			for _, branch := range branches {
				if strings.HasPrefix(branch.Name, remote) {
					vals = append(vals, strings.TrimPrefix(branch.Name, remote+"/"), branch.Message)
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}

// ActionCurrentBranch completes the current branch
//   currentBranch
func ActionCurrentBranch() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if output, err := exec.Command("git", "branch", "--show-current").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionValues(strings.Split(string(output), "\n")[0])
		}
	})
}

type branch struct {
	Name    string
	Message string
}

func branches(refOption RefOption) ([]branch, error) {
	args := []string{"branch", "--format", "%(refname)\n%(subject)"}
	if refOption.LocalBranches && refOption.RemoteBranches {
		args = append(args, "--all")
	} else if !refOption.LocalBranches && refOption.RemoteBranches {
		args = append(args, "--remote")
	} else if !refOption.LocalBranches && !refOption.RemoteBranches {
		return []branch{}, nil
	}

	if output, err := exec.Command("git", args...).Output(); err != nil {
		return nil, err
	} else {
		lines := strings.Split(string(output), "\n")
		branches := make([]branch, len(lines)/2)
		for index, line := range lines[:len(lines)-1] {
			if index%2 == 0 {
				trimmed := strings.TrimPrefix(line, "refs/heads/")
				trimmed = strings.TrimPrefix(trimmed, "refs/remotes/")
				branches[index/2] = branch{trimmed, lines[index+1]}
			}
		}
		return branches, err
	}
}
