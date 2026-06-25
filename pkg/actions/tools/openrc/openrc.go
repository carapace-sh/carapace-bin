// Package openrc provides completions for the OpenRC init system.
package openrc

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionServices completes OpenRC service script names from system init.d directories.
//
//	sshd (/etc/init.d/sshd)
//	networking (/etc/init.d/networking)
func ActionServices() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0)
		for _, name := range readEntries(serviceDirs(), false) {
			vals = append(vals, name, describeService(name))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionRunlevels completes OpenRC runlevels, including known defaults and
// any runlevels discovered in /etc/runlevels.
func ActionRunlevels() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		defaults := map[string]string{
			"boot":      "early boot services",
			"default":   "default runlevel",
			"nonetwork": "non-network runlevel",
			"reboot":    "shutdown and reboot the host",
			"shutdown":  "shutdown and halt the host",
			"single":    "single-user runlevel",
			"sysinit":   "system initialization services",
		}

		seen := make(map[string]bool)
		vals := make([]string, 0)
		for _, name := range sortedKeys(defaults) {
			seen[name] = true
			vals = append(vals, name, defaults[name], style.Blue)
		}
		for _, name := range readEntries(runlevelDirs(), true) {
			if seen[name] {
				continue
			}
			vals = append(vals, name, "user-defined runlevel", style.Default)
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}

// ActionServiceCommands completes the common verbs accepted by OpenRC
// service scripts.
func ActionServiceCommands() carapace.Action {
	return carapace.ActionValuesDescribed(
		"start", "start the service",
		"stop", "stop the service",
		"restart", "restart the service",
		"status", "show service status",
		"pause", "stop the service but do not mark it as stopped",
		"zap", "forget the service's started state",
		"describe", "show service description",
		"depend", "show service dependencies",
		"needsme", "show services that need this service",
		"ineed", "show services this service needs",
		"iuse", "show services this service uses",
		"usesme", "show services that use this service",
		"broken", "show broken dependencies",
		"provide", "show virtual services provided",
		"cgroup_cleanup", "clean cgroups for a stopped service",
	)
}

// ActionServiceStates completes service state values accepted by
// `rc-status --in-state`.
func ActionServiceStates() carapace.Action {
	return carapace.ActionValuesDescribed(
		"started", "service is running",
		"starting", "service is starting",
		"stopped", "service is stopped",
		"stopping", "service is stopping",
		"inactive", "service is inactive",
		"hotplugged", "service was hotplugged",
		"failed", "service failed to start",
		"scheduled", "service is scheduled to start",
		"crashed", "service crashed",
	)
}

// serviceDirs returns the directories searched for service scripts.
func serviceDirs() []string {
	dirs := []string{
		"/etc/init.d",
		"/usr/local/etc/init.d",
	}
	if configHome := xdgConfigHome(); configHome != "" {
		dirs = append(dirs, filepath.Join(configHome, "rc", "init.d"))
	}
	return dirs
}

// runlevelDirs returns the directories searched for runlevel definitions.
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

// readEntries returns sorted, de-duplicated entry names from the given
// directories. When dirsOnly is true only sub-directories are returned;
// otherwise files and symlinks are returned.
func readEntries(dirs []string, dirsOnly bool) []string {
	seen := make(map[string]struct{})
	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			name := entry.Name()
			if name == "" || name[0] == '.' {
				continue
			}
			if dirsOnly && !entry.IsDir() {
				continue
			}
			seen[name] = struct{}{}
		}
	}
	return sortedKeys(seen)
}

// describeService returns a short description for an OpenRC service by
// reading the `description=` line of its init script if present.
func describeService(name string) string {
	for _, dir := range serviceDirs() {
		path := filepath.Join(dir, name)
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			line = strings.TrimSpace(line)
			if !strings.HasPrefix(line, "description=") {
				continue
			}
			value := strings.TrimPrefix(line, "description=")
			value = strings.Trim(value, "\"'")
			if value != "" {
				return value
			}
		}
	}
	return ""
}

func sortedKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
