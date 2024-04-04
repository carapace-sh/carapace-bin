package action

import (
	"os"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionScenarios completes molecule scenarios
// Does not use ActionDirectories since this completion is not recursive
func ActionScenarios() carapace.Action {
	moleculeFiles, err := os.ReadDir("molecule")
	if err != nil {
		return carapace.ActionValues()
	}

	scenarios := make([]string, 0)
	for _, entry := range moleculeFiles {
		if !entry.IsDir() && strings.HasPrefix(".", entry.Name()) {
			continue
		}
		scenarios = append(scenarios, strings.TrimSuffix("/", entry.Name()))
	}

	return carapace.ActionValues(scenarios...).StyleF(style.ForPath).Tag("scenarios")
}
