package action

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

type iteration struct {
	Id        int
	Title     string
	State     string
	StartDate string `json:"start_date"`
	DueDate   string `json:"due_date"`
}

func ActionIterations(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		query := "projects/:fullpath/iterations?per_page=100"
		if flag := cmd.Flag("group"); flag != nil && flag.Value.String() != "" {
			query = fmt.Sprintf("groups/%v/iterations?per_page=100", url.PathEscape(flag.Value.String()))
		}

		var queryResult []iteration
		return actionApi(cmd, query, &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, iteration := range queryResult {
				description := iteration.State
				if iteration.StartDate != "" || iteration.DueDate != "" {
					description = fmt.Sprintf("%v - %v (%v)", iteration.StartDate, iteration.DueDate, iteration.State)
				}
				vals = append(vals, strconv.Itoa(iteration.Id), description)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
