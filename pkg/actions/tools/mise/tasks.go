package mise

import (
	"encoding/json"
	"sort"

	"github.com/carapace-sh/carapace"
)

type task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ActionTasks completes mise tasks.
//
//	build (Build the project)
//	test (Run tests)
func ActionTasks() carapace.Action {
	return carapace.ActionExecCommand("mise", "tasks", "--json")(func(output []byte) carapace.Action {
		var tasks []task
		if err := json.Unmarshal(output, &tasks); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		sort.SliceStable(tasks, func(i, j int) bool { return tasks[i].Name < tasks[j].Name })

		vals := make([]string, 0, len(tasks)*2)
		for _, task := range tasks {
			if task.Name == "" {
				continue
			}
			vals = append(vals, task.Name, task.Description)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
