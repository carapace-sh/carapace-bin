package action

import (
	"regexp"
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

func ActionInstalledVersions() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return ActionLocalVersions().Invoke(c).Merge(
			ActionAliases().Invoke(c),
		).ToA()
	})
}

type fnmVersion struct {
	version string
	aliases []string
}

var (
	versionRe = regexp.MustCompile(`\*\s(?P<version>[^\s]+)\s?(?P<aliases>.*)`)
)

func withLocalFnmVersions(callback func([]fnmVersion) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("fnm", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		var versions []fnmVersion

		for _, line := range lines {
			// line format: https://github.com/Schniz/fnm/blob/master/src/commands/ls_local.rs#L38

			matches := versionRe.FindStringSubmatch(line)

			if len(matches) < 3 {
				continue
			}

			version := fnmVersion{
				version: matches[1],
			}

			for _, alias := range strings.Split(matches[2], ", ") {
				value := strings.TrimSpace(alias)
				if value != "" {
					version.aliases = append(version.aliases, value)
				}
			}

			versions = append(versions, version)
		}

		return callback(versions)
	})
}
