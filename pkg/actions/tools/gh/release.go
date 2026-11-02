package gh

import (
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/number"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace-bin/pkg/util"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionReleaseFields completes release fields
//
//	name
//	tagName
func ActionReleaseFields() carapace.Action {
	return carapace.ActionValues(
		"createdAt",
		"isDraft",
		"isLatest",
		"isPrerelease",
		"name",
		"publishedAt",
		"tagName",
	)
}

type release struct {
	Name         string
	IsPrerelease bool
	CreatedAt    time.Time
	Tag          struct {
		Name string
	}
	ReleaseAssets struct {
		Nodes []struct {
			Name string
			Size int
		}
	}
}

type releaseQuery struct {
	Data struct {
		Repository struct {
			Releases struct {
				Edges []struct {
					Node release
				}
			}
		}
	}
}

type ReleaseOpts struct {
	Host  string
	Owner string
	Name  string

	Tag string
}

func (o ReleaseOpts) Default() ReleaseOpts {
	return o
}

func (o ReleaseOpts) repo() RepoOpts {
	return RepoOpts{Host: o.Host, Owner: o.Owner, Name: o.Name}
}

// ActionReleases completes releases
//
//	v1.0.0 (a day ago)
//	v0.9.0 (a month ago)
func ActionReleases(opts ReleaseOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult releaseQuery
		return graphQlAction(opts.repo(), `repository(owner: $owner, name: $repo) { releases(first: 100, orderBy: {direction: DESC, field: CREATED_AT}) { edges { node { createdAt isPrerelease name tag { name } } } } }`, &queryResult, func() carapace.Action {
			releases := queryResult.Data.Repository.Releases.Edges
			vals := make([]string, 0, len(releases)*2)
			for _, release := range releases {
				if release.Node.Tag.Name != "" {
					vals = append(vals, release.Node.Tag.Name, util.FuzzyAgo(time.Since(release.Node.CreatedAt)))
				}
			}
			return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Tag)
		})
	}).Tag("releases")
}

// ActionNextReleases completes next release versions
//
//	v1.0.1
//	v1.1.0
func ActionNextReleases(opts ReleaseOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult releaseQuery
		return graphQlAction(opts.repo(), `repository(owner: $owner, name: $repo) { releases(first: 100, orderBy: {direction: DESC, field: CREATED_AT}) { edges { node { createdAt isPrerelease name tag { name } } } } }`, &queryResult, func() carapace.Action {
			releases := queryResult.Data.Repository.Releases.Edges
			vals := make([]string, 0, len(releases))
			for _, release := range releases {
				if release.Node.Tag.Name != "" {
					vals = append(vals, release.Node.Tag.Name)
				}
			}
			return number.ActionSemanticVersions(vals...).Style(styles.Git.Tag)
		})
	}).Tag("release versions")
}

// ActionReleaseAssets completes assets of a recent release
func ActionReleaseAssets(opts ReleaseOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult releaseQuery
		return graphQlAction(opts.repo(), `repository(owner: $owner, name: $repo) { releases(first: 100, orderBy: {direction: DESC, field: CREATED_AT}) { edges { node { createdAt isPrerelease name tag { name } releaseAssets(first: 100) { nodes { name size } } } } } }`, &queryResult, func() carapace.Action {
			releases := queryResult.Data.Repository.Releases.Edges
			vals := make([]string, 0)
			for _, release := range releases {
				if release.Node.Tag.Name == opts.Tag {
					for _, asset := range release.Node.ReleaseAssets.Nodes {
						vals = append(vals, asset.Name, byteCountSI(asset.Size))
					}
					break
				}
			}
			return carapace.ActionValuesDescribed(vals...).StyleF(style.ForPathExt)
		})
	}).Tag("release assets")
}
