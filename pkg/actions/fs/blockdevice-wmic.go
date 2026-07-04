package fs

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func actionBlockdevicesWmic(f func(blockdevices []blockdevice) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("wmic", "logicaldisk", "get", "DeviceID,Description,FileSystem,Size,DriveType", "/format:csv")(func(output []byte) carapace.Action {
		lines := strings.Split(strings.ReplaceAll(string(output), "\r", ""), "\n")
		if len(lines) < 2 {
			return carapace.ActionValues()
		}

		// CSV header: Node,DeviceID,Description,DriveType,FileSystem,Size
		header := strings.Split(lines[0], ",")
		indexes := make(map[string]int)
		for i, col := range header {
			indexes[strings.TrimSpace(col)] = i
		}

		devices := make([]blockdevice, 0)
		for _, line := range lines[1:] {
			fields := strings.Split(line, ",")
			if len(fields) < len(header) {
				continue
			}

			deviceID := strings.TrimSpace(fields[indexes["DeviceID"]])
			if deviceID == "" {
				continue
			}

			size := strings.TrimSpace(fields[indexes["Size"]])
			size = formatBlockSize(parseWmicSize(size))

			devices = append(devices, blockdevice{
				Kname:        deviceID,
				Label:        strings.TrimSpace(fields[indexes["Description"]]),
				Parttypename: strings.TrimSpace(fields[indexes["FileSystem"]]),
				Path:         deviceID,
				Size:         size,
				Type:         strings.TrimSpace(fields[indexes["DriveType"]]),
			})
		}
		return f(devices)
	})
}

func parseWmicSize(s string) int64 {
	if s == "" {
		return 0
	}
	var n int64
	for _, c := range s {
		if c >= '0' && c <= '9' {
			n = n*10 + int64(c-'0')
		}
	}
	return n
}
