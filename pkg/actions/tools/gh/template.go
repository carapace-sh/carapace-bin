package gh

import "github.com/rsteube/carapace"

// ActionIssueTemplates completes issue templates
//
//	⭐ Submit a request (Surface a feature or problem that you think should be solved)
//	🐛 Bug report (Report a bug or unexpected behavior while using GitHub CLI)
func ActionIssueTemplates(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult struct {
			Data struct {
				Repository struct {
					IssueTemplates []struct {
						About string
						Name  string
					}
				}
			}
		}
		return graphQlAction(opts, `repository(owner: $owner, name: $repo){ issueTemplates { about, name } }`, &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, t := range queryResult.Data.Repository.IssueTemplates {
				vals = append(vals, t.Name, t.About)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

// ActionPullRequestTemplates completes pull request templates
//
//	PULL_REQUEST_TEMPLATE.md (content)
func ActionPullRequestTemplates(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult struct {
			Data struct {
				Repository struct {
					PullRequestTemplates []struct {
						Body     string
						Filename string
					}
				}
			}
		}
		return graphQlAction(opts, `repository(owner: $owner, name: $repo){ pullRequestTemplates { body, filename } }`, &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, t := range queryResult.Data.Repository.PullRequestTemplates {
				vals = append(vals, t.Filename, t.Body)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
