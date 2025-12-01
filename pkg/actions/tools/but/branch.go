package but

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
)

type branchList struct {
	AppliedStacks []struct {
		ID    string `json:"id"`
		Heads []struct {
			Name         string `json:"name"`
			Reviews      []any  `json:"reviews"`
			LastCommitAt int64  `json:"lastCommitAt"`
			CommitsAhead any    `json:"commitsAhead"`
			LastAuthor   struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"lastAuthor"`
			MergesCleanly bool `json:"mergesCleanly"`
		} `json:"heads"`
	} `json:"appliedStacks"`
	Branches []any `json:"branches"` // TODO what are these?
}

// ActionLocalBranches completes local branches
//
//	branch (branch description)
//	another (another description)
func ActionLocalBranches() carapace.Action {
	return carapace.ActionExecCommand("but", "branch", "list", "--local", "--json")(func(output []byte) carapace.Action {
		var list branchList
		if err := json.Unmarshal(output, &list); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		branches := make([]string, 0)
		for _, appliedStack := range list.AppliedStacks {
			for _, head := range appliedStack.Heads { // TODO all heads correct?
				branches = append(branches, head.Name)
			}
		}
		// but branches are git branches and lack context, so just use those from git
		return git.ActionLocalBranches().Retain(branches...)
	})
}
