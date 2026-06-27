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
				return carapace.ActionValuesDescribed(
					"async", "All I/O to the filesystem should be done asynchronously.",
					"atime", "Do not use the noatime feature.",
					"noatime", "Do not update inode access times on this filesystem.",
					"auto", "Can be mounted with the -a option.",
					"noauto", "Can only be mounted explicitly.",
					"context=", "Set context",
					"fscontext=", "Set context",
					"defcontext=", "Set context",
					"rootcontext=", "Set context",
					"defaults", "Use the default options",
					"dev", "Interpret character or block special devices on the filesystem.",
					"nodev", "Do not interpret character or block special devices on the filesystem.",
					"diratime", "Update directory inode access times on this filesystem.",
					"nodiratime", "Do not update directory inode access times on this filesystem.",
					"dirsync", "All directory updates within the filesystem should be done synchronously.",
					"exec", "Permit execution of binaries.",
					"noexec", "Do not permit direct execution of any binaries on the mounted filesystem.",
					"group", "Allow an ordinary user to mount the filesystem if one of that user's groups matches the group of the device.",
					"iversion", "Every time the inode is modified, the i_version field will be incremented.",
					"noiversion", "Do not increment the i_version inode field.",
					"mand", "Allow mandatory locks on this filesystem.",
					"nomand", "Do not allow mandatory locks on this filesystem.",
					"_netdev", "The filesystem resides on a device that requires network access",
					"nofail", "Do not report errors for this device if it does not exist.",
					"relatime", "Update inode access times relative to modify or change time.",
					"norelatime", "Do not use the relatime feature.",
					"strictatime", "Allows to explicitly request full atime updates.",
					"nostrictatime", "Use the kernel's default behavior for inode access time updates.",
					"lazytime", "Only update times on the in-memory version of the file inode.",
					"nolazytime", "Do not use the lazytime feature.",
					"suid", "Honor set-user-ID and set-group-ID bits or file capabilities",
					"nosuid", "Do not honor set-user-ID and set-group-ID bits or file capabilities",
					"silent", "Turn on the silent flag.",
					"loud", "Turn off the silent flag.",
					"owner", "Allow an ordinary user to mount the filesystem if that user is the owner of the device.",
					"remount", "Attempt to remount an already-mounted filesystem.",
					"ro", "Mount the filesystem read-only.",
					"rw", "Mount the filesystem read-write.",
					"sync", "All I/O to the filesystem should be done synchronously.",
					"user", "Allow an ordinary user to mount the filesystem.",
					"nouser", "Forbid an ordinary user to mount the filesystem.",
					"users", "Allow any user to mount and to unmount the filesystem",
					"X-mount.mkdir", "Allow to make a target directory if it does not exist yet.",
					"X-mount.subdir=", "Allow mounting sub-directory from a filesystem instead of the root directory.",
					"nosymfollow", "Do not follow symlinks when resolving paths.",
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
