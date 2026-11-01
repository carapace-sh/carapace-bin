package action

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

type iterationState string

// GitLab serializes the iteration state as an integer enum in the REST API,
// though some instances return the plain string representation.
func (s *iterationState) UnmarshalJSON(data []byte) error {
	var state string
	if err := json.Unmarshal(data, &state); err == nil {
		*s = iterationState(state)
		return nil
	}

	var n int
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	switch n {
	case 1:
		*s = "upcoming"
	case 2:
		*s = "current"
	case 3:
		*s = "closed"
	default:
		*s = iterationState(strconv.Itoa(n))
	}
	return nil
}

type iteration struct {
	Id        int
	Title     string
	State     iterationState
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
				var s string
				switch string(iteration.State) {
				case "upcoming":
					s = style.Blue
				case "current":
					s = style.Green
				case "closed":
					s = style.Red
				default:
					s = style.Default
				}

				description := string(iteration.State)
				if iteration.StartDate != "" || iteration.DueDate != "" {
					description = fmt.Sprintf("%v - %v (%v)", iteration.StartDate, iteration.DueDate, iteration.State)
				}
				vals = append(vals, strconv.Itoa(iteration.Id), description, s)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	}).Tag("iterations")
}
