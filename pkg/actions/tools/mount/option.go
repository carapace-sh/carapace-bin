package mount

import (
	"runtime"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
)

// ActionMountOptions completes mount options
//
//	noatime
//	nodev
func ActionMountOptions() carapace.Action {
	return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			switch runtime.GOOS {
			case "darwin":
				return carapace.ActionValuesDescribed(
					"async", "All I/O to the filesystem should be done asynchronously.",
					"auto", "Can be mounted with the -a option.",
					"force", "Force revocation of write access when downgrading to read-only.",
					"noasync", "This filesystem should not force all I/O to be written asynchronously.",
					"noatime", "Do not update the file access time when reading from a file.",
					"noauto", "This filesystem should be skipped when mount is run with -a.",
					"nobrowse", "This option indicates that the mount point should not be visible via the GUI.",
					"nodev", "Do not interpret character or block special devices on the filesystem.",
					"noexec", "Do not allow execution of any binaries on the mounted filesystem.",
					"nofollow", "Don't follow any symlinks in the provided mount-on directory.",
					"noowners", "Ignore the ownership field for the entire volume.",
					"nosuid", "Do not allow set-user-identifier or set-group-identifier bits to take effect.",
					"rdonly", "Mount the filesystem read-only.",
					"strictatime", "Always update the file access time when reading from a file.",
					"sync", "All I/O to the filesystem should be done synchronously.",
					"union", "Causes the namespace to appear as the union of directories.",
					"update", "Change the status of an already mounted filesystem.",
				)
			case "windows":
				return carapace.ActionValuesDescribed(
					"async", "All I/O to the filesystem should be done asynchronously.",
					"noatime", "Do not update the file access time when reading from a file.",
					"nodev", "Do not interpret character or block special devices on the filesystem.",
					"noexec", "Do not allow execution of any binaries on the mounted filesystem.",
					"nosuid", "Do not allow set-user-identifier or set-group-identifier bits to take effect.",
					"ro", "Mount the volume read-only.",
					"rw", "Mount the volume read-write.",
					"sync", "All I/O to the filesystem should be done synchronously.",
				)
			default:
				return carapace.ActionValuesDescribed(
					"X-mount.mkdir", "Allow to make a target directory if it does not exist yet.",
					"X-mount.subdir=", "Allow mounting sub-directory from a filesystem instead of the root directory.",
					"_netdev", "The filesystem resides on a device that requires network access",
					"async", "All I/O to the filesystem should be done asynchronously.",
					"atime", "Do not use the noatime feature.",
					"auto", "Can be mounted with the -a option.",
					"context=", "Set context",
					"defaults", "Use the default options",
					"defcontext=", "Set context",
					"dev", "Interpret character or block special devices on the filesystem.",
					"diratime", "Update directory inode access times on this filesystem.",
					"dirsync", "All directory updates within the filesystem should be done synchronously.",
					"exec", "Permit execution of binaries.",
					"fscontext=", "Set context",
					"group", "Allow an ordinary user to mount the filesystem if one of that user's groups matches the group of the device.",
					"iversion", "Every time the inode is modified, the i_version field will be incremented.",
					"lazytime", "Only update times on the in-memory version of the file inode.",
					"loud", "Turn off the silent flag.",
					"mand", "Allow mandatory locks on this filesystem.",
					"noatime", "Do not update inode access times on this filesystem.",
					"noauto", "Can only be mounted explicitly.",
					"nodev", "Do not interpret character or block special devices on the filesystem.",
					"nodiratime", "Do not update directory inode access times on this filesystem.",
					"noexec", "Do not permit direct execution of any binaries on the mounted filesystem.",
					"nofail", "Do not report errors for this device if it does not exist.",
					"noiversion", "Do not increment the i_version inode field.",
					"nolazytime", "Do not use the lazytime feature.",
					"nomand", "Do not allow mandatory locks on this filesystem.",
					"norelatime", "Do not use the relatime feature.",
					"nostrictatime", "Use the kernel's default behavior for inode access time updates.",
					"nosuid", "Do not honor set-user-ID and set-group-ID bits or file capabilities",
					"nosymfollow", "Do not follow symlinks when resolving paths.",
					"nouser", "Forbid an ordinary user to mount the filesystem.",
					"owner", "Allow an ordinary user to mount the filesystem if that user is the owner of the device.",
					"relatime", "Update inode access times relative to modify or change time.",
					"remount", "Attempt to remount an already-mounted filesystem.",
					"ro", "Mount the filesystem read-only.",
					"rootcontext=", "Set context",
					"rw", "Mount the filesystem read-write.",
					"silent", "Turn on the silent flag.",
					"strictatime", "Allows to explicitly request full atime updates.",
					"suid", "Honor set-user-ID and set-group-ID bits or file capabilities",
					"sync", "All I/O to the filesystem should be done synchronously.",
					"user", "Allow an ordinary user to mount the filesystem.",
					"users", "Allow any user to mount and to unmount the filesystem",
				)
			}
		case 1:
			switch c.Parts[0] {
			case "X-mount.mkdir":
				return fs.ActionFileModesNumeric()
			default:
				return carapace.ActionValues()
			}
		default:
			return carapace.ActionValues()
		}
	}).Tag("mount options")
}
