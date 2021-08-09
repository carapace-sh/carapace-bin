package action

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type mergeRequest struct {
	Iid   int
	Title string
}

func ActionMergeRequests(cmd *cobra.Command, state string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		stateQuery := ""
		if state != "" {
			stateQuery = "&state=" + url.QueryEscape(state)
		}

		var queryResult []mergeRequest
		return actionApi(cmd, fmt.Sprintf("projects/:fullpath/merge_requests?order_by=updated_at&per_page=100%v", stateQuery), &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, mr := range queryResult {
				vals = append(vals, strconv.Itoa(mr.Iid), mr.Title)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func ActionMergeRequestsAndBranches(cmd *cobra.Command, state string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.Batch(
			ActionBranches(cmd),
			ActionMergeRequests(cmd, state),
		).Invoke(c).Merge().ToA()
	})
}
