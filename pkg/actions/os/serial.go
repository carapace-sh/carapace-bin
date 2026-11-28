// package os contains operating system related actions
package os

import (
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/carapace-sh/carapace"
)

var comPortRegex = regexp.MustCompile(`\((COM\d+)\)\s*$`)

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
