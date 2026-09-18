package compose

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

type project struct {
	Name   string
	Status string
}

// ActionProjects completes compose projects
//
//	carapace-bin (running)
func ActionProjects() carapace.Action {
	return carapace.ActionExecCommand("docker", "compose", "ls", "--format", "json")(func(output []byte) carapace.Action {
		var projects []project
		if err := json.Unmarshal(output, &projects); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, project := range projects {
			vals = append(vals, project.Name, project.Status)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
