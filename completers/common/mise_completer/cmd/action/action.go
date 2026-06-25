package action

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

// ActionShells completes shells supported by mise script/completion commands.
func ActionShells() carapace.Action {
	return carapace.ActionValues("bash", "elvish", "fish", "nu", "powershell", "pwsh", "xonsh", "zsh")
}

// ActionBackends completes mise backends.
func ActionBackends() carapace.Action {
	return carapace.ActionValues("aqua", "asdf", "cargo", "core", "dotnet", "gem", "go", "npm", "pipx", "spm", "ubi", "vfox")
}

// ActionTasks completes project tasks.
func ActionTasks() carapace.Action {
	return carapace.ActionExecCommand("mise", "tasks", "ls", "--name-only", "--hidden")(func(output []byte) carapace.Action {
		return actionLines(output).Tag("tasks")
	})
}

// ActionRegistryTools completes tools known by the mise registry.
func ActionRegistryTools() carapace.Action {
	return carapace.ActionExecCommand("mise", "registry", "--hide-aliased")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for _, line := range strings.Split(string(output), "\n") {
			fields := strings.Fields(line)
			if len(fields) == 0 {
				continue
			}
			vals = append(vals, fields[0])
			if len(fields) > 1 {
				vals = append(vals, strings.Join(fields[1:], " "))
			} else {
				vals = append(vals, "")
			}
		}
		return carapace.ActionValuesDescribed(vals...).Tag("registry tools")
	}).Cache(24 * time.Hour)
}

// ActionToolVersions completes TOOL@VERSION values for install/use-like commands.
func ActionToolVersions() carapace.Action {
	return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionRegistryTools().NoSpace()
		case 1:
			return ActionRemoteVersions(c.Parts[0])
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionRemoteVersions completes versions available for a registry tool.
func ActionRemoteVersions(tool string) carapace.Action {
	return carapace.ActionExecCommand("mise", "ls-remote", tool)(func(output []byte) carapace.Action {
		return actionLines(output).Tag("versions")
	}).Cache(24*time.Hour, key.String(tool))
}

// ActionInstalledToolVersions completes installed/active TOOL@VERSION values.
func ActionInstalledToolVersions() carapace.Action {
	return carapace.ActionExecCommand("mise", "ls", "--json")(func(output []byte) carapace.Action {
		var tools map[string][]struct {
			Version string `json:"version"`
		}
		if err := json.Unmarshal(output, &tools); err != nil {
			return carapace.ActionValues()
		}

		vals := make([]string, 0)
		for tool, versions := range tools {
			if len(versions) == 0 {
				vals = append(vals, tool)
				continue
			}
			for _, version := range versions {
				if version.Version == "" {
					vals = append(vals, tool)
					continue
				}
				vals = append(vals, tool+"@"+version.Version)
			}
		}
		return carapace.ActionValues(vals...).Tag("installed tools")
	})
}

// ActionPlugins completes installed plugins.
func ActionPlugins() carapace.Action {
	return carapace.ActionExecCommand("mise", "plugins", "ls")(func(output []byte) carapace.Action {
		return actionFirstFields(output).Tag("plugins")
	})
}

// ActionRemotePlugins completes remotely available plugins.
func ActionRemotePlugins() carapace.Action {
	return carapace.ActionExecCommand("mise", "plugins", "ls-remote", "--only-names")(func(output []byte) carapace.Action {
		return actionFirstFields(output).Tag("remote plugins")
	}).Cache(24 * time.Hour)
}

// ActionConfigKeys completes keys from active mise config.
func ActionConfigKeys() carapace.Action {
	return carapace.ActionExecCommand("mise", "config", "ls", "--no-header")(func(output []byte) carapace.Action {
		return actionFirstFields(output).Tag("config keys")
	})
}

// ActionSettings completes setting names from mise settings.
func ActionSettings() carapace.Action {
	return carapace.ActionExecCommand("mise", "settings", "ls")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for _, line := range strings.Split(string(output), "\n") {
			key, _, ok := strings.Cut(strings.TrimSpace(line), " ")
			if !ok {
				key, _, ok = strings.Cut(strings.TrimSpace(line), "=")
			}
			if ok && key != "" {
				vals = append(vals, key)
			}
		}
		return carapace.ActionValues(vals...).Tag("settings")
	})
}

// ActionShellAliases completes configured shell aliases.
func ActionShellAliases() carapace.Action {
	return carapace.ActionExecCommand("mise", "shell-alias", "ls", "--no-header")(func(output []byte) carapace.Action {
		return actionFirstFields(output).Tag("shell aliases")
	})
}

// ActionHosts completes common token hosts.
func ActionHosts() carapace.Action {
	return carapace.ActionValues("github.com", "gitlab.com", "codeberg.org")
}

func actionLines(output []byte) carapace.Action {
	vals := make([]string, 0)
	for _, line := range strings.Split(string(output), "\n") {
		line = strings.TrimSpace(strings.TrimPrefix(line, "*"))
		if line != "" {
			vals = append(vals, line)
		}
	}
	return carapace.ActionValues(vals...)
}

func actionFirstFields(output []byte) carapace.Action {
	vals := make([]string, 0)
	for _, line := range strings.Split(string(output), "\n") {
		fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "*"))
		if len(fields) > 0 {
			vals = append(vals, fields[0])
		}
	}
	return carapace.ActionValues(vals...)
}
