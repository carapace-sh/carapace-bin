package wt

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

// ActionHookTypes completes hook types
//
//	post-commit
//	post-create
func ActionHookTypes() carapace.Action {
	return carapace.ActionValues(
		"post-commit",
		"post-create",
		"post-merge",
		"post-remove",
		"post-start",
		"post-switch",
		"pre-commit",
		"pre-merge",
		"pre-remove",
		"pre-switch",
	).Tag("hook types")
}

// ActionHookCommands completes hook command names for a hook type
//
//	test (user)
//	lint (project)
func ActionHookCommands(hookType string) carapace.Action {
	return carapace.ActionExecCommandE("wt", "hook", "show", hookType, "--format", "json")(func(output []byte, err error) carapace.Action {
		if err != nil {
			return carapace.ActionValues()
		}

		var entries []struct {
			Source   string  `json:"source"`
			Name     *string `json:"name"`
			Template string  `json:"template"`
		}
		if err := json.Unmarshal(output, &entries); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, entry := range entries {
			if entry.Name != nil {
				vals = append(vals, *entry.Name, entry.Source)
			}
		}
		return carapace.ActionValuesDescribed(vals...).Tag("hook commands")
	})
}
