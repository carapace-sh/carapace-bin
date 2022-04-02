package action

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

type issue struct {
	Number int
	Title  string
	State  string
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
	Closed bool
	Open   bool
}

func (i *IssueOpts) states() string {
	states := make([]string, 0)
	for index, include := range []bool{i.Closed, i.Open} {
		if include {
			states = append(states, []string{"CLOSED", "OPEN"}[index])
		}
	}
	return fmt.Sprintf("[%v]", strings.Join(states, ","))
}

func ActionIssues(cmd *cobra.Command, opts IssueOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult issueQuery
		return GraphQlAction(cmd, fmt.Sprintf(`repository(owner: $owner, name: $repo){ issues(first: 100, states: %v, orderBy: {direction: DESC, field: UPDATED_AT}) { nodes { number, title, state } } }`, opts.states()), &queryResult, func() carapace.Action {
			issues := queryResult.Data.Repository.Issues.Nodes
			vals := make([]string, 0)
			for _, issue := range issues {
				s := style.Default
				switch issue.State {
				case "OPEN":
					s = style.Green
				case "CLOSED":
					s = style.Red
				case "MERGED":
					s = style.Magenta
				default:
				}

				vals = append(vals, strconv.Itoa(issue.Number), issue.Title, s)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}

func ActionIssueFields() carapace.Action {
	return carapace.ActionValues(
		"assignees",
		"author",
		"body",
		"closed",
		"comments",
		"createdAt",
		"closedAt",
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
