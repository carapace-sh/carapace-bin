// package espidf contains ESP-IDF related actions
package espidf

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"go.yaml.in/yaml/v3"
)

// ActionTargets completes ESP-IDF chip targets
//
//	esp32
//	esp32s3
func ActionTargets() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		currentTarget := readProjectTarget(c.Dir)
		return carapace.ActionValues(
			"esp32",
			"esp32c2",
			"esp32c3",
			"esp32c5",
			"esp32c6",
			"esp32c61",
			"esp32h2",
			"esp32h4",
			"esp32h21",
			"esp32p4",
			"esp32s2",
			"esp32s3",
			"esp32s31",
			"linux",
		).StyleF(func(s string, sc style.Context) string {
			if currentTarget != "" && s == currentTarget {
				return style.Green
			}
			return style.Default
		})
	})
}

func readProjectTarget(dir string) string {
	data, err := os.ReadFile(filepath.Join(dir, "build", "project_description.json"))
	if err != nil {
		return ""
	}
	var desc struct {
		Target string `json:"target"`
	}
	if err := json.Unmarshal(data, &desc); err != nil {
		return ""
	}
	return desc.Target
}

// ActionPresets completes CMake configure presets from CMakePresets.json and CMakeUserPresets.json
//
//	default (Default configuration)
func ActionPresets() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0)
		seen := make(map[string]bool)
		for _, filename := range []string{"CMakePresets.json", "CMakeUserPresets.json"} {
			data, err := os.ReadFile(filepath.Join(c.Dir, filename))
			if err != nil {
				continue
			}
			var presets struct {
				ConfigurePresets []struct {
					Name        string `json:"name"`
					DisplayName string `json:"displayName"`
					Description string `json:"description"`
					Hidden      bool   `json:"hidden"`
				} `json:"configurePresets"`
			}
			if err := json.Unmarshal(data, &presets); err != nil {
				continue
			}
			for _, p := range presets.ConfigurePresets {
				if p.Hidden || p.Name == "" || seen[p.Name] {
					continue
				}
				seen[p.Name] = true
				desc := p.DisplayName
				if desc == "" {
					desc = p.Description
				}
				vals = append(vals, p.Name, desc)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionDependencies completes project component dependencies from dependencies.lock or idf_component.yml
//
//	espressif/mdns
func ActionDependencies() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0)
		seen := make(map[string]bool)

		for _, relPath := range []string{
			"dependencies.lock",
			filepath.Join("main", "idf_component.yml"),
			"idf_component.yml",
		} {
			data, err := os.ReadFile(filepath.Join(c.Dir, relPath))
			if err != nil {
				continue
			}
			var doc struct {
				Dependencies map[string]any `yaml:"dependencies"`
			}
			if err := yaml.Unmarshal(data, &doc); err != nil {
				continue
			}
			for name := range doc.Dependencies {
				if name == "" || name == "idf" || seen[name] {
					continue
				}
				seen[name] = true
				vals = append(vals, name)
			}
		}
		return carapace.ActionValues(vals...)
	})
}
