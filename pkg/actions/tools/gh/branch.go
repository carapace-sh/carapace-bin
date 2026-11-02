package gh

import (
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

type branch struct {
	Name   string
	Target struct {
		AbbreviatedOid string
	}
}

type branchesQuery struct {
	Data struct {
		Repository struct {
			Refs struct {
				Nodes []branch
			}
		}
	}
}

// ActionBranches completes branches
//
//	main (cbc683c)
//	develop (cbc683c)
func ActionBranches(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult branchesQuery
		return graphQlAction(opts, `repository(owner: $owner, name: $repo){ refs(first: 100, refPrefix: "refs/heads/") { nodes { name, target { abbreviatedOid } } } }`, &queryResult, func() carapace.Action {
			branches := queryResult.Data.Repository.Refs.Nodes
			vals := make([]string, len(branches)*2)
			for index, branch := range branches {
				vals[index*2] = branch.Name
				vals[index*2+1] = branch.Target.AbbreviatedOid
			}
			return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Branch)
		})
	}).Tag("branches")
}

type commitRef struct {
	Target struct {
		History struct {
			Edges []struct {
				Node struct {
					AbbreviatedOid string
					Message        string
				}
			}
		}
	}
}

type branchCommitsQuery struct {
	Data struct {
		Repository struct {
			DefaultBranchRef commitRef
			Ref              commitRef
		}
	}
}

type BranchOpts struct {
	Host  string
	Owner string
	Name  string

	Branch string
}

func (o BranchOpts) Default() BranchOpts {
	return o
}

func (o BranchOpts) repo() RepoOpts {
	return RepoOpts{Host: o.Host, Owner: o.Owner, Name: o.Name}
}

// ActionBranchCommits completes commits of a branch
//
//	cbc683c (fix release notes generation)
func ActionBranchCommits(opts BranchOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		ref := `defaultBranchRef`
		if opts.Branch != "" {
			ref = fmt.Sprintf(`ref(qualifiedName: "refs/heads/%v")`, opts.Branch)
		}

		var queryResult branchCommitsQuery
		return graphQlAction(opts.repo(), fmt.Sprintf(`repository(owner: $owner, name: $repo) { %v { target { ... on Commit { history(first: 100) { edges { node { ... on Commit { abbreviatedOid message } } } } } } } }`, ref), &queryResult, func() carapace.Action {
			commits := queryResult.Data.Repository.DefaultBranchRef.Target.History.Edges
			if opts.Branch != "" {
				commits = queryResult.Data.Repository.Ref.Target.History.Edges
			}

			vals := make([]string, len(commits)*2)
			for index, commit := range commits {
				vals[index*2] = commit.Node.AbbreviatedOid
				vals[index*2+1] = commit.Node.Message
			}
			return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Branch)
		})
	}).Tag("branch commits")
}

// ActionSearchRepositoryFields completes repository search fields variant
//
//	stargazersCount
//	visibility
func ActionSearchRepositoryFields() carapace.Action {
	return ActionRepositorySearchFields()
}
