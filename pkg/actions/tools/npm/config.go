package npm

import (
	"encoding/json"

	"github.com/rsteube/carapace"
)

// ActionLocalConfigKeys completes local config keys
func ActionLocalConfigKeys() carapace.Action {
	return actionConfigKeys(false)
}

// ActionGlobalConfigKeys completes global config keys
func ActionGlobalConfigKeys() carapace.Action {
	return actionConfigKeys(false)
}

func actionConfigKeys(global bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {

		args := []string{"config", "list", "--json"}
		if global {
			args = append(args, "--global")
		}

		return carapace.ActionExecCommand("npm", args...)(func(output []byte) carapace.Action {
			var config map[string]interface{}
			if err := json.Unmarshal(output, &config); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0, len(config))
			for key := range config {
				vals = append(vals, key)
			}
			return carapace.ActionValues(vals...)
		})
	}).Tag("config keys")
}
