// package turbo contains turbo related actions
package turbo

import (
	"encoding/json"
	"os"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/util"
)

type turbo struct {
	Pipeline map[string]interface{}
}

// ActionPipelineTasks completes pipeline tasks
//
//	build
//	deploy
func ActionPipelineTasks() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		path, err := util.FindReverse(c.Dir, "turbo.json")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var t turbo
		if err := json.Unmarshal(content, &t); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for task := range t.Pipeline {
			vals = append(vals, task)
		}
		return carapace.ActionValues(vals...)
	})
}
