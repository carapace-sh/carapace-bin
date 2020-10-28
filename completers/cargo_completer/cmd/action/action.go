package action

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
)

type manifest struct {
	Targets []struct {
		Name string
		Kind []string
	}
}

func ActionColorModes() carapace.Action {
	return carapace.ActionValues("auto", "always", "never")
}

func ActionTargets() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("cargo", "read-manifest").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			var m manifest
			if err := json.Unmarshal(output, &m); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				vals := make([]string, len(m.Targets)*2)
				for index, target := range m.Targets {
					vals[index*2] = target.Name
					vals[(index*2)+1] = strings.Join(target.Kind, ", ")
				}
				return carapace.ActionValuesDescribed(vals...)
			}
		}
	})
}
