package action

import (
	"fmt"
	"strconv"
	"time"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type environment struct {
	Id        int
	Name      string
	CreatedAt time.Time `json:"created_at"`
}

type environmentQuery struct {
	Environments []environment
}

func ActionEnvironments(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repo, err := repoOverride(cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var queryResult environmentQuery
		return ApiV3Action(cmd, fmt.Sprintf(`repos/%v/%v/environments`, repo.RepoOwner(), repo.RepoName()), &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, environment := range queryResult.Environments {
				vals = append(vals, environment.Name, strconv.Itoa(environment.Id))
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
