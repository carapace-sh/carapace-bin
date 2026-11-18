package action

import (
	"encoding/json"
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

type PullRequestOpts struct {
	Closed    bool
	Merged    bool
	Open      bool
	DraftOnly bool
}

// ActionPullRequests completes pull requests
func ActionPullRequests(cmd *cobra.Command, opts PullRequestOpts) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionPullRequests(gh.PullRequestOpts{
		Host:      repo.Host,
		Owner:     repo.Owner,
		Name:      repo.Name,
		Closed:    opts.Closed,
		Merged:    opts.Merged,
		Open:      opts.Open,
		DraftOnly: opts.DraftOnly,
	})
}

// ActionPullRequestFields completes pull request fields
func ActionPullRequestFields() carapace.Action {
	return gh.ActionPullRequestFields()
}

// ActionPullRequestCheckFields completes pull request check fields
func ActionPullRequestCheckFields() carapace.Action {
	return gh.ActionPullRequestCheckFields()
}

// ActionPullRequestLabels completes labels of a pull request
func ActionPullRequestLabels(cmd *cobra.Command, issueId string) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionPullRequestLabels(gh.PullRequestOpts{
		Host:   repo.Host,
		Owner:  repo.Owner,
		Name:   repo.Name,
		Number: issueId,
	})
}

// ActionPullRequestAssignees completes assignees of a pull request
func ActionPullRequestAssignees(cmd *cobra.Command, issueId string) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionPullRequestAssignees(gh.PullRequestOpts{
		Host:   repo.Host,
		Owner:  repo.Owner,
		Name:   repo.Name,
		Number: issueId,
	})
}

// ActionPullRequestReviewers completes reviewers of a pull request
func ActionPullRequestReviewers(cmd *cobra.Command, issueId string) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionPullRequestReviewers(gh.PullRequestOpts{
		Host:   repo.Host,
		Owner:  repo.Owner,
		Name:   repo.Name,
		Number: issueId,
	})
}

// ActionPullRequestCommits completes commits of a pull request
func ActionPullRequestCommits(cmd *cobra.Command, issueId string) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionPullRequestCommits(gh.PullRequestOpts{
		Host:   repo.Host,
		Owner:  repo.Owner,
		Name:   repo.Name,
		Number: issueId,
	})
}

// ActionPullRequestUrls completes pull request urls
func ActionPullRequestUrls(cmd *cobra.Command, opts PullRequestOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"pr", "list", "--limit", "100", "--json", "number,title,url,state,isDraft"}
		if f := cmd.Flag("repo"); f != nil && f.Changed {
			args = append(args, "--repo", f.Value.String())
		}

		states := make([]string, 0)
		if opts.Open {
			states = append(states, "open")
		}
		if opts.Closed {
			states = append(states, "closed")
		}
		if opts.Merged {
			states = append(states, "merged")
		}
		switch len(states) {
		case 0:
			return carapace.ActionValues()
		case 1:
			args = append(args, "--state", states[0])
		case 2, 3:
			args = append(args, "--state", "all")
		}

		return carapace.ActionExecCommand("gh", args...)(func(output []byte) carapace.Action {
			var pullrequests []struct {
				Number  int
				Title   string
				URL     string `json:"url"`
				State   string
				IsDraft bool `json:"isDraft"`
			}
			if err := json.Unmarshal(output, &pullrequests); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, pr := range pullrequests {
				switch {
				case !opts.Open && pr.State == "OPEN":
					continue
				case !opts.Closed && pr.State == "CLOSED":
					continue
				case !opts.Merged && pr.State == "MERGED":
					continue
				case opts.DraftOnly && !pr.IsDraft:
					continue
				}

				s := style.Default
				switch pr.State {
				case "OPEN":
					if pr.IsDraft {
						s = styles.Gh.Draft
					} else {
						s = styles.Gh.StateOpen
					}
				case "CLOSED":
					s = styles.Gh.StateClosed
				case "MERGED":
					s = styles.Gh.StateMerged
				}

				vals = append(vals, pr.URL, fmt.Sprintf("#%v %v", pr.Number, pr.Title), s)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	}).Tag("pull request urls")
}
