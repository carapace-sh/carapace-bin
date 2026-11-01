package asdf

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
)

// ActionInstalledPlugins completes installed plugins
//
//	ruby
//	nodejs
//	python
func ActionInstalledPlugins() carapace.Action {
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

// ActionAvailablePlugins completes plugins available for installation
//
//	ruby
//	nodejs
//	python
func ActionAvailablePlugins() carapace.Action {
	return carapace.ActionExecCommand("asdf", "plugin", "list", "all")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		var vals []string
		for _, line := range lines {
			fields := strings.Fields(line)
			if len(fields) < 2 || strings.HasPrefix(fields[1], "*") {
				continue
			}
			vals = append(vals, fields[0])
		}
		return carapace.ActionValues(vals...).Tag("available plugins")
	})
}

// ActionPluginURL completes the repository url of a plugin
//
//	https://github.com/asdf-vm/asdf-ruby.git
func ActionPluginURL(plugin string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		name := plugin
		if name == "" && len(c.Args) > 0 {
			name = c.Args[0]
		}
		if name == "" {
			return carapace.ActionMessage("missing plugin name")
		}
		return carapace.ActionExecCommand("asdf", "plugin", "list", "all")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			for _, line := range lines {
				fields := strings.Fields(line)
				if len(fields) >= 2 && fields[0] == name {
					return carapace.ActionValues(strings.TrimPrefix(fields[1], "*"))
				}
			}
			return carapace.ActionValues()
		})
	})
}

// ActionPluginRefs completes git refs of a plugin
//
//	master
//	v0.1.0
func ActionPluginRefs(plugin string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		name := plugin
		if name == "" && len(c.Args) > 0 {
			name = c.Args[0]
		}
		if name == "" {
			return carapace.ActionMessage("missing plugin name")
		}
		return carapace.Batch(
			git.ActionRefs(git.RefOption{
				LocalBranches:  true,
				RemoteBranches: true,
				Heads:          true,
				Tags:           true,
			}),
			git.ActionRecentCommits(10),
		).ToA().Chdir(filepath.Join(dataDir(), "plugins", name))
	})
}

// ActionExtensionCommands completes extension commands of a plugin
//
//	foo
//	bar
func ActionExtensionCommands(plugin string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		name := plugin
		if name == "" && len(c.Args) > 0 {
			name = c.Args[0]
		}
		if name == "" {
			return carapace.ActionMessage("missing plugin name")
		}
		entries, err := os.ReadDir(filepath.Join(dataDir(), "plugins", name, "lib", "commands"))
		if err != nil {
			return carapace.ActionValues()
		}
		var vals []string
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			if command, ok := strings.CutPrefix(entry.Name(), "command-"); ok {
				vals = append(vals, command)
			}
		}
		return carapace.ActionValues(vals...).Tag("extension commands")
	})
}
