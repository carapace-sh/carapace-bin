package gh

import (
	"fmt"
	"strconv"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type project struct {
	Number int
	Title  string
	Closed bool
	Items  struct {
		Nodes []struct {
			Id         string
			Type       string
			IsArchived bool
			Content    struct {
				Title string
			}
		}
	}
}

type projectQuery struct {
	Data struct {
		User struct {
			ProjectV2 project
		}
		Organization struct {
			ProjectV2 project
		}
	}
}

type projectsQuery struct {
	Data struct {
		User struct {
			ProjectsV2 struct {
				Nodes []project
			}
		}
		Organization struct {
			ProjectsV2 struct {
				Nodes []project
			}
		}
	}
}

type ProjectOpts struct {
	Host   string
	Owner  string
	Open   bool
	Closed bool
}

func (o ProjectOpts) Default() ProjectOpts {
	o.Owner = "@me"
	o.Open = true
	o.Closed = true
	return o
}

func (o ProjectOpts) repo() RepoOpts {
	return RepoOpts{
		Host:  o.Host,
		Owner: o.Owner,
	}
}

type ProjectItemOpts struct {
	Host     string
	Owner    string
	Project  int
	Archived bool
}

func (o ProjectItemOpts) repo() RepoOpts {
	return RepoOpts{
		Host:  o.Host,
		Owner: o.Owner,
	}
}

// ActionProjects completes projects
//
//	1 (first project)
//	2 (second project)
func ActionProjects(opts ProjectOpts) carapace.Action {
	return carapace.Batch(
		actionUserProjects(opts),
		actionOrganizationProjects(opts),
	).ToA().Suppress("Could not resolve to")
}

func actionUserProjects(opts ProjectOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult projectsQuery
		return graphQlAction(opts.repo(), `user(login: $owner) { projectsV2(first: 100) { nodes { number title closed } } }`, &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, project := range queryResult.Data.User.ProjectsV2.Nodes {
				vals = append(vals, strconv.Itoa(project.Number), project.Title, style.ForKeyword(strconv.FormatBool(!project.Closed), c))
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}

func actionOrganizationProjects(opts ProjectOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult projectsQuery
		return graphQlAction(opts.repo(), `organization(login: $owner) { projectsV2(first: 100) { nodes { number title closed } } }`, &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, project := range queryResult.Data.Organization.ProjectsV2.Nodes {
				vals = append(vals, strconv.Itoa(project.Number), project.Title, style.ForKeyword(strconv.FormatBool(!project.Closed), c))
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}

// ActionProjectItems completes project items
//
//	PVTI_lADOA48Fh84ABd_DzgBCG7c (Issue commands do not work with non-classic Projects)
//	PVTI_lADOA48Fh84ABd_DzgBCHAo (Checkout branch for issue)
func ActionProjectItems(opts ProjectItemOpts) carapace.Action {
	return carapace.Batch(
		actionUserProjectItems(opts),
		actionOrganizationProjectItems(opts),
	).ToA().Suppress("Could not resolve to")
}

func actionUserProjectItems(opts ProjectItemOpts) carapace.Action {
	var queryResult projectQuery
	// TODO filter archived
	return graphQlAction(opts.repo(), fmt.Sprintf(`user(login: $owner) { projectV2(number: %v) { items(first: 100) { nodes { id type isArchived content { ... on DraftIssue { title }  ... on Issue { title } ... on PullRequest { title } } } } } }`, opts.Project),
		&queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, item := range queryResult.Data.User.ProjectV2.Items.Nodes {
				vals = append(vals, item.Id, item.Content.Title) // TODO style for type/state
			}
			return carapace.ActionValuesDescribed(vals...)
		})
}

func actionOrganizationProjectItems(opts ProjectItemOpts) carapace.Action {
	var queryResult projectQuery
	// TODO filter archived
	return graphQlAction(opts.repo(), fmt.Sprintf(`organization(login: $owner) { projectV2(number: %v) { items(first: 100) { nodes { id type isArchived content { ... on DraftIssue { title }  ... on Issue { title } ... on PullRequest { title } } } } } }`, opts.Project),
		&queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, item := range queryResult.Data.Organization.ProjectV2.Items.Nodes {
				vals = append(vals, item.Id, item.Content.Title) // TODO style for type/state
			}
			return carapace.ActionValuesDescribed(vals...)
		})
}
