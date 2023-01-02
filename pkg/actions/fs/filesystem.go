package fs

import (
	"github.com/rsteube/carapace"
)

// ActionFilesystemTypes completes file system types
//
//	ext (is an elaborate extension of the minix filesystem.)
//	ext3 (is a journaling version of the ext2 filesystem.)
func ActionFilesystemTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"bfs", "is the native file system for the BeOS",
		"brtfs", "combines the copy-on-write principle with a logical volume manager",
		"cramfs", "a free read-only Linux file system designed for simplicity and space-efficiency",
		"exfat", "is a file system optimized for flash memory such as USB flash drives and SD cards",
		"ext", "is an elaborate extension of the minix filesystem.",
		"ext2", "is the high performance disk filesystem used by Linux  for  fixed  disks  as  well  as removable  media.",
		"ext3", "is a journaling version of the ext2 filesystem.",
		"ext4", "is  a  set  of  upgrades  to  ext3  including  substantial performance and reliability enhancements.",
		"fat", "is a file system developed for personal computers",
		"f2fs", "s a flash file system initially developed by Samsung Electronics",
		"hpfs", "is  the High Performance Filesystem, used in OS/2.",
		"iso9660", "is a CD-ROM filesystem type conforming to the ISO 9660 standard.",
		"jfs", "is a journaling filesystem, developed by IBM.",
		"minix", "is  the  filesystem  used in the Minix operating system, the first to run under Linux.",
		"msdos", "is  the filesystem used by DOS, Windows, and some OS/2 computers.",
		"ncpfs", "is  a  network  filesystem that supports the NCP protocol, used by Novell NetWare.",
		"nfs", "is the network filesystem used to access disks located on remote computers.",
		"ntfs", "is the filesystem native to Microsoft Windows NT.",
		"proc", "is a pseudo filesystem which is used as an interface to kernel data structures.",
		"reiserfs", "is a journaling filesystem, designed by Hans Reiser.",
		"smb", "is  a  network  filesystem  that  supports the SMB protocol.",
		"sysv", "is an implementation of the System V/Coherent filesystem for Linux.",
		"tmpfs", "is  a  filesystem  whose  contents  reside in virtual memory.",
		"umsdos", "is  an  extended DOS filesystem used by Linux.",
		"vfat", "is an extended FAT filesystem used by Microsoft Windows95 and Windows NT.",
		"xfs", "is a journaling filesystem, developed by SGI.",
		"xiafs", "was designed and implemented to be a stable, safe filesystem by  extending  the  Minix filesystem code.",
	).Tag("filesystem types")
}
