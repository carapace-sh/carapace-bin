package gh

import (
	"strconv"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace/pkg/style"
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

type issueType struct {
	ID          string
	Name        string
	Description string
	Color       string
}

type issueTypesQuery struct {
	Data struct {
		Repository struct {
			IssueTypes struct {
				Nodes []issueType
			}
		}
	}
}

// issueTypeColor maps GitHub IssueTypeColor enum values to carapace styles
var issueTypeColor = map[string]string{
	"BLUE":   style.Blue,
	"GRAY":   style.BrightBlack,
	"GREEN":  style.Green,
	"ORANGE": style.XTerm256Color(208),
	"PINK":   style.XTerm256Color(213),
	"PURPLE": style.Magenta,
	"RED":    style.Red,
	"YELLOW": style.Yellow,
}

// ActionIssueTypes completes issue types
//
//	Bug (Something that doesn't work)
//	Feature (A new feature)
func ActionIssueTypes(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult issueTypesQuery
		return graphQlAction(opts, `repository(owner: $owner, name: $repo){ issueTypes(first: 50) { nodes { id, name, description, color } } }`, &queryResult, func() carapace.Action {
			issueTypes := queryResult.Data.Repository.IssueTypes.Nodes
			vals := make([]string, 0)
			for _, t := range issueTypes {
				s, ok := issueTypeColor[t.Color]
				if !ok {
					s = style.Default
				}
				vals = append(vals, t.Name, t.Description, s)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		}).Cache(24*time.Hour, opts.cacheKey())
	}).Tag("issue types")
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
