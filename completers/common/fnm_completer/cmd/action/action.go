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
			aliases = append(aliases, version.aliases...)
		}

		return carapace.ActionValues(aliases...)
	})
}

func ActionInstalledVersions() carapace.Action {
	return carapace.ActionCallback(func(carapace.Context) carapace.Action {
		return carapace.Batch(
			ActionLocalVersions(),
			ActionAliases(),
		).ToA()
	})
}

type fnmVersion struct {
	version string
	aliases []string
}

func withLocalFnmVersions(callback func([]fnmVersion) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("fnm", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		// line format: https://github.com/Schniz/fnm/blob/master/src/commands/ls_local.rs#L38
		versionRe := regexp.MustCompile(`\*\s(?P<version>[^\s]+)\s?(?P<aliases>.*)`)

		var versions []fnmVersion
		for _, line := range lines {
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
