package newrelic

import (
	"encoding/json"

	"github.com/rsteube/carapace"
)

func actionNerdGraph[T any](query string, transform func(result T) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("newrelic", "nerdgraph", "query", query)(func(output []byte) carapace.Action {
			var result T
			if err := json.Unmarshal(output, &result); err != nil {
				println(string(output))
				return carapace.ActionMessage(err.Error())
			}
			return transform(result)
		})
	})
}
