package pnpm

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/util"
)

type pnpmWorkspace struct {
	Packages []string `json:"packages"`
}

// ActionWorkspaces completes workspaces
func ActionWorkspaces() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// First try to find pnpm-workspace.yaml
		if workspaceFile, err := util.FindReverse(c.Dir, "pnpm-workspace.yaml"); err == nil {
			return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				content, err := os.ReadFile(workspaceFile)
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}

				// Parse YAML-like structure for packages
				lines := strings.Split(string(content), "\n")
				vals := make([]string, 0)
				inPackages := false

				for _, line := range lines {
					line = strings.TrimSpace(line)
					if line == "packages:" {
						inPackages = true
						continue
					}
					if inPackages && strings.HasPrefix(line, "-") {
						packagePath := strings.TrimSpace(strings.TrimPrefix(line, "-"))
						vals = append(vals, packagePath)
					}
				}

				return carapace.ActionValues(vals...)
			})
		}

		// Fallback to package.json workspaces
		if pj, err := loadPackageJson(c); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionValues(pj.Workspaces...)
		}
	})
}

// ActionWorkspacePackages completes packages within workspaces
func ActionWorkspacePackages() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommandE("pnpm", "ls", "-r", "--json", "--depth", "-1")(func(output []byte, err error) carapace.Action {
			if err != nil {
				return carapace.ActionValues()
			}

			var packages []struct {
				Name    string `json:"name"`
				Version string `json:"version"`
				Path    string `json:"path"`
				Private bool   `json:"private"`
			}

			if err := json.Unmarshal(output, &packages); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0, len(packages)*2)
			for _, pkg := range packages {
				if pkg.Name != "" && !pkg.Private {
					desc := pkg.Version
					if desc == "" {
						desc = pkg.Path
					}
					vals = append(vals, pkg.Name, desc)
				}
			}
			return carapace.ActionValuesDescribed(vals...).Tag("workspace packages")
		})
	})
}
