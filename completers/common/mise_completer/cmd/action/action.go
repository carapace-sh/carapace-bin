package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionTools() carapace.Action {
	return carapace.ActionExecCommand("mise", "ls", "--installed")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 0 {
				vals = append(vals, fields[0])
			}
		}
		return carapace.ActionValues(vals...).UniqueList(",")
	})
}

func ActionInstalledToolVersions() carapace.Action {
	return carapace.ActionExecCommand("mise", "ls", "--installed")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) >= 2 {
				vals = append(vals, fields[0]+"@"+fields[1], fields[0])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionTasks() carapace.Action {
	return carapace.ActionExecCommand("mise", "tasks", "ls")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 0 {
				desc := ""
				if len(fields) > 1 {
					desc = strings.Join(fields[1:], " ")
				}
				vals = append(vals, fields[0], desc)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionPlugins() carapace.Action {
	return carapace.ActionExecCommand("mise", "plugins", "ls")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 0 {
				vals = append(vals, fields[0])
			}
		}
		return carapace.ActionValues(vals...)
	})
}

func ActionSettings() carapace.Action {
	return carapace.ActionExecCommand("mise", "settings", "--all")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 0 {
				desc := ""
				if len(fields) > 1 {
					desc = fields[1]
				}
				vals = append(vals, fields[0], desc)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionAliases() carapace.Action {
	return carapace.ActionExecCommand("mise", "alias", "ls")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 0 {
				vals = append(vals, fields[0])
			}
		}
		return carapace.ActionValues(vals...)
	})
}

func ActionBackends() carapace.Action {
	return carapace.ActionValues(
		"aqua",
		"asdf",
		"cargo",
		"conda",
		"dotnet",
		"forgejo",
		"gem",
		"github",
		"gitlab",
		"go",
		"http",
		"npm",
		"packslip",
		"pipx",
		"pkgx",
		"spm",
		"ubi",
		"vfox",
	)
}

func ActionShells() carapace.Action {
	return carapace.ActionValues("bash", "elvish", "fish", "nu", "powershell", "pwsh", "xonsh", "zsh")
}
