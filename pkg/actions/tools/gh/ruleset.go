package gh

import (
	"fmt"
	"strconv"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type ruleset struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Enforcement string `json:"enforcement"`
}

// ActionRulesets completes rulesets
//
//	1 (Rules)
//	2 (Another)
func ActionRulesets(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return resolveRepo(c, opts, func(resolved RepoOpts) carapace.Action {
			var queryResult []ruleset
			return apiV3Action(resolved, fmt.Sprintf("repos/%v/%v/rulesets", resolved.Owner, resolved.Name), &queryResult, func() carapace.Action {
				vals := make([]string, 0)
				for _, ruleset := range queryResult {
					switch ruleset.Enforcement {
					case "active":
						vals = append(vals, strconv.Itoa(ruleset.ID), ruleset.Name, style.Green)
					default:
						vals = append(vals, strconv.Itoa(ruleset.ID), ruleset.Name, style.Red)
					}
				}
				return carapace.ActionStyledValuesDescribed(vals...)
			})
		})
	}).Tag("rulesets")
}
