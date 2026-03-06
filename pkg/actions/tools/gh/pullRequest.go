package gh

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace/pkg/style"
)

type pullrequest struct {
	Number    int
	Title     string
	IsDraft   bool
	State     string
	Labels    []label
	Assignees []struct {
		Login string
	}
	RequestedReviewers []struct {
		Login string
	} `json:"requested_reviewers"`
}

type pullRequestsQuery struct {
	Data struct {
		Repository struct {
			PullRequests struct {
				Nodes []pullrequest
			}
		}
	}
}

type PullRequestOpts struct {
	Host  string
	Owner string
	Name  string

	Closed    bool
	Merged    bool
	Open      bool
	DraftOnly bool
}

func (o PullRequestOpts) Default() PullRequestOpts {
	o.Closed = true
	o.Merged = true
	o.Open = true
	o.DraftOnly = true
	return o
}

func (o PullRequestOpts) repo() RepoOpts {
	return RepoOpts{Host: o.Host, Owner: o.Owner, Name: o.Name}
}

func (o *PullRequestOpts) states() string {
	states := make([]string, 0)
	for index, include := range []bool{o.Closed, o.Merged, o.Open} {
		if include {
			states = append(states, []string{"CLOSED", "MERGED", "OPEN"}[index])
		}
	}
	return fmt.Sprintf("[%v]", strings.Join(states, ","))
}

// ActionPullRequests completes pull requests
// TODO quickly moved over from gh_completer - needs polishing (and removal of the old one)
func ActionPullRequests(opts PullRequestOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult pullRequestsQuery
		return graphQlAction(opts.repo(), fmt.Sprintf(`repository(owner: $owner, name: $repo){ pullRequests(first: 100, states: %v, orderBy: {direction: DESC, field: UPDATED_AT}) { nodes { number, title, isDraft, state } } }`, opts.states()), &queryResult, func() carapace.Action {
			pullrequests := queryResult.Data.Repository.PullRequests.Nodes
			vals := make([]string, 0, len(pullrequests)*2)
			for _, pullrequest := range pullrequests {
				if !opts.DraftOnly || pullrequest.IsDraft {

					s := style.Default
					switch pullrequest.State {
					case "OPEN":
						if pullrequest.IsDraft {
							s = styles.Gh.Draft
						} else {
							s = styles.Gh.StateOpen
						}
					case "CLOSED":
						s = styles.Gh.StateClosed
					case "MERGED":
						s = styles.Gh.StateMerged
					default:
					}

					vals = append(vals, strconv.Itoa(pullrequest.Number), pullrequest.Title, s)
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	}).Tag("pull requests")
}
