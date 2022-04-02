package action

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
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

func ActionLabels(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult labelsQuery
		return GraphQlAction(cmd, `repository(owner: $owner, name: $repo){ labels(first: 100) { nodes { name, description, color } } }`, &queryResult, func() carapace.Action {
			labels := queryResult.Data.Repository.Labels.Nodes
			vals := make([]string, 0)
			for _, label := range labels {
				vals = append(vals, label.Name, label.Description, style.Hex256(label.Color))
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	}).Cache(5*time.Minute, repoCacheKey(cmd))
}
