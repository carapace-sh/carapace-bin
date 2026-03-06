package wt

import (
	"encoding/json"
	"net/url"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace/pkg/uid"
)

type worktree struct {
	Branch string `json:"branch"`
	Path   string `json:"path"`
	Kind   string `json:"kind"`
	Commit struct {
		Sha       string `json:"sha"`
		ShortSha  string `json:"short_sha"`
		Message   string `json:"message"`
		Timestamp int    `json:"timestamp"`
	} `json:"commit"`
	WorkingTree struct {
		Staged    bool `json:"staged"`
		Modified  bool `json:"modified"`
		Untracked bool `json:"untracked"`
		Renamed   bool `json:"renamed"`
		Deleted   bool `json:"deleted"`
		Diff      struct {
			Added   int `json:"added"`
			Deleted int `json:"deleted"`
		} `json:"diff"`
	} `json:"working_tree"`
	MainState string `json:"main_state"`
	Main      struct {
		Ahead  int `json:"ahead"`
		Behind int `json:"behind"`
	} `json:"main"`
	Worktree struct {
		Detached bool `json:"detached"`
	} `json:"worktree"`
	IsMain     bool   `json:"is_main"`
	IsCurrent  bool   `json:"is_current"`
	IsPrevious bool   `json:"is_previous"`
	Statusline string `json:"statusline"`
	Symbols    string `json:"symbols"`
}

// ActionWorktrees completes worktrees
//
//	first (commit message)
//	second (commit message)
func ActionWorktrees() carapace.Action {
	return carapace.ActionExecCommand("wt", "list", "--format", "json")(func(output []byte) carapace.Action {
		var worktrees []worktree
		if err := json.Unmarshal(output, &worktrees); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		batch := carapace.Batch()
		for _, wt := range worktrees {
			batch = append(batch, carapace.ActionValuesDescribed(wt.Branch, wt.Commit.Message).
				UidF(func(s string, uc uid.Context) (*url.URL, error) {
					return git.Uid("ref")(wt.Commit.ShortSha, uc)
				}),
			)
		}
		return batch.ToA().Style(styles.Git.Branch)
	}).Tag("worktrees")
}
