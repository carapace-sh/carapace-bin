package pixi

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

type pixiGlobalEnvironment struct {
	Name         string                 `json:"name"`
	Dependencies []pixiGlobalDependency `json:"dependencies"`
}

type pixiGlobalDependency struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ActionGlobalEnvironments completes global environment names
//
//	gh (2.86.0)
//	jj (0.38.0)
func ActionGlobalEnvironments() carapace.Action {
	return actionExecPixi("global", "list", "--json")(func(output []byte) carapace.Action {
		var environments []pixiGlobalEnvironment
		if err := json.Unmarshal(output, &environments); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0, len(environments)*2)
		for _, environment := range environments {
			version := ""
			for _, dependency := range environment.Dependencies {
				if dependency.Name == environment.Name {
					version = dependency.Version
					break
				}
			}
			vals = append(vals, environment.Name, version)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("global environments")
}
