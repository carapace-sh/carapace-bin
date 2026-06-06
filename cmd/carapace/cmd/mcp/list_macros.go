package mcp

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace-bin/pkg/actions"
)

type mcpListMacrosRequest struct {
	Executable string `json:"executable,omitempty"`
}

type macroEntry struct {
	Name        string `json:"name"`
	Signature   string `json:"signature"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Reference   string `json:"reference,omitempty"`
	Function    string `json:"function,omitempty"`
}

type macroList struct {
	Version string       `json:"version"`
	Macros  []macroEntry `json:"macros"`
}

func (s *MCPServer) handleListMacros(args json.RawMessage) (map[string]any, error) {
	var request mcpListMacrosRequest
	if len(args) != 0 {
		if err := json.Unmarshal(args, &request); err != nil {
			return nil, fmt.Errorf("invalid list_macros arguments: %w", err)
		}
	}

	if request.Executable != "" {
		executable, err := s.resolveMacroExecutable(request.Executable)
		if err != nil {
			return mcpTextResult(err.Error(), true), nil
		}
		output, err := runCommand(executable, []string{"_carapace", "macro"}, os.Environ())
		if err != nil {
			return mcpTextResult(err.Error(), true), nil
		}
		return mcpTextResult(output, false), nil
	}

	buildInfo, _ := debug.ReadBuildInfo()
	depVersions := resolveDepVersions(buildInfo)

	var names []string
	for name := range actions.Macros {
		names = append(names, name)
	}
	sort.Strings(names)

	entries := make([]macroEntry, 0, len(names))
	for _, name := range names {
		m := actions.Macros[name]
		sig := m.Signature()
		if sig == "" {
			sig = "—"
		}
		pkgPath, _, _ := strings.Cut(m.Function, "#")
		entries = append(entries, macroEntry{
			Name:        "carapace." + name,
			Signature:   sig,
			Description: m.Description,
			Version:     depVersions[pkgPath],
			Reference:   m.Function,
		})
	}

	return mcpJSONResult(macroList{
		Version: depVersions["github.com/carapace-sh/carapace-bin"],
		Macros:  entries,
	}), nil
}

func resolveDepVersions(info *debug.BuildInfo) map[string]string {
	m := make(map[string]string)
	if info == nil {
		return m
	}
	for _, dep := range info.Deps {
		m[dep.Path] = dep.Version
	}
	return m
}
