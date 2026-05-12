package openrc

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/carapace-sh/carapace"
)

// ActionServices completes OpenRC service scripts from system and user init directories.
func ActionServices() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionValues(readNames(serviceDirs(), false)...)
	})
}

// ActionRunlevels completes OpenRC runlevels from system and user runlevel directories.
func ActionRunlevels() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		static := map[string]string{
			"boot":     "early boot services",
			"default":  "default runlevel",
			"reboot":   "shutdown and reboot the host",
			"shutdown": "shutdown and halt the host",
			"single":   "single-user runlevel",
			"sysinit":  "system initialization services",
		}

		described := make([]string, 0, len(static)*2)
		for _, name := range sortedKeys(static) {
			described = append(described, name, static[name])
		}

		var dynamic []string
		for _, name := range readNames(runlevelDirs(), true) {
			if _, ok := static[name]; !ok {
				dynamic = append(dynamic, name)
			}
		}

		return carapace.Batch(
			carapace.ActionValuesDescribed(described...),
			carapace.ActionValues(dynamic...),
		).ToA()
	})
}

// ActionServiceCommands completes common commands supported by OpenRC service scripts.
func ActionServiceCommands() carapace.Action {
	return carapace.ActionValuesDescribed(
		"start", "start the service",
		"stop", "stop the service",
		"restart", "restart the service",
		"status", "show service status",
		"zap", "forget the service's started state",
		"describe", "show service description",
		"needsme", "show services that need this service",
		"ineed", "show services this service needs",
		"iuse", "show services this service uses",
		"broken", "show broken dependencies",
		"provide", "show virtual services provided",
		"cgroup_cleanup", "clean cgroups for a stopped service",
	)
}

// ActionServiceStates completes service states accepted by rc-status --in-state.
func ActionServiceStates() carapace.Action {
	return carapace.ActionValues(
		"stopped",
		"started",
		"stopping",
		"starting",
		"inactive",
		"hotplugged",
		"failed",
		"scheduled",
		"crashed",
	)
}

func serviceDirs() []string {
	dirs := []string{
		"/etc/init.d",
		"/usr/local/etc/init.d",
		"/etc/user/init.d",
	}
	if configHome := xdgConfigHome(); configHome != "" {
		dirs = append(dirs, filepath.Join(configHome, "rc", "init.d"))
	}
	return dirs
}

func runlevelDirs() []string {
	dirs := []string{"/etc/runlevels"}
	if configHome := xdgConfigHome(); configHome != "" {
		dirs = append(dirs, filepath.Join(configHome, "rc", "runlevels"))
	}
	return dirs
}

func xdgConfigHome() string {
	if configHome := os.Getenv("XDG_CONFIG_HOME"); configHome != "" {
		return configHome
	}
	if home, err := os.UserHomeDir(); err == nil {
		return filepath.Join(home, ".config")
	}
	return ""
}

func readNames(dirs []string, dirsOnly bool) []string {
	seen := make(map[string]bool)
	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if entry.Name() == "" || entry.Name()[0] == '.' {
				continue
			}
			if dirsOnly && !entry.IsDir() {
				continue
			}
			seen[entry.Name()] = true
		}
	}
	return sortedKeys(seen)
}

func sortedKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
