package env

import (
	"os/exec"

	"github.com/rsteube/carapace"
)

type variables struct {
	Condition func(c carapace.Context) bool
	Names     map[string]string
	Values    map[string]carapace.Action
}

func checkPath(s string) func(c carapace.Context) bool {
	return func(c carapace.Context) bool {
		_, err := exec.LookPath(s) // TODO copy function to carapace as this needs to use carapace.Context$Env
		return err == nil
	}
}

var knownVariables = map[string]variables{}

func ActionKnownEnvironmentVariables() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0)
		for _, v := range knownVariables {
			if v.Condition != nil && !v.Condition(c) {
				continue
			}

			for name, description := range v.Names {
				vals = append(vals, name, description)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("known environment variables")
}

func ActionEnvironmentVariableValues(s string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		for _, v := range knownVariables {
			if action, ok := v.Values[s]; ok {
				return action
			}
		}
		return carapace.ActionFiles()
	})
}
