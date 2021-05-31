package action

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type issue struct {
	Iid   int
	Title string
}

func ActionIssues(cmd *cobra.Command, state string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		stateQuery := ""
		if state != "" {
			stateQuery = "&state=" + url.QueryEscape(state)
		}

		var queryResult []issue
		return actionApi(cmd, fmt.Sprintf("projects/:fullpath/issues?order_by=updated_at&per_page=100%v", stateQuery), &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, issue := range queryResult {
				vals = append(vals, strconv.Itoa(issue.Iid), issue.Title)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
