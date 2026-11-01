package asdf

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionInstalledVersions completes installed versions for a given plugin
//
//	3.2.2
//	3.3.0
func ActionInstalledVersions(plugin string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		name := plugin
		if name == "" && len(c.Args) > 0 {
			name = c.Args[0]
		}
		if name == "" {
			return carapace.ActionMessage("missing plugin name")
		}
		return carapace.ActionExecCommand("asdf", "list", name)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			var vals []string
			for _, line := range lines {
				line = strings.TrimPrefix(strings.TrimSpace(line), "*")
				if line != "" {
					vals = append(vals, line)
				}
			}
			return carapace.ActionValues(vals...).Tag("installed versions")
		})
	})
}

// ActionAllVersions completes all available versions for a given plugin
//
//	3.2.2
//	3.3.0
func ActionAllVersions(plugin string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		name := plugin
		if name == "" && len(c.Args) > 0 {
			name = c.Args[0]
		}
		if name == "" {
			return carapace.ActionMessage("missing plugin name")
		}
		return carapace.ActionExecCommand("asdf", "list", "all", name)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			var vals []string
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line != "" {
					vals = append(vals, line)
				}
			}
			return carapace.ActionValues(vals...).Tag("available versions")
		})
	})
}

// ActionLatestVersions completes the latest version syntax for a plugin
//
//	latest
//	latest:3
func ActionLatestVersions(plugin string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		name := plugin
		if name == "" && len(c.Args) > 0 {
			name = c.Args[0]
		}
		if name == "" {
			return carapace.ActionMessage("missing plugin name")
		}
		return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("latest", "latest:")
			case 1:
				if c.Parts[0] == "latest" {
					return ActionAllVersions(name)
				}
				return carapace.ActionValues()
			default:
				return carapace.ActionValues()
			}
		})
	})
}

// ActionInstallVersions completes versions accepted by install
//
//	latest
//	latest:3
//	3.2.2
func ActionInstallVersions(plugin string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		name := plugin
		if name == "" && len(c.Args) > 0 {
			name = c.Args[0]
		}
		if name == "" {
			return carapace.ActionMessage("missing plugin name")
		}
		return carapace.Batch(
			ActionAllVersions(name),
			ActionLatestVersions(name),
		).ToA()
	})
}

// ActionSetVersions completes versions accepted by set
//
//	latest
//	latest:3
//	system
//	3.2.2
func ActionSetVersions(plugin string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		name := plugin
		if name == "" && len(c.Args) > 0 {
			name = c.Args[0]
		}
		if name == "" {
			return carapace.ActionMessage("missing plugin name")
		}
		return carapace.Batch(
			ActionInstalledVersions(name),
			ActionLatestVersions(name),
			carapace.ActionValues("system"),
		).ToA()
	})
}
