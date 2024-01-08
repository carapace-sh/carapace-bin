package npm

import "github.com/rsteube/carapace"

// ActionScripts completes scripts
func ActionScripts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if pj, err := loadPackageJson(c); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)
			for name := range pj.Scripts {
				vals = append(vals, name)
			}
			return carapace.ActionValues(vals...)
		}
	})
}
