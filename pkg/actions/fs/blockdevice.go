package fs

import (
	"encoding/json"
	"fmt"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
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
	return carapace.ActionExecCommand("lsblk", "--json", "-o", "KNAME,LABEL,PARTLABEL,PARTUUID,PATH,SIZE,PARTTYPENAME,TYPE,UUID")(func(output []byte) carapace.Action {
		var b struct {
			Blockdevices []blockdevice
		}
		if err := json.Unmarshal(output, &b); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(b.Blockdevices)
	})
}

// ActionBlockDevices completes block devices
//   /dev/sda (10G)
//   /dev/sda1 (2G Linux swap)
func ActionBlockDevices() carapace.Action {
	return actionBlockdevices(func(blockdevices []blockdevice) carapace.Action {
		vals := make([]string, 0)
		for _, b := range blockdevices {
			vals = append(vals, b.Path, fmt.Sprintf("%v %v", b.Size, b.Parttypename), style.ForPath(b.Path))
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
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
	})
}

// ActionPartLabels completes partition labels
func ActionPartLabels() carapace.Action {
	return actionBlockdevices(func(blockdevices []blockdevice) carapace.Action {
		vals := make([]string, 0)
		for _, b := range blockdevices {
			if b.Label != "" {
				vals = append(vals, b.Partlabel, b.Kname)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
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
	})
}

// ActionPartUuids completes partition uuids
func ActionPartUuids() carapace.Action {
	return actionBlockdevices(func(blockdevices []blockdevice) carapace.Action {
		vals := make([]string, 0)
		for _, b := range blockdevices {
			if b.Partuuid != "" {
				vals = append(vals, b.Partuuid, b.Kname)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
