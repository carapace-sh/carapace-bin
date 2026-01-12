package fs

import (
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/carapace-sh/carapace/pkg/style"
)

type blockdevice struct {
	Kname        string
	Label        string
	Partlabel    string
	Partuuid     string
	Parttypename string
	Path         string
	Size         string
	Type         string
	Uuid         string
}

func actionBlockdevices(f func(blockdevices []blockdevice) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch {
		case condition.Executable("lsblk")(c): // unix
			return actionBlockdevicesLsblk(f)
		case condition.Executable("diskutil")(c): // darwin
			return actionBlockdevicesDiskutil(f)
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionBlockDevices completes block devices
//
//	/dev/sda (10G)
//	/dev/sda1 (2G Linux swap)
func ActionBlockDevices() carapace.Action {
	return actionBlockdevices(func(blockdevices []blockdevice) carapace.Action {
		vals := make([]string, 0)
		for _, b := range blockdevices {
			vals = append(vals, b.Path, fmt.Sprintf("%v %v", b.Size, b.Parttypename))
		}
		return carapace.ActionValuesDescribed(vals...).StyleF(style.ForPath)
	}).Tag("block devices")
}

// ActionLabels completes disk labels
//
//	ROOT (sda)
//	BOOT (sda1)
func ActionLabels() carapace.Action {
	return actionBlockdevices(func(blockdevices []blockdevice) carapace.Action {
		vals := make([]string, 0)
		for _, b := range blockdevices {
			if b.Label != "" {
				vals = append(vals, b.Label, b.Kname)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("labels")
}

// ActionPartitionLabels completes partition labels
//
//	EFI System Partition (sda1)
//	Linux filesystem (sda2)
func ActionPartitionLabels() carapace.Action {
	return actionBlockdevices(func(blockdevices []blockdevice) carapace.Action {
		vals := make([]string, 0)
		for _, b := range blockdevices {
			if b.Label != "" {
				vals = append(vals, b.Partlabel, b.Kname)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("partition labels")
}

// ActionUuids completes disk uuids
//
//	a1b2c3d4-e5f6-7890-abcd-ef1234567890 (sda)
//	11223344-5566-7788-99aa-bbccddeeff00 (sda1)
func ActionUuids() carapace.Action {
	return actionBlockdevices(func(blockdevices []blockdevice) carapace.Action {
		vals := make([]string, 0)
		for _, b := range blockdevices {
			if b.Uuid != "" {
				vals = append(vals, b.Uuid, b.Kname)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("uuids")
}

// ActionPartitionUuids completes partition uuids
//
//	a1b2c3d4-e5f6-7890-abcd-ef1234567890 (sda1)
//	11223344-5566-7788-99aa-bbccddeeff00 (sda2)
func ActionPartitionUuids() carapace.Action {
	return actionBlockdevices(func(blockdevices []blockdevice) carapace.Action {
		vals := make([]string, 0)
		for _, b := range blockdevices {
			if b.Partuuid != "" {
				vals = append(vals, b.Partuuid, b.Kname)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("partition uuids")
}
