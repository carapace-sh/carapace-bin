package yarn

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionBins completes bins
//
//	example (/tmp/yarn/example.js)
//	another (/tmp/yarn/another.js)
func ActionBins() carapace.Action {
	return actionYarn("bin", "--json")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines[:len(lines)-1] {
			var bin struct {
				Name string
				Path string
			}
			if err := json.Unmarshal([]byte(line), &bin); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			vals = append(vals, bin.Name, bin.Path)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
