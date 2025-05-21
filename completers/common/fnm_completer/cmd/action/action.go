package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionLogLevel() carapace.Action {
	return carapace.ActionValues(
		"quiet",
		"info",
		"error",
	)
}

func ActionVersionFileStrategy() carapace.Action {
	return carapace.ActionValuesDescribed(
		"local", "Use the local version of Node defined within the current directory",
		"recursive", "Use the version of Node defined within the current dirctory and all parent directories",
	)
}

func ActionResolveEngines() carapace.Action {
	return carapace.ActionValues("true", "false")
}

func ActionLocalVersions() carapace.Action {
	return withLocalFnmVersions(func(versions []fnmVersion) carapace.Action {
		var versionsStr []string
		for _, version := range versions {
			versionsStr = append(versionsStr, version.version)
		}

		return carapace.ActionValues(versionsStr...)
	})
}

func ActionAliases() carapace.Action {
	return withLocalFnmVersions(func(versions []fnmVersion) carapace.Action {
		var aliases []string
		for _, version := range versions {
			for _, alias := range version.aliases {
				aliases = append(aliases, alias)
			}
		}

		return carapace.ActionValues(aliases...)
	})
}

type fnmVersion struct {
	version string
	aliases []string
}

func withLocalFnmVersions(callback func([]fnmVersion) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("fnm", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		var versions []fnmVersion

		for _, line := range lines {
			// line format: https://github.com/Schniz/fnm/blob/master/src/commands/ls_local.rs#L38

			parts := strings.Split(line, " ")
			if len(parts) < 2 {
				continue
			}

			version := fnmVersion{
				version: parts[1],
			}

			if len(parts) > 2 {
				for _, alias := range parts[2:] {
					version.aliases = append(version.aliases, strings.Trim(alias, ","))
				}
			}
		}

		return callback(versions)
	})
}
