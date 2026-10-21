package action

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
)

func dataDir() string {
	dir := os.Getenv("ASDF_DATA_DIR")
	if dir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return ""
		}
		return filepath.Join(home, ".asdf")
	}

	if dir == "~" || strings.HasPrefix(dir, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return ""
		}
		dir = filepath.Join(home, dir[1:])
	}
	return filepath.Clean(dir)
}

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

// ActionAvailablePlugins completes plugins available in the plugin repository
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

// ActionShims completes shims
//
//	ruby
//	node
//	python
func ActionShims() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		entries, err := os.ReadDir(filepath.Join(dataDir(), "shims"))
		if err != nil {
			return carapace.ActionValues()
		}
		var vals []string
		for _, entry := range entries {
			if !entry.IsDir() {
				vals = append(vals, entry.Name())
			}
		}
		return carapace.ActionValues(vals...).Tag("shims")
	})
}
