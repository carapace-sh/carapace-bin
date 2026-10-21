package tsh

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"go.yaml.in/yaml/v3"
)

// ActionProfiles completes tsh profiles
//
//	proxy.example.com (alice@root)
//	staging.example.com (bob@staging)
func ActionProfiles() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		dir := os.Getenv("TELEPORT_HOME")
		if dir == "" {
			dir = "~/.tsh"
		}

		dir, err := c.Abs(dir)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		current := ""
		if content, err := os.ReadFile(filepath.Join(dir, "current-profile")); err == nil {
			current = strings.TrimSpace(string(content))
		}

		profiles := make(map[string]bool)
		vals := make([]string, 0)
		if entries, err := os.ReadDir(dir); err == nil {
			for _, entry := range entries {
				name, ok := strings.CutSuffix(entry.Name(), ".yaml")
				if !ok || entry.IsDir() {
					continue
				}

				profiles[name] = true
				vals = append(vals, name, profileDescription(filepath.Join(dir, entry.Name()), name == current))
			}
		}

		if entries, err := os.ReadDir(filepath.Join(dir, "keys")); err == nil {
			for _, entry := range entries {
				if !entry.IsDir() || profiles[entry.Name()] {
					continue
				}
				vals = append(vals, entry.Name(), profileDescription("", entry.Name() == current))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func profileDescription(path string, current bool) string {
	desc := ""
	if path != "" {
		if content, err := os.ReadFile(path); err == nil {
			var profile struct {
				Username string `yaml:"user"`
				Cluster  string `yaml:"cluster"`
				Proxy    string `yaml:"web_proxy_addr"`
			}
			if err := yaml.Unmarshal(content, &profile); err == nil {
				if profile.Username != "" && profile.Cluster != "" {
					desc = profile.Username + "@" + profile.Cluster
				} else if profile.Proxy != "" {
					desc = profile.Proxy
				}
			}
		}
	}

	if current {
		if desc != "" {
			desc += ", "
		}
		desc += "current"
	}
	return desc
}
