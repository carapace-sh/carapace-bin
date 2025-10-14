package action

import (
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/spf13/cobra"
)

type branch struct {
	Name   string
	Target struct{ AbbreviatedOid string } // TODO last commit message?
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

func ActionBranches(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult branchesQuery
		return GraphQlAction(cmd, `repository(owner: $owner, name: $repo){ refs(first: 100, refPrefix: "refs/heads/") { nodes { name, target { abbreviatedOid } } } }`, &queryResult, func() carapace.Action {
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

func ActionBranchCommits(cmd *cobra.Command, branch string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		ref := `defaultBranchRef`
		if branch != "" {
			ref = fmt.Sprintf(`ref(qualifiedName: "refs/heads/%v")`, branch)
		}

		var queryResult branchCommitsQuery
		return GraphQlAction(cmd, fmt.Sprintf(`repository(owner: $owner, name: $repo) { %v { target { ... on Commit { history(first: 100) { edges { node { ... on Commit { abbreviatedOid message } } } } } } } }`, ref), &queryResult, func() carapace.Action {
			commits := queryResult.Data.Repository.DefaultBranchRef.Target.History.Edges
			if branch != "" {
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
