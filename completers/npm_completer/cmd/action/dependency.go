package action

import (
	"encoding/json"
	"github.com/rsteube/carapace"
)

type dependency struct {
	Version      string                `json:",omitempty"`
	Dependencies map[string]dependency `json:",omitempty"`
}

func (d dependency) flatMap() map[string][]string {
	deps := make(map[string][]string)
	for name, dep := range d.Dependencies {
		deps[name] = append(deps[name], dep.Version)
		for key, value := range dep.flatMap() {
			deps[key] = append(deps[key], value...)
		}
	}
	return deps
}

func ActionDependencyNames() carapace.Action {
	return carapace.ActionExecCommand("npm", "ls", "--json", "--all")(func(output []byte) carapace.Action {
		var deps dependency
		if err := json.Unmarshal(output, &deps); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for name := range deps.flatMap() {
			vals = append(vals, name)
		}
		return carapace.ActionValues(vals...)
	})
}

func ActionDependencies() carapace.Action {
	return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("npm", "ls", "--json", "--all")(func(output []byte) carapace.Action {
			var deps dependency
			if err := json.Unmarshal(output, &deps); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			// dirty fix for dependencies with `@` prefix (which will be split by ActionMultiParts)
			if len(c.Parts) > 0 && c.Parts[0] == "" {
				c.Parts = c.Parts[1:]
				if len(c.Parts) == 0 {
					c.Parts = append(c.Parts, "@")
				} else {
					c.Parts[0] = "@" + c.Parts[0]
				}
			}

			switch len(c.Parts) {
			case 0:
				vals := make([]string, 0)
				for name := range deps.flatMap() {
					vals = append(vals, name)
				}
				return carapace.ActionValues(vals...)
			case 1:
				return carapace.ActionValues(deps.flatMap()[c.Parts[0]]...)
			default:
				return carapace.ActionValues()
			}
		})
	})
}
