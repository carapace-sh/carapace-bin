// package os contains operating system related actions
package os

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionSerialPorts completes available serial ports
//
//	COM5 (USB Serial Device)
//	/dev/ttyUSB0 (usb-serial)
func ActionSerialPorts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch runtime.GOOS {
		case "windows":
			return actionSerialPortsWindows()
		default:
			return actionSerialPortsUnix()
		}
	}).Tag("serial ports")
}

func actionSerialPortsWindows() carapace.Action {
	return carapace.Batch(
		actionSerialPortsWindowsPnp(),
		actionSerialPortsWindowsSerialComm(),
	).ToA()
}

func actionSerialPortsWindowsPnp() carapace.Action {
	return carapace.ActionExecCommand("powershell", "-NoProfile", "-Command", `Get-CimInstance Win32_PnPEntity -Filter "Name LIKE '%(COM%)'" -Property Name | Select-Object Name | ConvertTo-Json -Compress`)(func(output []byte) carapace.Action {
		var devices []struct {
			Name string
		}
		if err := json.Unmarshal(output, &devices); err != nil {
			var device struct {
				Name string
			}
			if err := json.Unmarshal(output, &device); err != nil {
				return carapace.ActionValues()
			}
			devices = []struct{ Name string }{device}
		}

		comPortRegex := regexp.MustCompile(`\((COM\d+)\)\s*$`)
		vals := make([]string, 0)
		for _, device := range devices {
			m := comPortRegex.FindStringSubmatch(device.Name)
			if len(m) != 2 {
				continue
			}
			port := m[1]
			desc := strings.TrimSpace(comPortRegex.ReplaceAllString(device.Name, ""))
			if desc == "" {
				desc = "serial port"
			}
			vals = append(vals, port, desc)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func actionSerialPortsWindowsSerialComm() carapace.Action {
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

func readWindowsUSBSerialNames() map[string]string {
	result := make(map[string]string)
	comPortRegex := regexp.MustCompile(`\((COM\d+)\)\s*$`)
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

func actionSerialPortsUnix() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		ports := make(map[string]string)

		addPort := func(path, desc string) {
			if path == "" {
				return
			}
			if _, err := os.Stat(path); err != nil {
				return
			}
			if existing, ok := ports[path]; !ok || existing == "" {
				ports[path] = desc
			}
		}

		for path, desc := range readSerialPortsSysfs() {
			addPort(path, desc)
		}

		for _, pattern := range []string{
			"/dev/cu.*",
			"/dev/rfcomm*",
			"/dev/serial/by-id/*",
			"/dev/ttyACM*",
			"/dev/ttyAMA*",
			"/dev/ttyUSB*",
		} {
			matches, err := filepath.Glob(pattern)
			if err != nil {
				continue
			}
			for _, match := range matches {
				addPort(match, describeSerialPort(match))
			}
		}

		if runtime.GOOS == "darwin" {
			for path := range ports {
				if strings.HasPrefix(path, "/dev/tty.") {
					delete(ports, path)
				}
			}
		}

		paths := make([]string, 0, len(ports))
		for path := range ports {
			paths = append(paths, path)
		}
		sort.Strings(paths)

		result := make([]string, 0, len(paths)*2)
		for _, path := range paths {
			result = append(result, path, ports[path])
		}
		return carapace.ActionValuesDescribed(result...)
	})
}

func readSerialPortsSysfs() map[string]string {
	entries, err := os.ReadDir("/sys/class/tty")
	if err != nil {
		return map[string]string{}
	}

	ports := make(map[string]string)
	for _, entry := range entries {
		name := entry.Name()
		devicePath := filepath.Join("/sys/class/tty", name, "device")
		if _, err := os.Stat(devicePath); err != nil {
			if !isSerialTTYName(name) {
				continue
			}
		}

		devPath := "/dev/" + name
		if _, err := os.Stat(devPath); err != nil {
			continue
		}
		if !isLikelySerialTTY(name, devicePath) {
			continue
		}

		ports[devPath] = describeSysfsSerialPort(name, devicePath)
	}
	return ports
}

func isSerialTTYName(name string) bool {
	prefixes := []string{"cu.", "rfcomm", "ttyACM", "ttyAMA", "ttyS", "ttyUSB"}
	for _, prefix := range prefixes {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}
	return false
}

func isLikelySerialTTY(name, devicePath string) bool {
	if strings.HasPrefix(name, "ttyS") {
		_, err := os.Stat(filepath.Join(devicePath, "driver"))
		return err == nil
	}

	if !isSerialTTYName(name) {
		return false
	}

	if strings.HasPrefix(name, "cu.") || strings.HasPrefix(name, "ttyACM") || strings.HasPrefix(name, "ttyAMA") || strings.HasPrefix(name, "ttyUSB") || strings.HasPrefix(name, "rfcomm") {
		return true
	}

	_, err := os.Stat(devicePath)
	return err == nil
}

func describeSerialPort(path string) string {
	base := filepath.Base(path)
	if strings.HasPrefix(path, "/dev/serial/by-id/") {
		if target, err := filepath.EvalSymlinks(path); err == nil {
			return filepath.Base(target)
		}
		return "serial device"
	}

	devicePath := filepath.Join("/sys/class/tty", base, "device")
	if _, err := os.Stat(devicePath); err == nil {
		return describeSysfsSerialPort(base, devicePath)
	}
	return "serial port"
}

func describeSysfsSerialPort(name, devicePath string) string {
	subsystem := ""
	if target, err := filepath.EvalSymlinks(filepath.Join(devicePath, "subsystem")); err == nil {
		subsystem = filepath.Base(target)
	}

	if subsystem == "usb-serial" || subsystem == "usb" {
		if desc := describeUSBSerialPort(devicePath, subsystem); desc != "" {
			return desc
		}
	}

	if subsystem != "" {
		return subsystem
	}

	switch {
	case strings.HasPrefix(name, "cu."):
		return "callout device"
	case strings.HasPrefix(name, "ttyACM"):
		return "USB ACM"
	case strings.HasPrefix(name, "ttyAMA"):
		return "ARM UART"
	case strings.HasPrefix(name, "ttyS"):
		return "serial port"
	case strings.HasPrefix(name, "ttyUSB"):
		return "USB serial"
	case strings.HasPrefix(name, "rfcomm"):
		return "Bluetooth RFCOMM"
	default:
		return "serial port"
	}
}

func describeUSBSerialPort(devicePath, subsystem string) string {
	usbDevicePath := devicePath
	if subsystem == "usb-serial" {
		usbDevicePath = filepath.Dir(devicePath)
	}

	product := readSysfsString(filepath.Join(usbDevicePath, "product"))
	manufacturer := readSysfsString(filepath.Join(usbDevicePath, "manufacturer"))
	iface := readSysfsString(filepath.Join(devicePath, "interface"))

	switch {
	case product != "" && iface != "" && iface != product:
		return product + " - " + iface
	case product != "":
		return product
	case manufacturer != "":
		return manufacturer
	case iface != "":
		return iface
	default:
		return "USB serial"
	}
}

func readSysfsString(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(content))
}
