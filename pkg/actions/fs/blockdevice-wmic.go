package fs

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

func actionBlockdevicesWmic(f func(blockdevices []blockdevice) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("powershell", "-NoProfile", "-Command", "Get-CimInstance -ClassName Win32_LogicalDisk -Property DeviceID,Description,FileSystem,Size,DriveType | Select-Object DeviceID,Description,FileSystem,Size,DriveType | ConvertTo-Json -Compress")(func(output []byte) carapace.Action {
		var disks []struct {
			DeviceID    string
			Description string
			FileSystem  string
			Size        int64
			DriveType   int
		}
		if err := json.Unmarshal(output, &disks); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		devices := make([]blockdevice, 0)
		for _, d := range disks {
			if d.DeviceID == "" {
				continue
			}
			devices = append(devices, blockdevice{
				Kname:        d.DeviceID,
				Label:        d.Description,
				Parttypename: d.FileSystem,
				Path:         d.DeviceID,
				Size:         formatBlockSize(d.Size),
				Type:         formatDriveType(d.DriveType),
			})
		}
		return f(devices)
	})
}

func formatDriveType(driveType int) string {
	switch driveType {
	case 1:
		return "disk"
	case 2:
		return "removable"
	case 3:
		return "fixed"
	case 4:
		return "network"
	case 5:
		return "cd"
	case 6:
		return "ram"
	default:
		return "unknown"
	}
}
