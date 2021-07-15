package git

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
)

type ChangeOption struct {
	Staged   bool
	Unstaged bool
}

// ActionChanges completes (un)staged changes
//   fileA ( M)
//   pathA/fileB (??)
func ActionChanges(opts ChangeOption) carapace.Action {
	// TODO multiparts action to complete step by step
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("git", "status", "--porcelain")(func(output []byte) carapace.Action {
			if root, err := rootDir(); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				if wd, err := os.Getwd(); err != nil {
					return carapace.ActionMessage(err.Error())
				} else {
					untracked := make([]string, 0)
					for _, line := range strings.Split(string(output), "\n") {
						if len(line) > 3 {
							if (opts.Staged && line[1] == ' ') ||
								(opts.Unstaged && line[1] != ' ') {
								if relativePath, err := filepath.Rel(wd, root+"/"+line[3:]); err != nil {
									return carapace.ActionMessage(err.Error())
								} else {
									untracked = append(untracked, relativePath, line[:2])
								}
							}
						}
					}
					return carapace.ActionValuesDescribed(untracked...)
				}
			}
		})
	})
}
