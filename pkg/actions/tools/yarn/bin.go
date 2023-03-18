package yarn

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionBins completes bins
//
//	example (/tmp/yarn/example.js)
//	another (/tmp/yarn/another.js)
func ActionBins() carapace.Action {
	return carapace.ActionExecCommandE("yarn", "bin", "--json")(func(output []byte, err error) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		if err != nil {
			if _, ok := err.(*exec.ExitError); ok {
				return carapace.ActionMessage(lines[0]) // error is printed to stdout
			}
			return carapace.ActionMessage(err.Error())
		}

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
