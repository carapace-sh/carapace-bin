package gh

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace/pkg/cache/key"
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

type issueQuery struct {
	Data struct {
		Repository struct {
			Issues struct {
				Nodes []issue
			}
		}
	}
}

type IssueOpts struct {
	Host  string
	Owner string
	Name  string

	Closed bool
	Open   bool

	Number string
}

func (o IssueOpts) Default() IssueOpts {
	o.Closed = true
	o.Open = true
	return o
}

func (o IssueOpts) repo() RepoOpts {
	return RepoOpts{Host: o.Host, Owner: o.Owner, Name: o.Name}
}

func (o IssueOpts) cacheKey() key.Key {
	return key.String(o.Host, o.Owner, o.Name, strconv.FormatBool(o.Closed), strconv.FormatBool(o.Open))
}

func (o IssueOpts) states() string {
	states := make([]string, 0)
	for index, include := range []bool{o.Closed, o.Open} {
		if include {
			states = append(states, []string{"CLOSED", "OPEN"}[index])
		}
	}
	return fmt.Sprintf("[%v]", strings.Join(states, ","))
}

// ActionIssues completes issues
//
//	42 (Fix login bug)
//	43 (Add feature)
func ActionIssues(opts IssueOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult issueQuery
		return graphQlAction(opts.repo(), fmt.Sprintf(`repository(owner: $owner, name: $repo){ issues(first: 100, states: %v, orderBy: {direction: DESC, field: UPDATED_AT}) { nodes { number, title, state } } }`, opts.states()), &queryResult, func() carapace.Action {
			issues := queryResult.Data.Repository.Issues.Nodes
			vals := make([]string, 0)
			for _, issue := range issues {
				s := style.Default
				switch issue.State {
				case "OPEN":
					s = styles.Gh.StateOpen
				case "CLOSED":
					s = styles.Gh.StateClosed
				case "MERGED":
					s = styles.Gh.StateMerged
				default:
				}

				vals = append(vals, strconv.Itoa(issue.Number), issue.Title, s)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		}).Cache(24*time.Hour, opts.cacheKey())
	}).Tag("issues")
}

// ActionIssueFields completes issue fields
//
//	author
//	labels
func ActionIssueFields() carapace.Action {
	return carapace.ActionValues(
		"assignees",
		"author",
		"body",
		"closed",
		"closedAt",
		"comments",
		"createdAt",
		"id",
		"labels",
		"milestone",
		"number",
		"projectCards",
		"reactionGroups",
		"state",
		"title",
		"updatedAt",
		"url",
	)
}

// ActionIssueLabels completes labels of an issue
//
//	bug (Something that doesn't work)
func ActionIssueLabels(opts IssueOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return resolveRepo(c, opts.repo(), func(resolved RepoOpts) carapace.Action {
			var queryResult issue
			return apiV3Action(resolved, fmt.Sprintf("repos/%v/%v/issues/%v", resolved.Owner, resolved.Name, opts.Number), &queryResult, func() carapace.Action {
				vals := make([]string, 0)
				for _, label := range queryResult.Labels {
					vals = append(vals, label.Name, label.Description, "#"+label.Color)
				}
				return carapace.ActionStyledValuesDescribed(vals...)
			})
		})
	})
}

// ActionIssueAssignees completes assignees of an issue
//
//	userA
//	userB
func ActionIssueAssignees(opts IssueOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return resolveRepo(c, opts.repo(), func(resolved RepoOpts) carapace.Action {
			var queryResult issue
			return apiV3Action(resolved, fmt.Sprintf("repos/%v/%v/issues/%v", resolved.Owner, resolved.Name, opts.Number), &queryResult, func() carapace.Action {
				vals := make([]string, 0)
				for _, assignee := range queryResult.Assignees {
					vals = append(vals, assignee.Login)
				}
				return carapace.ActionValues(vals...)
			})
		})
	})
}
