package git

import (
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

type ChangeOpts struct {
	Staged   bool
	Unstaged bool
}

func (o ChangeOpts) Default() ChangeOpts {
	o.Staged = true
	o.Unstaged = true
	return o
}

// ActionChanges completes (un)staged changes
//   fileA ( M)
//   pathA/fileB (??)
func ActionChanges(opts ChangeOpts) carapace.Action {
	// TODO multiparts action to complete step by step
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("git", "status", "--porcelain")(func(output []byte) carapace.Action {
			if root, err := rootDir(c); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				untracked := make([]string, 0)
				for _, line := range strings.Split(string(output), "\n") {
					if len(line) > 3 {
						if (opts.Staged && line[1] == ' ') ||
							(opts.Unstaged && line[1] != ' ') {
							if relativePath, err := filepath.Rel(c.Dir, root+"/"+line[3:]); err != nil {
								return carapace.ActionMessage(err.Error())
							} else {
								untracked = append(untracked, relativePath, line[:2], style.ForPath(relativePath))
							}
						}
					}
				}
				return carapace.ActionStyledValuesDescribed(untracked...)
			}
		})
	})
}
