package tea

import (
	"strconv"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

// ActionPullrequestFields completes pullrequest fields
//
//	index
//	state
func ActionPullrequestFields() carapace.Action {
	return carapace.ActionValues(
		"index",
		"state",
		"author",
		"author-id",
		"url",
		"title",
		"body",
		"mergeable",
		"base",
		"base-commit",
		"head",
		"diff",
		"patch",
		"created",
		"updated",
		"deadline",
		"assignees",
		"milestone",
		"labels",
		"comments",
	)
}

type PullrequestOpts struct {
	Login  string
	Remote string
	Repo   string
	Open   bool
	Closed bool
}

// ActionPullrequests completes pull requests
func ActionPullrequests(opts PullrequestOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var pullrequests []struct {
			Index int    `yaml:"index"`
			Title string `yaml:"title"`
			State string `yaml:"state"`
		}

		args := []string{"pull", "list"}
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
		return actionYamlQuery(repoOpts, &pullrequests, args...)(func() carapace.Action {
			vals := make([]string, 0)
			for _, pr := range pullrequests {
				vals = append(vals, strconv.Itoa(pr.Index), pr.Title, styles.Tea.ForState(pr.State, c))
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}
