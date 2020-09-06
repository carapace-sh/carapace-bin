package action

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
)

func rootDir() (string, error) {
	if output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output(); err != nil {
		return "", err
	} else {
		return strings.Split(string(output), "\n")[0], nil
	}
}

func ActionRemotes() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("git", "remote").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionValues(strings.Split(string(output), "\n")...)
		}
	})
}

func ActionRemoteBranches() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("git", "branch", "--remote", "--format", "%(refname)").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			refs := make([]string, 0)
			for _, line := range strings.Split(string(output), "\n") {
				if len(line) > 14 {
					refs = append(refs, line[13:])
				}
			}
			return carapace.ActionValues(refs...)
		}
	})
}

// TODO multiparts action to complete step by step
func ActionUnstagedChanges() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("git", "status", "--porcelain").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			if root, err := rootDir(); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				if wd, err := os.Getwd(); err != nil {
					return carapace.ActionMessage(err.Error())
				} else {
					untracked := make([]string, 0)
					for _, line := range strings.Split(string(output), "\n") {
						if len(line) > 3 && line[1] != ' ' { // no blank line and not already staged
							if relativePath, err := filepath.Rel(wd, root+"/"+line[3:]); err != nil {
								return carapace.ActionMessage(err.Error())
							} else {
								untracked = append(untracked, relativePath, line[:2])
							}
						}
					}
					return carapace.ActionValuesDescribed(untracked...)
				}
			}
		}
	})
}
