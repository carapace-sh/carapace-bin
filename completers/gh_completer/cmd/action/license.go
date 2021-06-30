package action

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type license struct {
	Key  string
	Name string
}

func ActionLicenses(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []license
		return ApiV3Action(cmd, `licenses`, &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, license := range queryResult {
				vals = append(vals, license.Key, license.Name)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
