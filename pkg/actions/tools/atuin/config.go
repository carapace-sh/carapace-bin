package atuin

import (
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/pelletier/go-toml"
)

// ActionConfigKeys completes config keys
//
//	sync
//	sync.records (true)
func ActionConfigKeys() carapace.Action {
	return carapace.ActionExecCommand("atuin", "config", "print")(func(output []byte) carapace.Action {
		var entries map[string]any
		if err := toml.Unmarshal(output, &entries); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for k, v := range entries {
			switch v := v.(type) {
			case map[string]any:
				vals = append(vals, k, "")
				for nk, v := range v { // TODO do further nested maps exist?
					vals = append(vals, fmt.Sprintf("%v.%v", k, nk), fmt.Sprintf("%v", v))
				}
			default:
				vals = append(vals, k, fmt.Sprintf("%v", v))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).MultiParts(".")
}
