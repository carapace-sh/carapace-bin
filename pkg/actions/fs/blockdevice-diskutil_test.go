package fs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	_ "embed"

	"github.com/micromdm/plist"
)

//go:embed blockdevice-diskutil.plist
var blockdeviceDiskutil []byte

func TestBlockDevicesPlist(t *testing.T) {
	var parsed struct {
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
	err := plist.NewXMLDecoder(bytes.NewReader(blockdeviceDiskutil)).Decode(&parsed)
	if err != nil {
		t.Error(err)
	}

	m, _ := json.MarshalIndent(parsed, "", "  ")
	fmt.Println(string(m))
}
