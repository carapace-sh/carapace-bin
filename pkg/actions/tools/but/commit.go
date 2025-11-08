package but

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

type logEntry struct {
	BranchDetails []struct {
		Commits []struct {
			ID           string `json:"id"`
			Message      string `json:"message"`
			HasConflicts bool   `json:"hasConflicts"`
			State        struct {
				Type string `json:"type"`
			} `json:"state"`
			CreatedAt int64 `json:"createdAt"`
		} `json:"commits"`
	} `json:"branchDetails"`
}

// ActionCommits completes commits
//
//	36ae34b (some commit)
//	e1b2490 (another commit)
func ActionCommits() carapace.Action {
	return carapace.ActionExecCommand("but", "--json", "log")(func(output []byte) carapace.Action {
		var entries []logEntry
		if err := json.Unmarshal(output, &entries); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, entry := range entries {
			for _, branch := range entry.BranchDetails {
				for _, commit := range branch.Commits {
					vals = append(vals, commit.ID[:7], commit.Message, styles.Git.Commit)
				}
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Tag("commits").UidF(git.Uid("ref")) // TODO custom uid
}
