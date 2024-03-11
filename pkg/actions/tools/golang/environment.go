package golang

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

// ActionEnvironmentVariables completes go environment variables.
//
//	GOHOSTARCH (amd64)
//	GOHOSTOS (linux)
func ActionEnvironmentVariables() carapace.Action {
	return carapace.ActionExecCommand("go", "env", "--json")(func(output []byte) carapace.Action {
		var env map[string]string
		if err := json.Unmarshal(output, &env); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for key, value := range env {
			vals = append(vals, key, value)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
