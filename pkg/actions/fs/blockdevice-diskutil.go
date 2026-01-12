package fs

import (
	"bytes"
	"fmt"
	"runtime"

	"github.com/carapace-sh/carapace"
	"github.com/micromdm/plist"
)

type blockdevicesDiskutil struct {
	AllDisksAndPartitions []struct {
		DeviceIdentifier string `plist:"DeviceIdentifier"`
		Content          string `plist:"Content"`
		Size             int    `plist:"Size"`
		APFSVolumes      []struct {
			DeviceIdentifier string `plist:"DeviceIdentifier"`
			DiskUUID         string `plist:"DiskUUID"`
			VolumeName       string `plist:"VolumeName"`
			VolumeUUID       string `plist:"VolumeUUID"`
			Size             int    `plist:"Size"`
		} `plist:"APFSVolumes"`
		Partitions []struct {
			DeviceIdentifier string `plist:"DeviceIdentifier"`
			Content          string `plist:"Content"`
			DiskUUID         string `plist:"DiskUUID"`
			VolumeName       string `plist:"VolumeName"`
			VolumeUUID       string `plist:"VolumeUUID"`
			Size             int    `plist:"Size"`
		} `plist:"Partitions"`
	} `plist:"AllDisksAndPartitions"`
}

func formatBlockSize(bytes int) string {
	const (
		kiB = 1024
		miB = kiB * 1024
		giB = miB * 1024
		tiB = giB * 1024
	)
	switch {
	case bytes >= tiB:
		return fmt.Sprintf("%.0fT", float64(bytes)/float64(tiB))
	case bytes >= giB:
		return fmt.Sprintf("%.0fG", float64(bytes)/float64(giB))
	case bytes >= miB:
		return fmt.Sprintf("%.0fM", float64(bytes)/float64(miB))
	case bytes >= kiB:
		return fmt.Sprintf("%.0fK", float64(bytes)/float64(kiB))
	default:
		return fmt.Sprintf("%dB", bytes)
	}
}

func (b blockdevicesDiskutil) toBlockdevices() []blockdevice {
	devices := make([]blockdevice, 0)
	for _, disk := range b.AllDisksAndPartitions {
		devices = append(devices, blockdevice{
			Kname:        disk.DeviceIdentifier,
			Path:         "/dev/" + disk.DeviceIdentifier,
			Size:         formatBlockSize(disk.Size),
			Type:         "disk",
			Parttypename: disk.Content,
		})
		for _, part := range disk.Partitions {
			devices = append(devices, blockdevice{
				Kname:        part.DeviceIdentifier,
				Label:        part.VolumeName,
				Partlabel:    part.VolumeName,
				Partuuid:     part.DiskUUID,
				Parttypename: part.Content,
				Path:         "/dev/" + part.DeviceIdentifier,
				Size:         formatBlockSize(part.Size),
				Type:         "part",
				Uuid:         part.VolumeUUID,
			})
		}
		for _, vol := range disk.APFSVolumes {
			devices = append(devices, blockdevice{
				Kname:        vol.DeviceIdentifier,
				Label:        vol.VolumeName,
				Partlabel:    vol.VolumeName,
				Partuuid:     vol.DiskUUID,
				Parttypename: "APFSVolume",
				Path:         "/dev/" + vol.DeviceIdentifier,
				Size:         formatBlockSize(vol.Size),
				Type:         "part",
				Uuid:         vol.VolumeUUID,
			})
		}
	}
	return devices
}

func actionBlockdevicesDiskutil(f func(blockdevices []blockdevice) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if runtime.GOOS != "darwin" {
			return carapace.ActionValues()
		}
		return carapace.ActionExecCommand("diskutil", "list", "-plist")(func(output []byte) carapace.Action {
			var b blockdevicesDiskutil
			if err := plist.NewXMLDecoder(bytes.NewReader(output)).Decode(&b); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return f(b.toBlockdevices())
		})
	})
}
