package gh

import (
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type team struct {
	Name        string `json:"name"`
	ID          int    `json:"id"`
	Description string `json:"description"`
	Privacy     string `json:"privacy"`
}

// ActionTeams completes teams
//
//	first-team (description of first team)
//	second-team (description of second team)
func ActionTeams(opts OwnerOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []team
		return apiV3Action(opts.repo(), fmt.Sprintf(`orgs/%v/teams`, opts.Owner), &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, team := range queryResult {
				switch team.Privacy {
				case "secret":
					vals = append(vals, team.Name, team.Description, style.Carapace.KeywordNegative)
				default:
					vals = append(vals, team.Name, team.Description, style.Default)
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})

	})
}
