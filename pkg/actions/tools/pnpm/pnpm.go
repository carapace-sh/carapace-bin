package pnpm

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
)

// ActionFilters completes filters
func ActionFilters() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if index := strings.LastIndexAny(c.Value, "[]"); index != -1 && []rune(c.Value)[index] == '[' {
			return git.ActionRefs(git.RefOption{}.Default()).Invoke(c).Prefix(c.Value[:index+1]).Suffix("]").ToA().NoSpace()
		}
		if index := strings.LastIndexAny(c.Value, "{}"); index != -1 && []rune(c.Value)[index] == '{' {
			prefix := c.Value[:index+1]
			c.Value = c.Value[index+1:]
			return carapace.ActionDirectories().Invoke(c).Prefix(prefix).ToA().NoSpace()
		}
		return carapace.Batch(
			ActionWorkspaceFilters(),
			ActionWorkspacePackages().NoSpace(),
		).ToA()
	})
}

// ActionDependencyNames completes dependency names from package.json
func ActionDependencyNames() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		pj, err := loadPackageJson(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		seen := make(map[string]bool)

		for name := range pj.Dependencies {
			if !seen[name] {
				seen[name] = true
				vals = append(vals, name)
			}
		}
		for name := range pj.DevDependencies {
			if !seen[name] {
				seen[name] = true
				vals = append(vals, name)
			}
		}
		for name := range pj.OptionalDependencies {
			if !seen[name] {
				seen[name] = true
				vals = append(vals, name)
			}
		}
		return carapace.ActionValues(vals...).Tag("dependencies")
	})
}

// ActionDependencies completes dependencies with their version from package.json
func ActionDependencies() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		pj, err := loadPackageJson(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		seen := make(map[string]bool)

		// Helper to add dependencies with description (version)
		addDeps := func(deps map[string]string) {
			for name, version := range deps {
				if !seen[name] {
					seen[name] = true
					vals = append(vals, name, version)
				}
			}
		}

		addDeps(pj.Dependencies)
		addDeps(pj.DevDependencies)
		addDeps(pj.OptionalDependencies)

		return carapace.ActionValuesDescribed(vals...).Tag("dependencies")
	})
}

// ActionStorePath completes store paths
func ActionStorePath() carapace.Action {
	return carapace.ActionExecCommand("pnpm", "store", "path")(func(output []byte) carapace.Action {
		path := strings.TrimSpace(string(output))
		return carapace.ActionValues(path)
	})
}

// ActionStoreStatus completes store status information
func ActionStoreStatus() carapace.Action {
	return carapace.ActionExecCommand("pnpm", "store", "status")(func(output []byte) carapace.Action {
		lines := strings.Split(strings.TrimSpace(string(output)), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if line != "" {
				vals = append(vals, line)
			}
		}
		return carapace.ActionValues(vals...)
	})
}

// ActionStorePackages completes packages in the store
func ActionStorePackages() carapace.Action {
	return carapace.ActionExecCommand("pnpm", "list", "--parseable", "--depth", "0")(func(output []byte) carapace.Action {
		lines := strings.Split(strings.TrimSpace(string(output)), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if line != "" {
				if idx := strings.LastIndex(line, "/"); idx != -1 {
					vals = append(vals, line[idx+1:])
				} else {
					vals = append(vals, line)
				}
			}
		}
		return carapace.ActionValues(vals...).Tag("store packages")
	})
}

// ActionStorePrune completes store prune options
func ActionStorePrune() carapace.Action {
	return carapace.ActionValues()
}
