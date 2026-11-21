package gh

import (
	"time"

	"github.com/carapace-sh/carapace"
)

type label struct {
	Name        string
	Description string
	Color       string
}

type labelsQuery struct {
	Data struct {
		Repository struct {
			Labels struct {
				Nodes []label
			}
		}
	}
}

// ActionLabels completes labels
//
//	enhancement (New feature or request)
//	good first issue (Good for newcomers)
func ActionLabels(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult labelsQuery
		return graphQlAction(opts, `repository(owner: $owner, name: $repo){ labels(first: 100) { nodes { name, description, color } } }`, &queryResult, func() carapace.Action {
			labels := queryResult.Data.Repository.Labels.Nodes
			vals := make([]string, 0)
			for _, label := range labels {
				vals = append(vals, label.Name, label.Description, "#"+label.Color)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	}).Tag("labels").Cache(24*time.Hour, opts.cacheKey())
}

// ActionLabelFields completes label fields
//
//	color
//	createdAt
func ActionLabelFields() carapace.Action {
	return carapace.ActionValues(
		"color",
		"createdAt",
		"description",
		"id",
		"isDefault",
		"name",
		"updatedAt",
		"url",
	)
}
