package git

import (
	"net/url"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/carapace-sh/carapace/pkg/uid"
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
				evaluatedDir, err := filepath.EvalSymlinks(c.Dir)
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}

				untracked := make([]string, 0)
				for _, line := range strings.Split(string(output), "\n") {
					switch {
					case len(line) < 4:
						// skip
					case opts.Staged && line[1] == ' ',
						opts.Unstaged && line[1] != ' ' && line[1] != '!',
						opts.Ignored && line[1] == '!':

						path := line[3:]
						if splitted := strings.SplitN(path, " -> ", 2); len(splitted) > 1 { // renamed
							path = splitted[1]
						}

						relativePath, err := filepath.Rel(evaluatedDir, filepath.Join(root, path))
						if err != nil {
							return carapace.ActionMessage(err.Error())
						}

						switch status := line[:2]; {
						case strings.Contains(status, "D"): // deleted
							untracked = append(untracked, relativePath, status, style.ForPathExt(relativePath, c))
						default:
							untracked = append(untracked, relativePath, status, style.ForPath(relativePath, c))
						}
					}
				}

				action := carapace.ActionStyledValuesDescribed(untracked...)
				if strings.HasPrefix(c.Value, "./") {
					action = action.Prefix("./")
				}
				return action.UidF(func(s string, uc uid.Context) (*url.URL, error) {
					abs, err := filepath.Abs(filepath.Join(evaluatedDir, s))
					if err != nil {
						return nil, err
					}
					rel, err := filepath.Rel(root, abs)
					if err != nil {
						return nil, err
					}
					return Uid("change")(rel, uc)
				})
			}
		})
	}).Tag("changes")
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
	}).Tag("ref changes")
}
