package gh

import "github.com/carapace-sh/carapace"

type coc struct {
	Key  string
	Name string
}

// ActionCocs completes codes of conduct
//
//	agpl-3.0 (GNU Affero General Public License v3.0)
//	mit (MIT License)
func ActionCocs(opts HostOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []coc
		return apiV3Action(opts.repo(), `codes_of_conduct`, &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, coc := range queryResult {
				vals = append(vals, coc.Key, coc.Name)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("code of conducts")
}
