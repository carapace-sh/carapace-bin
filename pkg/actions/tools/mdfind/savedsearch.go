package mdfind

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionSavedSearches completes Spotlight saved search names
//
//	Recent Files
//	Search Terms
func ActionSavedSearches() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		home, err := os.UserHomeDir()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		dir := filepath.Join(home, "Library", "Saved Searches")
		entries, err := os.ReadDir(dir)
		if err != nil {
			return carapace.ActionValues()
		}

		vals := make([]string, 0)
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".savedSearch") {
				name := strings.TrimSuffix(entry.Name(), ".savedSearch")
				vals = append(vals, name)
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("saved searches").Uid("tools.mdfind", "savedsearches")
}
