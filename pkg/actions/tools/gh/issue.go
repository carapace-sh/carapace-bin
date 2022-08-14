package gh

import (
	"strconv"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
	"github.com/rsteube/carapace/pkg/style"
)

type issue struct {
	Number    int
	Title     string
	State     string
	Labels    []label
	Assignees []struct {
		Login string
	}
}

type pinnedIssueQuery struct {
	Data struct {
		Repository struct {
			PinnedIssues struct {
				Nodes []struct {
					Issue issue
				}
			}
		}
	}
}

// ActionPinnedIssues completes pinned issues
//
//	11 (issue description)
//	12 (issue description)
func ActionPinnedIssues(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult pinnedIssueQuery
		return graphQlAction(opts, `repository(owner: $owner, name: $repo){ pinnedIssues(first: 3) { nodes { issue { number, title, state } } } }`, &queryResult, func() carapace.Action {
			issues := queryResult.Data.Repository.PinnedIssues.Nodes
			vals := make([]string, 0)
			for _, issue := range issues {
				s := style.Default
				switch issue.Issue.State {
				case "OPEN":
					s = styles.Gh.StateOpen
				case "CLOSED":
					s = styles.Gh.StateClosed
				case "MERGED":
					s = styles.Gh.StateMerged
				default:
				}

				vals = append(vals, strconv.Itoa(issue.Issue.Number), issue.Issue.Title, s)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}
