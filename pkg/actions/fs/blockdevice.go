package fs

import (
	"fmt"

	"github.com/carapace-sh/carapace"
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

// TODO add examples to actions

// ActionLabels completes disk labels
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
