package but

import (
	"encoding/json"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

type butCommit struct {
	CliID       string    `json:"cliId"`
	CommitID    string    `json:"commitId"`
	CreatedAt   time.Time `json:"createdAt"`
	Message     string    `json:"message"`
	AuthorName  string    `json:"authorName"`
	AuthorEmail string    `json:"authorEmail"`
	Conflicted  bool      `json:"conflicted"`
	ReviewID    any       `json:"reviewId"`
	Changes     any       `json:"changes"`
}

type butStatus struct {
	Stacks []struct {
		Branches []struct {
			Commits []butCommit `json:"commits"`
		} `json:"branches"`
	} `json:"stacks"`
	MergeBase     butCommit `json:"mergeBase"`
	UpstreamState struct {
		Behind       int       `json:"behind"`
		LatestCommit butCommit `json:"latestCommit"`
		LastFetched  any       `json:"lastFetched"`
	} `json:"upstreamState"`
}

// ActionCommits completes commits
//
//	36ae34b (some commit)
//	e1b2490 (another commit)
func ActionCommits() carapace.Action {
	return carapace.ActionExecCommand("but", "status", "--json")(func(output []byte) carapace.Action {
		var status butStatus
		if err := json.Unmarshal(output, &status); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, stack := range status.Stacks {
			for _, branch := range stack.Branches {
				for _, commit := range branch.Commits {
					vals = append(vals, commit.CommitID[:7], commit.Message, styles.Git.Commit)
				}
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Tag("commits").UidF(git.Uid("ref")) // TODO custom uid
}
