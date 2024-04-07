package molecule

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionScenarios completes molecule scenarios
func ActionScenarios() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		entries, err := os.ReadDir(filepath.Join(c.Dir, "molecule"))
		if err != nil {
			return carapace.ActionValues()
		}

		scenarios := make([]string, 0)
		for _, entry := range entries {
			if !entry.IsDir() || strings.HasPrefix(entry.Name(), ".") {
				continue
			}
			if _, err := os.Stat(filepath.Join(c.Dir, "molecule", entry.Name(), "molecule.yml")); err == nil {
				scenarios = append(scenarios, entry.Name())
			}
		}

		return carapace.ActionValues(scenarios...)
	}).Tag("ansible scenarios")
}
