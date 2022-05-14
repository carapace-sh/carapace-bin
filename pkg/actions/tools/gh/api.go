package gh

import (
	"encoding/json"

	"github.com/rsteube/carapace"
)

func apiV3Action(repo repo, query string, v interface{}, transform func() carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("gh", "api", "--hostname", repo.host(), "--preview", "mercy", query)(func(output []byte) carapace.Action {
			if err := json.Unmarshal(output, &v); err != nil {
				return carapace.ActionMessage("failed to unmarshall response: " + err.Error())
			}
			return transform()
		})
	})
}
