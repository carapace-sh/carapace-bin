package fs

import "github.com/rsteube/carapace"

// ActionFileSystemTypes completes file system types
//   ext (is an elaborate extension of the minix filesystem.)
//   ext3 (is a journaling version of the ext2 filesystem.)
func ActionFileSystemTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"ext", "is an elaborate extension of the minix filesystem.",
		"ext2", "is the high performance disk filesystem used by Linux  for  fixed  disks  as  well  as removable  media.",
		"ext3", "is a journaling version of the ext2 filesystem.",
		"ext4", "is  a  set  of  upgrades  to  ext3  including  substantial performance and reliability enhancements.",
		"hpfs", "is  the High Performance Filesystem, used in OS/2.",
		"iso9660", "is a CD-ROM filesystem type conforming to the ISO 9660 standard.",
		"JFS", "is a journaling filesystem, developed by IBM.",
		"minix", "is  the  filesystem  used in the Minix operating system, the first to run under Linux.",
		"msdos", "is  the filesystem used by DOS, Windows, and some OS/2 computers.",
		"ncpfs", "is  a  network  filesystem that supports the NCP protocol, used by Novell NetWare.",
		"nfs", "is the network filesystem used to access disks located on remote computers.",
		"ntfs", "is the filesystem native to Microsoft Windows NT.",
		"proc", "is a pseudo filesystem which is used as an interface to kernel data structures.",
		"Reiserfs", "is a journaling filesystem, designed by Hans Reiser.",
		"smb", "is  a  network  filesystem  that  supports the SMB protocol.",
		"sysv", "is an implementation of the System V/Coherent filesystem for Linux.",
		"umsdos", "is  an  extended DOS filesystem used by Linux.",
		"tmpfs", "is  a  filesystem  whose  contents  reside in virtual memory.",
		"vfat", "is an extended FAT filesystem used by Microsoft Windows95 and Windows NT.",
		"XFS", "is a journaling filesystem, developed by SGI.",
		"xiafs", "was designed and implemented to be a stable, safe filesystem by  extending  the  Minix filesystem code.",
	)
}
