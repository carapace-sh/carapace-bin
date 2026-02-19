package pixi

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
)

type pixiPackage struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ActionPackages completes installed package names
//
//	numpy (1.24.0)
//	python (3.11.0)
func ActionPackages() carapace.Action {
	return actionExecPixi("list", "--json")(func(output []byte) carapace.Action {
		var packages []pixiPackage
		if err := json.Unmarshal(output, &packages); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0, len(packages)*2)
		for _, p := range packages {
			vals = append(vals, p.Name, p.Version)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("packages")
}

// ActionPackageSearch completes conda package names from the registry
//
//	numpy (2.4.2)
//	numpyro (0.20.0)
func ActionPackageSearch() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if c.Value == "" {
			return carapace.ActionValues()
		}
		return carapace.ActionExecCommand("pixi", "search", "--limit", "20", c.Value+"*")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[1:] {
				fields := strings.Fields(line)
				if len(fields) >= 2 && fields[0] != "..." {
					vals = append(vals, fields[0], fields[1])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("packages")
}
