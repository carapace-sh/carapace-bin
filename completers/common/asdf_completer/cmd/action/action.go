package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPlugins completes installed asdf plugins
//
//	ruby
//	nodejs
//	python
func ActionPlugins() carapace.Action {
	return carapace.ActionExecCommand("asdf", "plugin", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		var vals []string
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" {
				name := strings.Fields(line)[0]
				vals = append(vals, name)
			}
		}
		return carapace.ActionValues(vals...).Tag("plugins")
	})
}

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
				line = strings.TrimSpace(line)
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
