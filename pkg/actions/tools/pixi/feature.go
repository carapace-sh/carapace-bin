package pixi

import "github.com/carapace-sh/carapace"

// ActionFeatures completes feature names
//
//	default
//	cuda
func ActionFeatures() carapace.Action {
	return actionExecPixi("info", "--json")(func(output []byte) carapace.Action {
		info, err := parseInfo(output)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		seen := make(map[string]bool)
		vals := make([]string, 0)
		for _, env := range info.EnvironmentsInfo {
			for _, f := range env.Features {
				if !seen[f] {
					seen[f] = true
					vals = append(vals, f)
				}
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("features")
}
