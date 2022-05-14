package gh

import (
	"github.com/rsteube/carapace"
)

type license struct {
	Key  string
	Name string
}

// ActionLicenses completes licenses
//   apache-2.0 (Apache License 2.0)
//   bsd-2-clause (BSD 2-Clause "Simplified" License)
func ActionLicenses(opts HostOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []license
		return apiV3Action(opts, `licenses`, &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, license := range queryResult {
				vals = append(vals, license.Key, license.Name)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
