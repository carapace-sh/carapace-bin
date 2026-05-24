package npm

import "github.com/carapace-sh/carapace"

// ActionScripts completes scripts
//
//	test (bun test --isolate)
//	test:coverage (bun test --isolate --coverage)
func ActionScripts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if pj, err := loadPackageJson(c); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)
			for name, command := range pj.Scripts {
				vals = append(vals, name, command)
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	}).Tag("scripts")
}
