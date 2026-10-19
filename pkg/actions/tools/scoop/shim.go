package scoop

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
)

type scoopShimInfo struct {
	Name   string `json:"Name"`
	Source string `json:"Source"`
}

// ActionShims completes installed shim names with their source as description
//
//	7z (7zip)
//	python (python)
func ActionShims() carapace.Action {
	return carapace.ActionExecCommand("powershell", "-NoProfile", "-Command", "scoop shim list | ConvertTo-Json -Compress")(func(output []byte) carapace.Action {
		output = []byte(strings.TrimSpace(string(output)))
		if len(output) == 0 || string(output) == "null" {
			return carapace.ActionValues()
		}
		var shims []scoopShimInfo
		if err := json.Unmarshal(output, &shims); err != nil {
			var shim scoopShimInfo
			if err2 := json.Unmarshal(output, &shim); err2 != nil {
				return carapace.ActionMessage(err.Error())
			}
			return carapace.ActionValuesDescribed(shim.Name, shim.Source).Tag("shims")
		}
		vals := make([]string, 0)
		for _, shim := range shims {
			vals = append(vals, shim.Name, shim.Source)
		}
		return carapace.ActionValuesDescribed(vals...).Tag("shims")
	})
}
