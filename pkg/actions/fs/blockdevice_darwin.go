package fs

import (
	"github.com/carapace-sh/carapace"
	"github.com/micromdm/plist"
)

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

func (d darwinPartition) ToBlockdevice() blockdevice {
	return blockdevice{
		Name:          d.DeviceIdentifier,
		Label:         d.VolumeName,
		PartitionType: d.PartitionType,
		Path:          d.MountPoint,
		// TODO Size         string
		Type: "part", // TODO
		// Uuid         string
	}
}

func actionBlockdevices(f func(blockdevices []blockdevice) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("diskutil", "list", "-plist")(func(output []byte) carapace.Action {
		var deviceList struct {
			Internal []darwinDevice
			External []darwinDevice
		}

		if err := plist.Unmarshal(output, &deviceList); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		return carapace.ActionValues("TODO")
	})
}
