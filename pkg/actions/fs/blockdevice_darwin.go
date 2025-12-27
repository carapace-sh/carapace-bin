package fs

import (
	"github.com/carapace-sh/carapace"
	"github.com/micromdm/plist"
)

type darwinDeviceList struct {
	Internal []darwinDevice
	External []darwinDevice
}

func (l darwinDeviceList) ToBlockdevices() []blockdevice {
	blockdevices := make([]blockdevice, 0)

	for _, device := range append(l.Internal, l.External...) {
		blockdevices = append(blockdevices, blockdevice{
			Name:  device.DeviceIdentifier,
			Label: device.MediaName,
			// TODO Size         string
			Type: "disk", // TODO
			// Uuid         string
		})

		for _, partition := range device.Partitions {

			blockdevices = append(blockdevices, blockdevice{
				Name:          partition.DeviceIdentifier,
				Label:         partition.VolumeName,
				PartitionType: partition.PartitionType,
				Path:          partition.MountPoint,
				// TODO Size         string
				Type: "part", // TODO
				// Uuid         string
			})
		}
	}
	return blockdevices
}

type darwinDevice struct {
	DeviceIdentifier string            `plist:"DeviceIdentifier"`
	DeviceNode       string            `plist:"DeviceNode"`
	DeviceSize       int               `plist:"DeviceSize"`
	MediaName        string            `plist:"MediaName"`
	MediaType        string            `plist:"MediaType"`
	Partitions       []darwinPartition `plist:"Partitions"`
}

type darwinPartition struct {
	DeviceIdentifier string `plist:"DeviceIdentifier"`
	DeviceNode       string `plist:"DeviceNode"`
	DeviceSize       int    `plist:"DeviceSize"`
	MountPoint       string `plist:"MountPoint"`
	PartitionType    string `plist:"PartitionType"`
	VolumeName       string `plist:"VolumeName"`
}

func actionBlockdevices(f func(blockdevices []blockdevice) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("diskutil", "list", "-plist")(func(output []byte) carapace.Action {
		var deviceList darwinDeviceList
		if err := plist.Unmarshal(output, &deviceList); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(deviceList.ToBlockdevices())
	})
}
