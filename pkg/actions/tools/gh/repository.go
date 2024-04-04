package gh

import (
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

type repository struct {
	Name        string
	Description string
	IsArchived  bool
	IsPrivate   bool
	IsFork      bool
	IsMirror    bool
	IsTemplate  bool
	IsLocked    bool
}

type repositoryQuery struct {
	Data struct {
		Search struct {
			Repos []struct {
				Repo repository
			}
		}
	}
}

// ActionRepositories completes repositories
//
//	carapace (command argument completion generator for spf13/cobra)
//	carapace-bin (multi-shell multi-command argument completer)
func ActionRepositories(opts OwnerOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult repositoryQuery
		return graphQlAction(opts.repo(), fmt.Sprintf(`search( type:REPOSITORY, query: """ user:%v "%v" in:name fork:true""", first: 100) { repos: edges { repo: node { ... on Repository { name description isArchived isPrivate isFork isMirror isTemplate isLocked } } } }`, opts.Owner, c.Value), &queryResult, func() carapace.Action {
			repositories := queryResult.Data.Search.Repos
			vals := make([]string, 0)
			for _, repo := range repositories {
				if r := repo.Repo; r.IsPrivate {
					vals = append(vals, repo.Repo.Name, repo.Repo.Description, styles.Gh.RepoPrivate)
				} else if r.IsArchived {
					vals = append(vals, repo.Repo.Name, repo.Repo.Description, styles.Gh.RepoArchived)
				} else if r.IsFork {
					vals = append(vals, repo.Repo.Name, repo.Repo.Description, styles.Gh.RepoFork)
				} else if r.IsMirror {
					vals = append(vals, repo.Repo.Name, repo.Repo.Description, styles.Gh.RepoMirror)
				} else if r.IsLocked {
					vals = append(vals, repo.Repo.Name, repo.Repo.Description, styles.Gh.RepoLocked)
				} else if r.IsTemplate {
					vals = append(vals, repo.Repo.Name, repo.Repo.Description, styles.Gh.RepoTemplate)
				} else {
					vals = append(vals, repo.Repo.Name, repo.Repo.Description, styles.Gh.RepoPublic)
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})

	})
}

// ActionOwnerRepositories completes owner/repository
//
//	carapace-sh/carapace
//	spf13/cobra
func ActionOwnerRepositories(opts HostOpts) carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionOwners(opts).Invoke(c).Suffix("/").ToA()
		case 1:
			return ActionRepositories(OwnerOpts{Owner: c.Parts[0]})
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionHostOwnerRepositories completes [host/]owner/repository
//
//	carapace-sh/carapace
//	github.com/carapace-sh/carapace
func ActionHostOwnerRepositories() carapace.Action {
	return carapace.ActionMultiPartsN("/", 3, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.Batch(
				ActionConfigHosts(),
				ActionOwners(HostOpts{}),
			).ToA().Suffix("/")
		case 1:
			if strings.Contains(c.Parts[0], ".") {
				return ActionOwners(HostOpts{Host: c.Parts[0]}).Suffix("/")
			}
			return ActionRepositories(OwnerOpts{Owner: c.Parts[0]})
		default:
			if strings.Contains(c.Parts[0], ".") {
				return ActionRepositories(OwnerOpts{Host: c.Parts[0], Owner: c.Parts[1]})
			}
			return carapace.ActionValues()
		}
	})
}
