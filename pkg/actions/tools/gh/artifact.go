package gh

import (
	"fmt"
	"strconv"

	"github.com/carapace-sh/carapace"
)

type artifact struct {
	Id          int
	Name        string
	SizeInBytes int `json:"size_in_bytes"`
}

type artifactsQuery struct {
	Artifacts []artifact
}

type ArtifactOpts struct {
	Host  string
	Owner string
	Name  string

	RunId string
}

func (o ArtifactOpts) Default() ArtifactOpts {
	return o
}

func (o ArtifactOpts) repo() RepoOpts {
	return RepoOpts{Host: o.Host, Owner: o.Owner, Name: o.Name}
}

func (o ArtifactOpts) query(owner, name string) string {
	if o.RunId != "" {
		return fmt.Sprintf("repos/%v/%v/actions/runs/%v/artifacts", owner, name, o.RunId)
	}
	return fmt.Sprintf("repos/%v/%v/actions/artifacts", owner, name)
}

// ActionWorkflowArtifactNames completes artifact names
//
//	binary (1.6 MB)
func ActionWorkflowArtifactNames(opts ArtifactOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return resolveRepo(c, opts.repo(), func(resolved RepoOpts) carapace.Action {
			var queryResult artifactsQuery
			return apiV3Action(resolved, opts.query(resolved.Owner, resolved.Name), &queryResult, func() carapace.Action {
				vals := make([]string, 0)
				for _, artifact := range queryResult.Artifacts {
					vals = append(vals, artifact.Name, byteCountSI(artifact.SizeInBytes))
				}
				return carapace.ActionValuesDescribed(vals...)
			})
		})
	}).Tag("artifact names")
}

// ActionWorkflowArtifactIds completes artifact ids
//
//	2036111109 (binary (1.6 MB))
func ActionWorkflowArtifactIds(opts ArtifactOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return resolveRepo(c, opts.repo(), func(resolved RepoOpts) carapace.Action {
			var queryResult artifactsQuery
			return apiV3Action(resolved, opts.query(resolved.Owner, resolved.Name), &queryResult, func() carapace.Action {
				vals := make([]string, 0)
				for _, artifact := range queryResult.Artifacts {
					vals = append(vals, strconv.Itoa(artifact.Id), fmt.Sprintf("%v (%v)", artifact.Name, byteCountSI(artifact.SizeInBytes)))
				}
				return carapace.ActionValuesDescribed(vals...)
			})
		})
	}).Tag("artifact ids")
}

// https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/
func byteCountSI(b int) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}
