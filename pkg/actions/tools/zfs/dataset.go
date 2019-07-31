// package zfs contains OpenZFS related actions
package zfs

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace/pkg/style"
)

// DatasetOpts configures which dataset types to complete
type DatasetOpts struct {
	Filesystem bool
	Volume     bool
	Snapshot   bool
	Bookmark   bool
}

func (o DatasetOpts) types() string {
	types := make([]string, 0, 4)
	if o.Filesystem {
		types = append(types, "filesystem")
	}
	if o.Volume {
		types = append(types, "volume")
	}
	if o.Snapshot {
		types = append(types, "snapshot")
	}
	if o.Bookmark {
		types = append(types, "bookmark")
	}
	if len(types) == 0 {
		return "all"
	}
	return strings.Join(types, ",")
}

// ActionDatasets completes ZFS dataset names filtered by type
//
//	rpool (filesystem)
//	rpool/data@snap1 (snapshot)
func ActionDatasets(opts DatasetOpts) carapace.Action {
	return carapace.ActionExecCommand("zfs", "list", "-H", "-o", "name,type", "-t", opts.types())(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines {
			if fields := strings.SplitN(line, "\t", 2); len(fields) == 2 {
				name := fields[0]
				dsType := strings.TrimSpace(fields[1])
				var s string
				switch dsType {
				case "filesystem":
					s = styles.Zfs.Filesystem
				case "volume":
					s = styles.Zfs.Volume
				case "snapshot":
					s = styles.Zfs.Snapshot
				case "bookmark":
					s = styles.Zfs.Bookmark
				default:
					s = style.Default
				}
				vals = append(vals, name, dsType, s)
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Tag("datasets")
}

// ActionFilesystems completes ZFS filesystem and volume names
//
//	rpool
//	rpool/data
func ActionFilesystems() carapace.Action {
	return ActionDatasets(DatasetOpts{Filesystem: true, Volume: true})
}

// ActionSnapshots completes ZFS snapshot names
//
//	rpool/data@snap1
//	rpool/data@snap2
func ActionSnapshots() carapace.Action {
	return ActionDatasets(DatasetOpts{Snapshot: true})
}

// ActionBookmarks completes ZFS bookmark names
//
//	rpool/data#mark1
//	rpool/data#mark2
func ActionBookmarks() carapace.Action {
	return ActionDatasets(DatasetOpts{Bookmark: true})
}

// ActionMountedFilesystems completes mounted ZFS filesystems
//
//	rpool
//	rpool/data
func ActionMountedFilesystems() carapace.Action {
	return carapace.ActionExecCommand("zfs", "list", "-H", "-o", "name,mountpoint", "-t", "filesystem")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines {
			if fields := strings.SplitN(line, "\t", 2); len(fields) == 2 {
				name := strings.TrimSpace(fields[0])
				mountpoint := strings.TrimSpace(fields[1])
				if name != "" && mountpoint != "none" && mountpoint != "-" {
					vals = append(vals, name, mountpoint)
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("filesystems")
}
