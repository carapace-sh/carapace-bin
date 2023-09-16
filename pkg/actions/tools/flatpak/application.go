package flatpak

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionApplications completes applications
//
//	org.kde.Platform (Shared libraries used by KDE applications)
//	org.qutebrowser.qutebrowser (A keyboard-driven web browser)
func ActionApplications() carapace.Action {
	return carapace.ActionExecCommand("flatpak", "list", "--columns", "application,description")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines {
			if splitted := strings.SplitN(line, "\t", 2); len(splitted) == 2 {
				vals = append(vals, splitted...)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
