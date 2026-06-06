package mcp

import (
	"sort"

	"github.com/carapace-sh/carapace-bin/pkg/actions"
)

func (s *MCPServer) handleListMacros() (map[string]any, error) {
	var names []string
	for name := range actions.Macros {
		names = append(names, name)
	}
	sort.Strings(names)

	macros := make([]map[string]string, 0, len(names))
	for _, name := range names {
		m := actions.Macros[name]
		sig := m.Signature()
		if sig == "" {
			sig = "—"
		}
		macros = append(macros, map[string]string{
			"name":        "carapace." + name,
			"signature":   sig,
			"description": m.Description,
			"reference":   m.Function,
		})
	}

	return mcpJSONResult(macros), nil
}
