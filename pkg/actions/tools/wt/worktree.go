package wt

import (
	"encoding/json"
	"net/url"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace/pkg/uid"
)

type worktreeList struct {
	Items []worktree `json:"items"`
}

type worktree struct {
	Branch string `json:"branch"`
	Head   struct {
		Sha         string `json:"sha"`
		ShortSha    string `json:"short_sha"`
		Subject     string `json:"subject"`
		CommittedAt string `json:"committed_at"`
	} `json:"head"`
	Worktree struct {
		Path     string `json:"path"`
		Main     bool   `json:"main"`
		Current  bool   `json:"current"`
		Previous bool   `json:"previous"`
		Detached bool   `json:"detached"`
	} `json:"worktree"`
	Display struct {
		State      string `json:"state"`
		Symbols    string `json:"symbols"`
		Statusline string `json:"statusline"`
	} `json:"display"`
}

type legacyWorktree struct {
	Branch string `json:"branch"`
	Commit struct {
		ShortSha string `json:"short_sha"`
		Message  string `json:"message"`
	} `json:"commit"`
}

// ActionWorktrees completes worktrees
//
//	first (commit message)
//	second (commit message)
func ActionWorktrees() carapace.Action {
	return carapace.ActionExecCommand("wt", "list", "--format", "json")(func(output []byte) carapace.Action {
		var raw json.RawMessage
		if err := json.Unmarshal(output, &raw); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var items []worktree
		if len(raw) > 0 && raw[0] == '{' { // schema 2 envelope (default since 0.79.0, opt-in `[list] json-schema = 2` before)
			var list worktreeList
			if err := json.Unmarshal(raw, &list); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			items = list.Items
		} else { // schema 1 bare array (default before 0.79.0, retained via `[list] json-schema = 1`)
			var legacy []legacyWorktree
			if err := json.Unmarshal(raw, &legacy); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			for _, wt := range legacy {
				item := worktree{Branch: wt.Branch}
				item.Head.ShortSha = wt.Commit.ShortSha
				item.Head.Subject = wt.Commit.Message
				items = append(items, item)
			}
		}

		batch := carapace.Batch()
		for _, wt := range items {
			worktree := wt
			batch = append(batch, carapace.ActionValuesDescribed(worktree.Branch, worktree.Head.Subject).
				UidF(func(s string, uc uid.Context) (*url.URL, error) {
					return git.Uid("ref")(worktree.Head.ShortSha, uc)
				}),
			)
		}
		return batch.ToA().Style(styles.Git.Branch)
	}).Tag("worktrees")
}
