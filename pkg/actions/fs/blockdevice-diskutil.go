package fs

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/micromdm/plist"
)

type blockdevicesDiskutil struct {
	AllDisksAndPartitions []struct {
		DeviceIdentifier string `plist:"DeviceIdentifier"`
		APFSVolumes      []struct {
			DeviceIdentifier string `plist:"DeviceIdentifier"`
			DiskUUID         string `plist:"DiskUUID"`
			VolumeName       string `plist:"VolumeName"`
			VolumeUUID       string `plist:"VolumeUUID"`
			Size             int    `plist:"Size"`
			MountPoint       string `plist:"MountPoint"`
		} `plist:"APFSVolumes"`
		Partitions []struct {
			DeviceIdentifier string `plist:"DeviceIdentifier"`
			DiskUUID         string `plist:"DiskUUID"`
			VolumeName       string `plist:"VolumeName"`
			VolumeUUID       string `plist:"VolumeUUID"`
			Size             int    `plist:"Size"`
			MountPoint       string `plist:"MountPoint"`
		} `plist:"Partitions"`
	} `plist:"AllDisksAndPartitions"`
}

func actionBlockdevicesDiskutil(f func(blockdevices []blockdevice) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("diskutil", "list", "-plist")(func(output []byte) carapace.Action {
		var b blockdevicesDiskutil

		err := plist.NewXMLDecoder(bytes.NewReader(output)).Decode(&b)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		// return f(b) // TODO convert
		return carapace.ActionValues()
	})
}
