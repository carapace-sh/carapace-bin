package gh

import (
	"fmt"
	"strconv"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type project struct {
	Id     string
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

// ActionProjectIds completes project IDs
//
//	PVT_kwDOA... (first project)
//	PVT_kwDOA... (second project)
func ActionProjectIds(opts ProjectOpts) carapace.Action {
	return carapace.Batch(
		actionUserProjectIds(opts),
		actionOrganizationProjectIds(opts),
	).ToA().Suppress("Could not resolve to")
}

func actionUserProjectIds(opts ProjectOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult projectsQuery
		return graphQlAction(opts.repo(), `user(login: $owner) { projectsV2(first: 100) { nodes { id number title closed } } }`, &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, project := range queryResult.Data.User.ProjectsV2.Nodes {
				vals = append(vals, project.Id, project.Title)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func actionOrganizationProjectIds(opts ProjectOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult projectsQuery
		return graphQlAction(opts.repo(), `organization(login: $owner) { projectsV2(first: 100) { nodes { id number title closed } } }`, &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, project := range queryResult.Data.Organization.ProjectsV2.Nodes {
				vals = append(vals, project.Id, project.Title)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func actionUserProjects(opts ProjectOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult projectsQuery
		return graphQlAction(opts.repo(), `user(login: $owner) { projectsV2(first: 100) { nodes { id number title closed } } }`, &queryResult, func() carapace.Action {
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
		return graphQlAction(opts.repo(), `organization(login: $owner) { projectsV2(first: 100) { nodes { id number title closed } } }`, &queryResult, func() carapace.Action {
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
//
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

type ProjectItemNodeOpts struct {
	Host      string
	ProjectId string
}

func (o ProjectItemNodeOpts) Default() ProjectItemNodeOpts {
	return o
}

func (o ProjectItemNodeOpts) repo() RepoOpts {
	return RepoOpts{
		Host: o.Host,
	}
}

type projectNodeQuery struct {
	Data struct {
		Node struct {
			Items struct {
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
	}
}

// ActionProjectNodeItems completes project items by project node ID
//
//	PVTI_lADOA48Fh84ABd_DzgBCG7c (Issue commands do not work with non-classic Projects)
//	PVTI_lADOA48Fh84ABd_DzgBCHAo (Checkout branch for issue)
func ActionProjectNodeItems(opts ProjectItemNodeOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult projectNodeQuery
		return graphQlAction(opts.repo(), fmt.Sprintf(`node(id: "%v") { ... on ProjectV2 { items(first: 100) { nodes { id type isArchived content { ... on DraftIssue { title } ... on Issue { title } ... on PullRequest { title } } } } } }`, opts.ProjectId),
			&queryResult, func() carapace.Action {
				vals := make([]string, 0)
				for _, item := range queryResult.Data.Node.Items.Nodes {
					vals = append(vals, item.Id, item.Content.Title)
				}
				return carapace.ActionValuesDescribed(vals...)
			})
	})
}

type ProjectFieldOpts struct {
	Host    string
	Owner   string
	Project int
}

func (o ProjectFieldOpts) Default() ProjectFieldOpts {
	o.Owner = "@me"
	return o
}

func (o ProjectFieldOpts) repo() RepoOpts {
	return RepoOpts{
		Host:  o.Host,
		Owner: o.Owner,
	}
}
type projectFieldNode struct {
	Typename string `json:"__typename"`
	Id       string
	Name     string
}

type projectFieldsQuery struct {
	Data struct {
		User struct {
			ProjectV2 struct {
				Fields struct {
					Nodes []projectFieldNode
				}
			}
		}
		Organization struct {
			ProjectV2 struct {
				Fields struct {
					Nodes []projectFieldNode
				}
			}
		}
	}
}

// ActionProjectFields completes project field IDs
//
//	PVTF_lADOA48Fh84ABd_DzgBCG7c (Status)
//	PVTF_lADOA48Fh84ABd_DzgBCHAo (Priority)
func ActionProjectFields(opts ProjectFieldOpts) carapace.Action {
	return carapace.Batch(
		actionUserProjectFields(opts),
		actionOrganizationProjectFields(opts),
	).ToA().Suppress("Could not resolve to")
}

func actionUserProjectFields(opts ProjectFieldOpts) carapace.Action {
	var queryResult projectFieldsQuery
	return graphQlAction(opts.repo(), fmt.Sprintf(`user(login: $owner) { projectV2(number: %v) { fields(first: 100) { nodes { __typename ... on ProjectV2Field { id name } ... on ProjectV2SingleSelectField { id name } ... on ProjectV2IterationField { id name } } } } }`, opts.Project),
		&queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, field := range queryResult.Data.User.ProjectV2.Fields.Nodes {
				vals = append(vals, field.Id, field.Name)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
}

func actionOrganizationProjectFields(opts ProjectFieldOpts) carapace.Action {
	var queryResult projectFieldsQuery
	return graphQlAction(opts.repo(), fmt.Sprintf(`organization(login: $owner) { projectV2(number: %v) { fields(first: 100) { nodes { __typename ... on ProjectV2Field { id name } ... on ProjectV2SingleSelectField { id name } ... on ProjectV2IterationField { id name } } } } }`, opts.Project),
		&queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, field := range queryResult.Data.Organization.ProjectV2.Fields.Nodes {
				vals = append(vals, field.Id, field.Name)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
}
