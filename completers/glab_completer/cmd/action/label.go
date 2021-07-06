package action

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type label struct {
	Name        string
	Description string
}

func ActionLabels(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		//repo := "" // TODO from config or var

		var queryResult []label
		return actionApi(cmd, "projects/:fullpath/labels", &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, label := range queryResult {
				vals = append(vals, label.Name, label.Description)
			}
			return carapace.ActionValuesDescribed(vals...)
			//}).Cache(1*time.Hour, cache.String(repo.FullName()))
		})
	})
}
