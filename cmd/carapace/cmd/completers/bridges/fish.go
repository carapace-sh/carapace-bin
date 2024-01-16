package bridges

import (
	"os"
	"runtime"
	"strings"

	"github.com/rsteube/carapace/pkg/execlog"
)

var fishFilter = []string{}

func Fish() []string {
	switch runtime.GOOS {
	case "windows":
		return []string{}
	}

	if _, err := execlog.LookPath("fish"); err != nil {
		return []string{}
	}

	// TODO handle different OS/locations
	entries, err := os.ReadDir("/usr/share/fish/completions")
	if err != nil {
		return []string{}
	}

	unique := make(map[string]bool)
	for _, entry := range entries {
		if !entry.IsDir() && entry.Name() != "..fish" && strings.HasSuffix(entry.Name(), ".fish") {
			unique[strings.TrimSuffix(entry.Name(), ".fish")] = true
		}
	}

	completers := make([]string, 0)
	for c := range filter(unique, genericFilter, fishFilter) {
		completers = append(completers, c)
	}
	return completers
}
