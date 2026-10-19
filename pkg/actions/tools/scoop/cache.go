package scoop

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
)

type scoopCacheInfo struct {
	Name    string `json:"Name"`
	Version string `json:"Version"`
}

// ActionCachedApps completes cached app names with their version as description
//
//	bat (0.24.0)
//	git (2.45.0)
func ActionCachedApps() carapace.Action {
	return carapace.ActionExecCommand("powershell", "-NoProfile", "-Command", "scoop cache show | ConvertTo-Json -Compress")(func(output []byte) carapace.Action {
		output = []byte(strings.TrimSpace(string(output)))
		if len(output) == 0 || string(output) == "null" {
			return carapace.ActionValues()
		}
		var cache []scoopCacheInfo
		if err := json.Unmarshal(output, &cache); err != nil {
			var item scoopCacheInfo
			if err2 := json.Unmarshal(output, &item); err2 != nil {
				return carapace.ActionMessage(err.Error())
			}
			return carapace.ActionValuesDescribed(item.Name, item.Version).Tag("cached apps")
		}
		vals := make([]string, 0)
		for _, item := range cache {
			vals = append(vals, item.Name, item.Version)
		}
		return carapace.ActionValuesDescribed(vals...).Tag("cached apps")
	})
}
