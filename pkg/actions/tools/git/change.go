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
	Unmerged bool
	Unstaged bool
}

func (o ChangeOpts) Default() ChangeOpts {
	o.Staged = true
	o.Unstaged = true
	return o
}

// change describes a change of a file (status code and path relative to the repository root)
type change struct {
	status string
	path   string
}

// actionChanges completes given changes
//
//	fileA (M)
//	pathA/fileB (??)
func actionChanges(c carapace.Context, changes []change, styleF func(change, string, style.Context) string, uidF func(change, uid.Context) (*url.URL, error)) carapace.Action {
	root, err := rootDir(c)
	if err != nil {
		return carapace.ActionMessage(err.Error())
	}
	evaluatedDir, err := filepath.EvalSymlinks(c.Dir)
	if err != nil {
		return carapace.ActionMessage(err.Error())
	}

	batch := carapace.Batch()
	for _, ce := range changes {
		relativePath, err := filepath.Rel(evaluatedDir, filepath.Join(root, ce.path))
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		action := carapace.ActionValuesDescribed(relativePath, ce.status).
			StyleF(func(s string, sc style.Context) string {
				return styleF(ce, relativePath, sc)
			}).
			UidF(func(s string, uc uid.Context) (*url.URL, error) {
				return uidF(ce, uc)
			})
		switch {
		case strings.HasPrefix(c.Value, "./"):
			action = action.Prefix("./")
		case strings.HasPrefix(c.Value, ".") && !strings.HasPrefix(relativePath, ".."):
			action = action.Prefix("./")
		}
		batch = append(batch, action)
	}
	return batch.ToA()
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
			changes := make([]change, 0)
			for line := range strings.SplitSeq(string(output), "\n") {
				switch {
				case len(line) < 4:
					continue
				case opts.Unmerged && (strings.Contains(line[:2], "U") || line[:2] == "DD" || line[:2] == "AA"),
					opts.Staged && line[1] == ' ',
					opts.Unstaged && line[1] != ' ' && line[1] != '!',
					opts.Ignored && line[1] == '!':

					path := line[3:]
					if renamed := strings.SplitN(path, " -> ", 2); len(renamed) > 1 {
						path = renamed[1]
					}
					changes = append(changes, change{line[:2], path})
				}
			}

			return actionChanges(c, changes,
				func(ce change, path string, sc style.Context) string {
					if strings.Contains(ce.status, "D") { // deleted
						return style.ForPathExt(path, sc)
					}
					return style.ForPath(path, sc)
				},
				func(ce change, uc uid.Context) (*url.URL, error) {
					return Uid("change")(ce.path, uc)
				},
			)
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
