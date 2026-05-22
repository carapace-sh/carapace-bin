package fs

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"

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

var darwinDiskIdentifierRegex = regexp.MustCompile(`^disk[0-9]+(s[0-9]+)*$`)

func actionBlockdevices(f func(blockdevices []blockdevice) carapace.Action) carapace.Action {
	switch runtime.GOOS {
	case "linux":
		return carapace.ActionExecCommand("lsblk", "--json", "-o", "KNAME,LABEL,PARTLABEL,PARTUUID,PATH,SIZE,PARTTYPENAME,TYPE,UUID")(func(output []byte) carapace.Action {
			var b struct {
				Blockdevices []blockdevice
			}
			if err := json.Unmarshal(output, &b); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return f(b.Blockdevices)
		})
	case "darwin":
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return f(darwinBlockdevices())
		})
	default:
		return f(nil)
	}
}

func darwinBlockdevices() []blockdevice {
	entries, err := os.ReadDir("/dev")
	if err != nil {
		return nil
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		names = append(names, entry.Name())
	}
	return darwinBlockdevicesFromNames(names)
}

func darwinBlockdevicesFromNames(names []string) []blockdevice {
	blockdevices := make([]blockdevice, 0)
	for _, name := range names {
		if !darwinDiskIdentifierRegex.MatchString(name) {
			continue
		}
		blockdevices = append(blockdevices, blockdevice{
			Kname: name,
			Path:  "/dev/" + name,
		})
	}
	return blockdevices
}

// ActionBlockDevices completes block devices
//
//	/dev/sda (10G)
//	/dev/sda1 (2G Linux swap)
func ActionBlockDevices() carapace.Action {
	return actionBlockdevices(func(blockdevices []blockdevice) carapace.Action {
		vals := make([]string, 0)
		for _, b := range blockdevices {
			vals = append(vals, b.Path, strings.TrimSpace(fmt.Sprintf("%v %v", b.Size, b.Parttypename)))
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
