package completers

import (
	"os"
	"strings"
)

func FishCompleters() []string {
	// TODO handle different OS/locations
	entries, err := os.ReadDir("/usr/share/fish/completions")
	if err != nil {
		return []string{}
	}

	completers := make([]string, 0)
	for _, entry := range entries {
		if !entry.IsDir() && entry.Name() != "..fish" && strings.HasSuffix(entry.Name(), ".fish") {
			completers = append(completers, strings.TrimSuffix(entry.Name(), ".fish"))
		}
	}
	return completers
}
