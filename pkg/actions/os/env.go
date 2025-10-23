package os

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionEnvironmentVariables completes environment values
// see also env.ActionNameValues
//
//	SHELL (/bin/elvish)
//	LANG (en_US.utf8)
func ActionEnvironmentVariables() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		m := make(map[string]string)
		for _, e := range c.Env {
			if name, value, ok := strings.Cut(e, "="); ok {
				m[name] = value
			}
		}

		vals := make([]string, 0, len(m))
		for key, value := range m {
			vals = append(vals, key, value)
		}
		return carapace.ActionValuesDescribed(vals...).Filter(
			"CARAPACE_COMPLINE",
			"CARAPACE_SHELL",
			"CARAPACE_SHELL_BUILTINS",
			"CARAPACE_SHELL_FUNCTIONS",
		)
	}).Tag("environment variables")
}
