package but

import (
	"encoding/json"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type butChange struct {
	CliID      string `json:"cliId"`
	FilePath   string `json:"filePath"`
	ChangeType string `json:"changeType"`
}

func (c butChange) Description() string {
	m := map[string]string{
		"added":    "A",
		"removed":  "D",
		"modified": "M",
		"renamed":  "R",
	}
	return m[c.ChangeType] + " " + c.FilePath
}

type butCommit struct {
	CliID       string      `json:"cliId"`
	CommitID    string      `json:"commitId"`
	CreatedAt   time.Time   `json:"createdAt"`
	Message     string      `json:"message"`
	AuthorName  string      `json:"authorName"`
	AuthorEmail string      `json:"authorEmail"`
	Conflicted  bool        `json:"conflicted"`
	ReviewID    any         `json:"reviewId"`
	Changes     []butChange `json:"changes"`
}

type butBranch struct {
	CliID           string      `json:"cliId"`
	Name            string      `json:"name"`
	Commits         []butCommit `json:"commits"`
	UpstreamCommits []any       `json:"upstreamCommits"`
	BranchStatus    string      `json:"branchStatus"`
	ReviewID        any         `json:"reviewId"`
}

type butStack struct {
	CliID           string      `json:"cliId"`
	AssignedChanges []butChange `json:"assignedChanges"`
	Branches        []butBranch `json:"branches"`
}

type butStatus struct {
	UnassignedChanges []butChange `json:"unassignedChanges"`
	Stacks            []butStack  `json:"stacks"`
	MergeBase         butCommit   `json:"mergeBase"`
	UpstreamState     struct {
		Behind       int       `json:"behind"`
		LatestCommit butCommit `json:"latestCommit"`
		LastFetched  any       `json:"lastFetched"`
	} `json:"upstreamState"`
}

func actionStatus(includeCommitted bool, f func(status butStatus) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"status", "--json"}
		if includeCommitted {
			args = append(args, "-f")
		}
		return carapace.ActionExecCommand("but", args...)(func(output []byte) carapace.Action {
			var status butStatus
			if err := json.Unmarshal(output, &status); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return f(status)
		})
	})
}

func styleForChangeType(s string) string {
	switch s {
	case "added": /// The file was newly added (it was not tracked before)
		return style.Carapace.KeywordPositive
	case "removed": /// The file was deleted
		return style.Carapace.KeywordNegative
	case "modified": /// The file was modified
		return style.Carapace.KeywordAmbiguous
	case "renamed": /// The file was renamed
		return style.Dim
	default:
		return style.Default
	}
}
