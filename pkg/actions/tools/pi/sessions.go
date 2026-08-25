package pi

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionSessions completes session files
//
//	2026-08-25T14-30-00-000Z_01J8XYZABC1234567890.jsonl
//	2026-08-24T10-00-00-000Z_01J8XYZABC1234567890.jsonl
func ActionSessions() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		home, err := os.UserHomeDir()
		if err != nil {
			return carapace.ActionValues()
		}
		sessionsDir := filepath.Join(home, ".pi", "agent", "sessions")
		entries, err := os.ReadDir(sessionsDir)
		if err != nil {
			return carapace.ActionValues()
		}

		vals := make([]string, 0)
		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}
			sessionFiles, err := filepath.Glob(filepath.Join(sessionsDir, entry.Name(), "*.jsonl"))
			if err != nil {
				continue
			}
			for _, sf := range sessionFiles {
				name := filepath.Base(sf)
				name = strings.TrimSuffix(name, ".jsonl")
				vals = append(vals, name)
			}
		}
		return carapace.ActionValues(vals...).Tag("sessions")
	})
}

// ActionTools completes built-in tool names
//
//	read
//	bash
func ActionTools() carapace.Action {
	return carapace.ActionValues(
		"read",
		"bash",
		"powershell",
		"edit",
		"write",
		"grep",
		"find",
		"ls",
	).Tag("tools")
}
