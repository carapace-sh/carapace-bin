package conda

import (
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionEnvironments completes environments
func ActionEnvironments() carapace.Action {
	return carapace.ActionExecCommand("conda", "env", "list", "--json")(func(output []byte) carapace.Action {
		var result struct{ Envs []string }
		if err := json.Unmarshal(output, &result); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, env := range result.Envs {
			if strings.Contains(env, ".conda/envs/") {
				vals = append(vals, filepath.Base(env), env)
			} else {
				vals = append(vals, "base", env)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("environments")
}

// ActionEnvironmentVariables completes environment variables.
//
//	FIRST (1)
//	SECOND (2)
func ActionEnvironmentVariables(name string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if name == "" {
			name = "base"
		}
		return carapace.ActionExecCommand("conda", "env", "config", "vars", "list", "--name", name, "--json")(func(output []byte) carapace.Action {
			var variables map[string]string
			if err := json.Unmarshal(output, &variables); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for key, value := range variables {
				vals = append(vals, key, value)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
