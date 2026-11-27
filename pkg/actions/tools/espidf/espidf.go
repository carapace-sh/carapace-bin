// package espidf contains ESP-IDF related actions
package espidf

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"go.yaml.in/yaml/v3"
)

var comPortRegex = regexp.MustCompile(`\((COM\d+)\)\s*$`)

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

// ActionBaudRates completes common serial baud rates
//
//	115200
//	460800
func ActionBaudRates() carapace.Action {
	return carapace.ActionValues(
		"9600",
		"57600",
		"115200",
		"230400",
		"460800",
		"921600",
		"1500000",
		"2000000",
	)
}

// ActionSerialPorts completes available serial ports
//
//	COM5 (USB-SERIAL CH340)
//	/dev/ttyUSB0
func ActionSerialPorts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if runtime.GOOS == "windows" {
			return carapace.ActionExecCommand("reg", "query", `HKLM\HARDWARE\DEVICEMAP\SERIALCOMM`)(func(output []byte) carapace.Action {
				friendlyNames := readWindowsUSBSerialNames()
				lines := strings.Split(string(output), "\n")
				vals := make([]string, 0)
				for _, line := range lines {
					fields := strings.Fields(line)
					if len(fields) >= 3 && fields[1] == "REG_SZ" {
						port := fields[2]
						desc := fields[0]
						if name, ok := friendlyNames[port]; ok {
							desc = name
						}
						vals = append(vals, port, desc)
					}
				}
				return carapace.ActionValuesDescribed(vals...)
			})
		}

		patterns := []string{
			"/dev/cu.*",
			"/dev/tty.*",
			"/dev/ttyACM*",
			"/dev/ttyS*",
			"/dev/ttyUSB*",
		}
		vals := make([]string, 0)
		for _, pattern := range patterns {
			if matches, err := filepath.Glob(pattern); err == nil {
				vals = append(vals, matches...)
			}
		}
		return carapace.ActionValues(vals...)
	})
}

func readWindowsUSBSerialNames() map[string]string {
	result := make(map[string]string)
	out, err := exec.Command("reg", "query", `HKLM\SYSTEM\CurrentControlSet\Enum\USB`, "/s", "/v", "FriendlyName").Output()
	if err != nil {
		return result
	}
	for _, line := range strings.Split(string(out), "\n") {
		idx := strings.Index(line, "REG_SZ")
		if idx == -1 {
			continue
		}
		val := strings.TrimSpace(line[idx+len("REG_SZ"):])
		if m := comPortRegex.FindStringSubmatch(val); len(m) == 2 {
			port := m[1]
			name := strings.TrimSpace(comPortRegex.ReplaceAllString(val, ""))
			if name != "" {
				result[port] = name
			}
		}
	}
	return result
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
