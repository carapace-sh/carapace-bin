// package apt contains Advanced Package Tool related actions
package apt

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPackages completes installed packages
//
//	adduser
//	cpp
func ActionPackages() carapace.Action {
	return carapace.ActionExecCommand("apt", "list", "--installed")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines {
			if splitted := strings.SplitN(line, "/", 2); len(splitted) == 2 {
				vals = append(vals, splitted...)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionPackageSearch completes installable packages
//
//	git
//	git-man
func ActionPackageSearch() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("apt-cache", "search", "^"+c.Value)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines {
				if splitted := strings.SplitN(line, " - ", 2); len(splitted) == 2 {
					vals = append(vals, splitted...)
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
