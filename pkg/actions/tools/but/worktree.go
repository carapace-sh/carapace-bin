package but

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/style"
)

type butWorktree struct {
	ID      *string `json:"id"`
	Name    string  `json:"name"`
	RefName *string `json:"refName"`
	Path    string  `json:"path"`
}

func (w butWorktree) ref() string {
	if w.RefName == nil {
		return ""
	}
	return *w.RefName
}

type worktreeListOutput struct {
	Active   []butWorktree `json:"active"`
	Archived []butWorktree `json:"archived"`
}

type WorktreesOpts struct {
	Archived bool // include archived worktrees
}

func (o WorktreesOpts) Default() WorktreesOpts {
	o.Archived = true
	return o
}

// ActionWorktrees completes worktrees
//
//	w1 (my-branch)
//	docs (docs-site)
func ActionWorktrees(opts WorktreesOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"--json", "worktree", "list", "--active"}
		if opts.Archived {
			args = append(args, "--archived")
		}
		return carapace.ActionExecCommand("but", args...)(func(output []byte) carapace.Action {
			var list worktreeListOutput
			if err := json.Unmarshal(output, &list); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, w := range append(list.Active, list.Archived...) {
				if w.ID != nil {
					vals = append(vals, *w.ID, w.Name, style.Default)
				}
				if ref := w.ref(); ref != "" && ref != w.Name {
					vals = append(vals, w.Name, ref, style.Default)
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...).Tag("worktrees").UidF(git.Uid("worktree"))
		})
	})
}
