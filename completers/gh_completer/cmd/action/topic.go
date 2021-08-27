package action

import (
	"fmt"
	"net/url"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type topic struct {
	Name             string
	ShortDescription string `json:"short_description"`
}

type topicQuery struct {
	Items []topic
}

func ActionTopicSearch(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.CallbackValue) < 2 {
			return carapace.ActionMessage("topic search needs at least two characters")
		}

		var queryResult topicQuery
		return ApiV3Action(cmd, fmt.Sprintf(`search/topics?per_page=100&q=%v`, url.QueryEscape(c.CallbackValue)), &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult.Items)*2)
			for _, topic := range queryResult.Items {
				vals = append(vals, topic.Name, topic.ShortDescription)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
