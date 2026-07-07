package qemu

import "github.com/carapace-sh/carapace"

// ActionImageFormats completes supported QEMU image formats
//
//	blkdebug (block debug)
//	blklogwrites (block log writes)
func ActionImageFormats() carapace.Action {
	return carapace.ActionValuesDescribed(
		"blkdebug", "block debug",
		"blklogwrites", "block log writes",
		"blkverify", "block verify",
		"bochs", "Bochs disk image",
		"cloop", "compressed loop",
		"compress", "compressed image",
		"copy-before-write", "copy before write filter",
		"copy-on-read", "copy on read filter",
		"dmg", "Apple Disk Management image",
		"file", "raw file",
		"ftp", "FTP",
		"ftps", "FTP over TLS",
		"gluster", "GlusterFS",
		"host_cdrom", "host CD-ROM",
		"host_device", "host device",
		"http", "HTTP",
		"https", "HTTPS",
		"iscsi", "iSCSI",
		"iser", "iSER",
		"luks", "LUKS encrypted",
		"nbd", "Network Block Device",
		"nfs", "NFS",
		"null-aio", "null AIO",
		"null-co", "null",
		"nvme", "NVMe",
		"parallels", "Parallels disk image",
		"preallocate", "preallocate filter",
		"qcow", "old QEMU Copy-On-Write",
		"qcow2", "QEMU Copy-On-Write v2",
		"qed", "QEMU Enhanced Disk",
		"quorum", "quorum filter",
		"raw", "raw disk image",
		"replication", "replication filter",
		"snapshot-access", "snapshot access filter",
		"ssh", "SSH",
		"throttle", "throttle filter",
		"vdi", "VirtualBox Disk Image",
		"vhdx", "VHDX disk image",
		"vmdk", "VMware disk image",
		"vpc", "VPC (VHD) disk image",
		"vvfat", "Virtual VFAT",
	).Tag("image formats").Uid("qemu", "image-format")
}

// ActionCacheModes completes QEMU cache modes
//
//	none (no caching)
//	writeback (write-back caching (default))
func ActionCacheModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"none", "no caching",
		"writeback", "write-back caching (default)",
		"writethrough", "write-through caching",
		"directsync", "direct sync",
		"unsafe", "unsafe caching",
	).Tag("cache modes").Uid("qemu", "cache-mode")
}

// ActionAioModes completes QEMU AIO modes
//
//	threads (pthreads based AIO)
//	native (Linux native AIO)
func ActionAioModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"threads", "pthreads based AIO",
		"native", "Linux native AIO",
		"io_uring", "Linux io_uring AIO",
	).Tag("aio modes").Uid("qemu", "aio-mode")
}

// ActionExportTypes completes QEMU storage daemon export types
//
//	nbd (Network Block Device)
//	fuse (Filesystem in Userspace)
func ActionExportTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"nbd", "Network Block Device",
		"fuse", "Filesystem in Userspace",
		"vhost-user-blk", "vhost-user block device",
		"vduse-blk", "VDUSE block device",
	).Tag("export types").Uid("qemu", "export-type")
}
