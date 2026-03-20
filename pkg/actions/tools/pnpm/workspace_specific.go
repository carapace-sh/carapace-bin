package pnpm

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/util"
	"gopkg.in/yaml.v3"
)

// ActionWorkspaceScripts completes scripts available across all workspace packages
func ActionWorkspaceScripts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		workspaceFile, err := util.FindReverse(c.Dir, "pnpm-workspace.yaml")
		if err != nil {
			return ActionScripts()
		}

		content, err := os.ReadFile(workspaceFile)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		workspaceDir := filepath.Dir(workspaceFile)

		var ws workspaceYaml
		if err := yaml.Unmarshal(content, &ws); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		seen := make(map[string]bool)

		for _, pattern := range ws.Packages {
			matches, err := filepath.Glob(filepath.Join(workspaceDir, pattern))
			if err != nil {
				continue
			}
			for _, match := range matches {
				data, err := os.ReadFile(filepath.Join(match, "package.json"))
				if err != nil {
					continue
				}
				var pj packageJson
				if json.Unmarshal(data, &pj) == nil {
					for name := range pj.Scripts {
						if !seen[name] {
							seen[name] = true
							vals = append(vals, name)
						}
					}
				}
			}
		}
		return carapace.ActionValues(vals...).Tag("workspace scripts")
	})
}

// ActionWorkspaceFilter completes common workspace filter patterns
func ActionWorkspaceFilter() carapace.Action {
	return ActionWorkspaceFilters()
}

// ActionWorkspaceFilters completes common workspace filter patterns
func ActionWorkspaceFilters() carapace.Action {
	return carapace.ActionValuesDescribed(
		"...", "Include all dependents (recursive)",
		"...^", "Include direct dependents only",
		"^...", "Include all dependencies (recursive)",
		"^...^", "Include direct dependencies only",
		"*", "All packages",
		"packages/*", "All packages in packages/",
		"apps/*", "All packages in apps/",
		"libs/*", "All packages in libs/",
		"tools/*", "All packages in tools/",
	).Tag("workspace filter patterns")
}

type location struct {
	Dependencies map[string]struct {
		Version string
	}
}

// ActionWorkspaceDependencies completes workspace package names with their versions
func ActionWorkspaceDependencies() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("pnpm", "list", "--json", "--depth", "0")(func(output []byte) carapace.Action {
			var locations []location
			if err := json.Unmarshal(output, &locations); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, l := range locations {
				for name, details := range l.Dependencies {
					vals = append(vals, name, details.Version)
				}
			}
			return carapace.ActionValuesDescribed(vals...).Tag("workspace dependencies")
		})
	})
}
