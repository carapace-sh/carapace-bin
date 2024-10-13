package git

import (
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/carapace-sh/carapace/pkg/util"
)

type ChangeOpts struct {
	Ignored  bool
	Staged   bool
	Unstaged bool
}

func (o ChangeOpts) Default() ChangeOpts {
	o.Staged = true
	o.Unstaged = true
	return o
}

// ActionChanges completes (un)staged changes
//
//	fileA ( M)
//	pathA/fileB (??)
func ActionChanges(opts ChangeOpts) carapace.Action {
	// TODO multiparts action to complete step by step
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"status", "--porcelain"}
		if opts.Ignored {
			args = append(args, "--ignored")
		}
		return carapace.ActionExecCommand("git", args...)(func(output []byte) carapace.Action {
			if root, err := rootDir(c); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				untracked := make([]string, 0)
				for _, line := range strings.Split(string(output), "\n") {
					if len(line) > 3 {
						if (opts.Staged && line[1] == ' ') ||
							(opts.Unstaged && line[1] != ' ' && line[1] != '!') ||
							(opts.Ignored && line[1] == '!') {
							path := line[3:]
							if splitted := strings.SplitN(path, " -> ", 2); len(splitted) > 1 { // renamed
								path = splitted[1]
							}

							evaluatedDir, err := filepath.EvalSymlinks(c.Dir)
							if err != nil {
								return carapace.ActionMessage(err.Error())
							}
							if relativePath, err := filepath.Rel(evaluatedDir, root+"/"+path); err != nil {
								return carapace.ActionMessage(err.Error())
							} else {
								if status := line[:2]; strings.Contains(status, "D") { // deleted
									untracked = append(untracked, relativePath, status, style.ForPathExt(relativePath, c))
								} else {
									untracked = append(untracked, relativePath, status, style.ForPath(relativePath, c))
								}
							}
						}
					}
				}

				action := carapace.ActionStyledValuesDescribed(untracked...)
				if strings.HasPrefix(c.Value, "./") {
					action = action.Prefix("./")
				}
				return action
			}
		})
	}).Tag("changed files")
}

// ActionRefChanges completes changes made in given ref
//
//	go.mod
//	cmd/carapace/main.go
func ActionRefChanges(ref string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("git", "diff-tree", "--name-only", "--no-commit-id", "-r", ref)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			path, err := util.FindReverse(c.Dir, ".git")
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			relativeRoot, err := filepath.Rel(c.Dir, filepath.Dir(path))
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			if relativeRoot == "." {
				relativeRoot = ""
			} else {
				relativeRoot += "/"
			}

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				vals = append(vals, relativeRoot+line)
			}

			return carapace.ActionValues(vals...).StyleF(style.ForPathExt)
		})
	}).Tag("changed files")
}
