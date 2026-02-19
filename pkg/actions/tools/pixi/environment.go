package pixi

import "github.com/carapace-sh/carapace"

// ActionEnvironments completes environment names
//
//	default
//	cuda
func ActionEnvironments() carapace.Action {
	return actionExecPixi("info", "--json")(func(output []byte) carapace.Action {
		info, err := parseInfo(output)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0, len(info.EnvironmentsInfo))
		for _, env := range info.EnvironmentsInfo {
			vals = append(vals, env.Name)
		}
		return carapace.ActionValues(vals...)
	}).Tag("environments")
}
