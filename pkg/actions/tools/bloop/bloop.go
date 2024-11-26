package bloop

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func autocomplete(mode string) carapace.Action {
	return carapace.ActionExecCommand("bloop", "autocomplete", "--format", "fish", "--mode", mode)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines {
			if line != "" {
				value, description, _ := strings.Cut(line, "\t")
				vals = append(vals, value, description)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionProjects completes projects
//
//	one
//	two
func ActionProjects() carapace.Action {
	return autocomplete("projects").Tag("projects")
}

// ActionProtocols completes protocols
//
//	local
//	tcp
func ActionProtocols() carapace.Action {
	return autocomplete("protocols").Tag("protocols")
}

// ActionReporters completes reporters
//
//	scalac
//	bloop
func ActionReporters() carapace.Action {
	return autocomplete("reporters").Tag("reporters")
}
