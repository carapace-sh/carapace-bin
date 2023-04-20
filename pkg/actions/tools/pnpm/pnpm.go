package pnpm

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/npm"
)

// ActionFilter completes filter
func ActionFilter() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if index := strings.LastIndexAny(c.Value, "[]"); index != -1 && []rune(c.Value)[index] == '[' {
			return git.ActionRefs(git.RefOption{}.Default()).Invoke(c).Prefix(c.Value[:index+1]).Suffix("]").ToA().NoSpace()
		}
		if index := strings.LastIndexAny(c.Value, "{}"); index != -1 && []rune(c.Value)[index] == '{' {
			prefix := c.Value[:index+1]
			c.Value = c.Value[index+1:]
			return carapace.ActionDirectories().Invoke(c).Prefix(prefix).ToA().NoSpace()
		}
		return npm.ActionDependencies().NoSpace()
	})
}

type location struct {
	Dependencies map[string]struct {
		Version string
	}
}

// ActionDependencyNames completes dependency names
func ActionDependencyNames() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("pnpm", "list", "--json")(func(output []byte) carapace.Action {
			var locations []location
			if err := json.Unmarshal(output, &locations); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, l := range locations {
				for name := range l.Dependencies {
					vals = append(vals, name)
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}

// ActionDependencies ocmpletes dependencies with their version
func ActionDependencies() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("pnpm", "list", "--json")(func(output []byte) carapace.Action {
			var locations []location
			if err := json.Unmarshal(output, &locations); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, l := range locations {
				for name, details := range l.Dependencies {
					vals = append(vals, fmt.Sprintf("%v@%v", name, details.Version))
				}
			}
			return carapace.ActionValues(vals...)
		}).Invoke(c).ToMultiPartsA("@")
	})
}
