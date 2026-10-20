// package mise contains mise related actions
package mise

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

// parseDescribed parses `name:description` lines as printed by `--complete` flags.
// A colon within the name is escaped as `\:`.
func parseDescribed(output []byte) []string {
	vals := make([]string, 0)
	for _, line := range strings.Split(string(output), "\n") {
		line = strings.TrimSuffix(line, "\r")
		if line == "" {
			continue
		}

		var name strings.Builder
		desc := ""
		for i := 0; i < len(line); i++ {
			switch {
			case line[i] == '\\' && i+1 < len(line) && line[i+1] == ':':
				name.WriteByte(':')
				i++
			case line[i] == ':':
				desc = line[i+1:]
				i = len(line)
			default:
				name.WriteByte(line[i])
			}
		}
		vals = append(vals, name.String(), desc)
	}
	return vals
}

func parseLines(output []byte) []string {
	vals := make([]string, 0)
	for _, line := range strings.Split(string(output), "\n") {
		if line = strings.TrimSpace(line); line != "" {
			vals = append(vals, line)
		}
	}
	return vals
}

type installedVersion struct {
	Version string `json:"version"`
}

// ActionTasks completes tasks
//
//	build (Build it)
//	test:unit
func ActionTasks() carapace.Action {
	return carapace.ActionExecCommand("mise", "tasks", "ls", "--complete")(func(output []byte) carapace.Action {
		return carapace.ActionValuesDescribed(parseDescribed(output)...)
	}).Tag("tasks")
}

// ActionSettings completes settings
//
//	activate_aggressive (Keep mise's tool directories ahead of `PATH` entries added after `mise activate`.)
//	age.key_file (Age private key file used to decrypt and as the default recipient when encrypting.)
func ActionSettings() carapace.Action {
	return carapace.ActionExecCommand("mise", "settings", "--complete")(func(output []byte) carapace.Action {
		return carapace.ActionValuesDescribed(parseDescribed(output)...)
	}).Tag("settings").Cache(24 * time.Hour)
}

// ActionEnvironmentKeys completes keys of configured environment variables
//
//	NODE_ENV
//	RAILS_ENV
func ActionEnvironmentKeys() carapace.Action {
	return carapace.ActionExecCommand("mise", "set", "--complete")(func(output []byte) carapace.Action {
		return carapace.ActionValues(parseLines(output)...)
	}).Tag("environment variables")
}

// ActionSecrets completes secrets
//
//	API_TOKEN
//	DATABASE_URL
func ActionSecrets() carapace.Action {
	return carapace.ActionExecCommand("mise", "secrets", "ls", "--complete")(func(output []byte) carapace.Action {
		return carapace.ActionValuesDescribed(parseDescribed(output)...)
	}).Tag("secrets")
}

// ActionPlugins completes installed plugins (core and user)
//
//	bun
//	node
func ActionPlugins() carapace.Action {
	return carapace.ActionExecCommand("mise", "plugins", "--core", "--user")(func(output []byte) carapace.Action {
		return carapace.ActionValues(parseLines(output)...)
	}).Tag("plugins")
}

// ActionRemotePlugins completes plugins available for installation
//
//	1password
//	act
func ActionRemotePlugins() carapace.Action {
	return carapace.ActionExecCommand("mise", "plugins", "--all")(func(output []byte) carapace.Action {
		return carapace.ActionValues(parseLines(output)...)
	}).Tag("remote plugins").Cache(24 * time.Hour)
}

// ActionBackends completes backends
//
//	aqua
//	npm
func ActionBackends() carapace.Action {
	return carapace.ActionExecCommand("mise", "backends")(func(output []byte) carapace.Action {
		return carapace.ActionValues(parseLines(output)...)
	}).Tag("backends").Cache(24 * time.Hour)
}

// ActionTools completes tools from the mise registry
//
//	1password (Password manager developed by AgileBits Inc)
//	7z (7-Zip is a file archiver with a high compression ratio)
func ActionTools() carapace.Action {
	return carapace.ActionExecCommand("mise", "search", "--complete")(func(output []byte) carapace.Action {
		return carapace.ActionValuesDescribed(parseDescribed(output)...)
	}).Tag("tools").Cache(24 * time.Hour)
}

// ActionInstalledTools completes installed tools
//
//	node
//	python
func ActionInstalledTools() carapace.Action {
	return carapace.ActionExecCommand("mise", "ls", "--installed", "--json")(func(output []byte) carapace.Action {
		installed := make(map[string][]installedVersion)
		if err := json.Unmarshal(output, &installed); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0, len(installed))
		for tool := range installed {
			vals = append(vals, tool)
		}
		return carapace.ActionValues(vals...)
	}).Tag("installed tools")
}

// ActionVersions completes versions available for a tool
//
//	24.0.0
//	24.1.0
func ActionVersions(tool string) carapace.Action {
	return carapace.ActionExecCommand("mise", "ls-remote", tool)(func(output []byte) carapace.Action {
		return carapace.ActionValues(parseLines(output)...)
	}).Tag("versions").Cache(1*time.Hour, key.String(tool))
}

// ActionInstalledVersions completes installed versions of a tool
//
//	20.20.2
//	24.0.0
func ActionInstalledVersions(tool string) carapace.Action {
	return carapace.ActionExecCommand("mise", "ls", "--installed", "--json", tool)(func(output []byte) carapace.Action {
		var installed []installedVersion
		if err := json.Unmarshal(output, &installed); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0, len(installed))
		for _, v := range installed {
			vals = append(vals, v.Version)
		}
		return carapace.ActionValues(vals...)
	}).Tag("installed versions")
}

// ActionToolVersions completes tools with their available versions
//
//	node@24.0.0
//	python@3.13
func ActionToolVersions() carapace.Action {
	return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionTools().Suffix("@")
		default:
			return ActionVersions(c.Parts[0])
		}
	})
}

// ActionInstalledToolVersions completes installed tools with their installed versions
//
//	node@20.20.2
//	python@3.13.1
func ActionInstalledToolVersions() carapace.Action {
	return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionInstalledTools().Suffix("@")
		default:
			return ActionInstalledVersions(c.Parts[0])
		}
	})
}

// ActionToolAliases completes version aliases of a tool
//
//	lts (24)
//	lts-argon (4)
func ActionToolAliases(tool string) carapace.Action {
	return carapace.ActionExecCommand("mise", "tool-alias", "ls", tool)(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for _, line := range parseLines(output) {
			if fields := strings.Fields(line); len(fields) >= 3 {
				vals = append(vals, fields[1], fields[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("tool aliases").Cache(1*time.Hour, key.String(tool))
}

// ActionBinNames completes executables provided by tools
//
//	node
//	npm
func ActionBinNames() carapace.Action {
	return carapace.ActionExecCommand("mise", "which", "--complete")(func(output []byte) carapace.Action {
		return carapace.ActionValues(parseLines(output)...)
	}).Tag("executables")
}
