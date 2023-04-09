package tea

import (
	"strconv"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
)

// ActionIssueFields completes issue fields
//
//	index
//	state
func ActionIssueFields() carapace.Action {
	return carapace.ActionValues(
		"index",
		"state",
		"kind",
		"author",
		"author-id",
		"url",
		"title",
		"body",
		"created",
		"updated",
		"deadline",
		"assignees",
		"milestone",
		"labels",
		"comments",
	)
}

type IssueOpts struct {
	Login  string
	Remote string
	Repo   string
	Open   bool
	Closed bool
}

// ActionIssues completes issues
func ActionIssues(opts IssueOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var issues []struct {
			Index int    `yaml:"index"`
			Title string `yaml:"title"`
			State string `yaml:"state"`
		}

		args := []string{"issue", "list"}
		switch {
		case opts.Closed && opts.Open:
			args = append(args, "--state", "all")
		case opts.Closed:
			args = append(args, "--state", "closed")
		case opts.Open:
			args = append(args, "--state", "open")
		default:
			return carapace.ActionValues()
		}

		repoOpts := RepoOpts{Login: opts.Login, Remote: opts.Remote, Repo: opts.Repo}
		return actionYamlQuery(repoOpts, &issues, args...)(func() carapace.Action {
			vals := make([]string, 0)
			for _, issue := range issues {
				vals = append(vals, strconv.Itoa(issue.Index), issue.Title, styles.Tea.ForState(issue.State, c))
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}
