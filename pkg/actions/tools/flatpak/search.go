package flatpak

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionApplicationSearch completes installable applications
//
//	org.qutebrowser.qutebrowser (A keyboard-driven web browser)
//	org.qutebrowser.qutebrowser.Userscripts (qutebrowser userscripts supporting Python modules, shared libraries, and exec...)
func ActionApplicationSearch() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("flatpak", "search", "--columns", "application,description", c.Value)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines {
				if splitted := strings.SplitN(line, "\t", 2); len(splitted) == 2 {
					vals = append(vals, splitted...)
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
