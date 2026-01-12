package mount

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
)

// ActionDarwinMountOptions completes macOS mount options
//
//	async
//	noatime
func ActionDarwinMountOptions() carapace.Action {
	return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValuesDescribed(
				"async", "All I/O to the filesystem should be done asynchronously.",
				"noasync", "This filesystem should not force all I/O to be written asynchronously.",
				"auto", "Can be mounted with the -a option.",
				"noauto", "This filesystem should be skipped when mount is run with -a.",
				"force", "Force revocation of write access when downgrading to read-only.",
				"nodev", "Do not interpret character or block special devices on the filesystem.",
				"noexec", "Do not allow execution of any binaries on the mounted filesystem.",
				"noowners", "Ignore the ownership field for the entire volume.",
				"nosuid", "Do not allow set-user-identifier or set-group-identifier bits to take effect.",
				"rdonly", "Mount the filesystem read-only.",
				"sync", "All I/O to the filesystem should be done synchronously.",
				"update", "Change the status of an already mounted filesystem.",
				"union", "Causes the namespace to appear as the union of directories.",
				"noatime", "Do not update the file access time when reading from a file.",
				"strictatime", "Always update the file access time when reading from a file.",
				"nobrowse", "This option indicates that the mount point should not be visible via the GUI.",
				"nofollow", "Don't follow any symlinks in the provided mount-on directory.",
			)
		default:
			return carapace.ActionValues()
		}
	}).Tag("mount options")
}

// ActionDarwinSources completes macOS mount sources
//
//	/dev/disk0s2
//	LABEL=Data
func ActionDarwinSources() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.Batch(
			fs.ActionBlockDevices(),
			carapace.ActionFiles(),
			carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"LABEL", "specifies device by volume name",
						"UUID", "specifies device by volume UUID",
					).Suffix("=")
				case 1:
					switch c.Parts[0] {
					case "LABEL":
						return fs.ActionLabels()
					case "UUID":
						return fs.ActionUuids()
					default:
						return carapace.ActionValues()
					}
				default:
					return carapace.ActionValues()
				}
			}),
		).ToA()
	}).Tag("sources")
}
