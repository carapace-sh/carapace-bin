package winetricks

import (
	"fmt"
	"os"

	"github.com/rsteube/carapace"
)

// ActionPrefixes completes prefixes
//
//	3m_library (3M Cloud Library (3M Company, 2015) [downloadable])
//	7zip (7-Zip 19.00 (Igor Pavlov, 2019) [downloadable])
func ActionPrefixes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		home, err := os.UserHomeDir()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		files, err := os.ReadDir(fmt.Sprintf("%v/.local/share/wineprefixes/", home))
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, file := range files {
			if file.IsDir() {
				vals = append(vals, file.Name())
			}
		}
		return carapace.ActionValues(vals...)
	})
}
