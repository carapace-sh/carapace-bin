package pixi

import "github.com/carapace-sh/carapace"

// ActionTasks completes task names
//
//	build
//	generate
func ActionTasks() carapace.Action {
	return actionExecPixi("info", "--json")(func(output []byte) carapace.Action {
		info, err := parseInfo(output)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		tasks := make(map[string]bool)
		for _, env := range info.EnvironmentsInfo {
			for _, t := range env.Tasks {
				tasks[t] = true
			}
		}

		vals := make([]string, 0, len(tasks))
		for t := range tasks {
			vals = append(vals, t)
		}
		return carapace.ActionValues(vals...)
	}).Tag("tasks")
}
