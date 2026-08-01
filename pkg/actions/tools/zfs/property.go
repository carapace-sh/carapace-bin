package zfs

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionProperties completes ZFS dataset property names
//
//	compression (compression algorithm)
//	mountpoint (mount point)
func ActionProperties() carapace.Action {
	return carapace.ActionValuesDescribed(
		"aclinherit", "ACL inherit mode",
		"aclmode", "ACL mode",
		"acltype", "ACL type",
		"atime", "update access time on read",
		"canmount", "allow mounting",
		"casesensitivity", "filename case sensitivity",
		"checksum", "checksum algorithm",
		"compression", "compression algorithm",
		"context", "SELinux context",
		"copies", "number of data copies",
		"dedup", "deduplication",
		"defaultgroupobjquota", "default group object quota",
		"defaultgroupquota", "default group quota",
		"defaultprojectobjquota", "default project object quota",
		"defaultprojectquota", "default project quota",
		"defaultuserobjquota", "default user object quota",
		"defaultuserquota", "default user quota",
		"defcontext", "SELinux default context",
		"devices", "allow device nodes",
		"direct", "direct I/O behavior",
		"dnodesize", "dnode size",
		"encryption", "encryption algorithm",
		"exec", "allow execution of binaries",
		"filesystem_limit", "max number of filesystems",
		"fscontext", "SELinux filesystem context",
		"jailed", "managed from FreeBSD jail",
		"keyformat", "encryption key format",
		"keylocation", "encryption key location",
		"logbias", "log bias",
		"mountpoint", "mount point",
		"nbmand", "non-blocking mandatory locks",
		"normalization", "unicode normalization",
		"overlay", "allow overlay mount",
		"pbkdf2iters", "PBKDF2 iterations",
		"prefetch", "prefetch behavior",
		"primarycache", "primary cache",
		"quota", "dataset quota",
		"readonly", "read-only",
		"recordsize", "suggested block size",
		"redundant_metadata", "redundant metadata",
		"refquota", "dataset referenced quota",
		"refreservation", "dataset referenced reservation",
		"relatime", "relative access time updates",
		"reservation", "dataset reservation",
		"rootcontext", "SELinux root context",
		"secondarycache", "secondary cache",
		"setuid", "allow setuid",
		"sharenfs", "NFS share options",
		"sharesmb", "SMB share options",
		"snapdev", "snapshot device visibility",
		"snapdir", "snapshot directory visibility",
		"snapshot_limit", "max number of snapshots",
		"special_small_blocks", "special vdev small block threshold",
		"sync", "sync behavior",
		"utf8only", "reject non-UTF-8 filenames",
		"version", "on-disk version",
		"volblocksize", "volume block size",
		"volmode", "volume mode",
		"volsize", "volume size",
		"volthreading", "volume threading",
		"vscan", "virus scan",
		"xattr", "extended attributes",
		"zoned", "managed from non-global zone",
	).Tag("properties")
}

// ActionReadonlyProperties completes read-only ZFS dataset properties
//
//	available (available space)
//	used (used space)
func ActionReadonlyProperties() carapace.Action {
	return carapace.ActionValuesDescribed(
		"available", "available space",
		"clones", "snapshot clones",
		"compressratio", "compression ratio",
		"createtxg", "creation transaction group",
		"creation", "creation time",
		"defer_destroy", "marked for deferred destroy",
		"encryptionroot", "encryption root",
		"filesystem_count", "number of filesystems",
		"guid", "dataset GUID",
		"keystatus", "encryption key status",
		"logicalreferenced", "logical referenced size",
		"logicalused", "logical used size",
		"mounted", "is mounted",
		"objsetid", "object set ID",
		"origin", "clone origin",
		"receive_resume_token", "receive resume token",
		"redact_snaps", "redaction snapshots",
		"refcompressratio", "referenced compression ratio",
		"referenced", "referenced space",
		"snapshot_count", "number of snapshots",
		"snapshots_changed", "time of last snapshot change",
		"type", "dataset type",
		"used", "used space",
		"usedbychildren", "used by children",
		"usedbydataset", "used by dataset",
		"usedbyrefreservation", "used by refreservation",
		"usedbysnapshots", "used by snapshots",
		"userrefs", "number of user holds",
		"written", "written since last snapshot",
	).Tag("properties")
}

// ActionPropertyValues completes values for a given ZFS property
//
//	on
//	off
func ActionPropertyValues(property string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch property {
		case "aclinherit":
			return carapace.ActionValues("discard", "noallow", "restricted", "passthrough", "passthrough-x")
		case "aclmode":
			return carapace.ActionValues("discard", "groupmask", "passthrough", "restricted")
		case "acltype":
			return carapace.ActionValues("off", "nfsv4", "posix", "noacl", "posixacl")
		case "atime", "relatime", "devices", "exec", "jailed", "nbmand", "overlay",
			"readonly", "setuid", "utf8only", "volthreading", "vscan", "zoned":
			return carapace.ActionValues("on", "off").StyleF(style.ForKeyword)
		case "canmount":
			return carapace.ActionValues("on", "off", "noauto").StyleF(style.ForKeyword)
		case "casesensitivity":
			return carapace.ActionValues("sensitive", "insensitive", "mixed")
		case "checksum":
			return carapace.ActionValues("on", "off", "fletcher2", "fletcher4", "sha256", "noparity", "sha512", "skein", "edonr", "blake3")
		case "compression":
			return carapace.ActionValues("on", "off", "gzip", "gzip-1", "gzip-2", "gzip-3", "gzip-4", "gzip-5",
				"gzip-6", "gzip-7", "gzip-8", "gzip-9", "lz4", "lzjb", "zle",
				"zstd", "zstd-1", "zstd-2", "zstd-3", "zstd-4", "zstd-5", "zstd-6", "zstd-7",
				"zstd-8", "zstd-9", "zstd-10", "zstd-11", "zstd-12", "zstd-13", "zstd-14",
				"zstd-15", "zstd-16", "zstd-17", "zstd-18", "zstd-19", "zstd-fast",
				"zstd-fast-1", "zstd-fast-2", "zstd-fast-3", "zstd-fast-4", "zstd-fast-5",
				"zstd-fast-6", "zstd-fast-7", "zstd-fast-8", "zstd-fast-9", "zstd-fast-10",
				"zstd-fast-20", "zstd-fast-30", "zstd-fast-40", "zstd-fast-50",
				"zstd-fast-60", "zstd-fast-70", "zstd-fast-80", "zstd-fast-90",
				"zstd-fast-100", "zstd-fast-500", "zstd-fast-1000")
		case "copies":
			return carapace.ActionValues("1", "2", "3")
		case "dedup":
			return carapace.ActionValues("on", "off", "verify", "sha256", "sha256,verify", "sha512", "sha512,verify",
				"skein", "skein,verify", "edonr,verify", "blake3", "blake3,verify")
		case "direct":
			return carapace.ActionValues("disabled", "standard", "always")
		case "dnodesize":
			return carapace.ActionValues("legacy", "auto", "1k", "2k", "4k", "8k", "16k")
		case "encryption":
			return carapace.ActionValues("off", "on", "aes-128-ccm", "aes-192-ccm", "aes-256-ccm",
				"aes-128-gcm", "aes-192-gcm", "aes-256-gcm")
		case "keyformat":
			return carapace.ActionValues("raw", "hex", "passphrase")
		case "keylocation":
			return carapace.Batch(
				carapace.ActionValues("prompt"),
				carapace.ActionFiles(),
			).ToA()
		case "logbias":
			return carapace.ActionValues("latency", "throughput")
		case "normalization":
			return carapace.ActionValues("none", "formC", "formD", "formKC", "formKD")
		case "prefetch":
			return carapace.ActionValues("all", "none", "metadata")
		case "mountpoint":
			return carapace.Batch(
				carapace.ActionValues("none", "legacy"),
				carapace.ActionDirectories(),
			).ToA()
		case "primarycache", "secondarycache":
			return carapace.ActionValues("all", "none", "metadata")
		case "defaultgroupobjquota", "defaultgroupquota", "defaultprojectobjquota", "defaultprojectquota",
			"defaultuserobjquota", "defaultuserquota", "quota", "refquota", "reservation", "volsize":
			return carapace.ActionValues("none")
		case "refreservation":
			return carapace.ActionValues("none", "auto")
		case "recordsize":
			return carapace.ActionValues("512", "1K", "2K", "4K", "8K", "16K", "32K", "64K", "128K", "256K", "512K", "1M")
		case "redundant_metadata":
			return carapace.ActionValues("all", "most", "some", "none")
		case "sharenfs":
			return carapace.ActionValues("on", "off").StyleF(style.ForKeyword)
		case "sharesmb":
			return carapace.ActionValues("on", "off").StyleF(style.ForKeyword)
		case "snapdev":
			return carapace.ActionValues("hidden", "visible")
		case "snapdir":
			return carapace.ActionValues("disabled", "hidden", "visible")
		case "special_small_blocks":
			return carapace.ActionValues("0", "512", "1K", "2K", "4K", "8K", "16K", "32K", "64K", "128K", "256K", "512K", "1M")
		case "sync":
			return carapace.ActionValues("standard", "always", "disabled")
		case "volblocksize":
			return carapace.ActionValues("512", "1K", "2K", "4K", "8K", "16K", "32K", "64K", "128K")
		case "volmode":
			return carapace.ActionValues("default", "full", "geom", "dev", "none")
		case "xattr":
			return carapace.ActionValues("on", "off", "dir", "sa")
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionPropertyAssignments completes property=value assignments
//
//	compression=lz4
//	mountpoint=/data
func ActionPropertyAssignments() carapace.Action {
	return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionProperties().Suffix("=")
		default:
			return ActionPropertyValues(c.Parts[0])
		}
	})
}

// ActionPoolFeatures completes ZFS pool feature flag names
//
//	feature@async_destroy (enabled)
//	feature@encryption (enabled)
func ActionPoolFeatures() carapace.Action {
	return carapace.ActionExecCommand("zpool", "get", "-H", "all")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		seen := make(map[string]bool)
		vals := make([]string, 0)

		for _, line := range lines {
			if fields := strings.Split(line, "\t"); len(fields) >= 3 {
				name := fields[1]
				if !strings.HasPrefix(name, "feature@") || seen[name] {
					continue
				}
				seen[name] = true
				vals = append(vals, name, fields[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...).Tag("pool features")
	})
}

// ActionPoolProperties completes ZFS pool property names
//
//	autoexpand (automatically expand pool)
//	autoreplace (automatically replace devices)
func ActionPoolProperties() carapace.Action {
	return carapace.Batch(
		carapace.ActionValuesDescribed(
			"altroot", "alternate root directory",
			"ashift", "pool sector size exponent",
			"autoexpand", "automatically expand pool when devices are replaced",
			"autoreplace", "automatically replace failed devices",
			"autotrim", "automatic TRIM",
			"bootfs", "default bootable dataset",
			"cachefile", "pool cache file location",
			"comment", "pool comment",
			"compatibility", "feature compatibility",
			"dedup_table_quota", "dedup table size quota",
			"dedupditto", "deprecated, no effect",
			"delegation", "allow delegated administration",
			"failmode", "failure mode behavior",
			"listsnapshots", "include snapshots in zfs list",
			"multihost", "multihost protection",
			"readonly", "read-only pool import",
			"version", "on-disk version",
		).Tag("pool properties"),
		ActionPoolFeatures(),
	).ToA()
}

// ActionReadonlyPoolProperties completes read-only ZFS pool properties
//
//	allocated (allocated space)
//	capacity (capacity percentage)
func ActionReadonlyPoolProperties() carapace.Action {
	return carapace.ActionValuesDescribed(
		"allocated", "allocated space",
		"bcloneratio", "block clone ratio",
		"bclonesaved", "block clone space saved",
		"bcloneused", "block clone space used",
		"capacity", "capacity percentage",
		"checkpoint", "checkpoint space",
		"dedup_table_size", "on-disk dedup table size",
		"dedupcached", "dedup table size cached in ARC",
		"dedupratio", "dedup ratio",
		"expandsize", "amount of expandable space",
		"fragmentation", "fragmentation percentage",
		"free", "free space",
		"freeing", "space being freed",
		"guid", "pool GUID",
		"health", "pool health status",
		"last_scrubbed_txg", "txg of last scrub",
		"leaked", "leaked space",
		"load_guid", "pool load GUID",
		"size", "total pool size",
	).Tag("pool properties")
}

// ActionPoolPropertyValues completes values for a given pool property
//
//	on
//	off
func ActionPoolPropertyValues(property string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if strings.HasPrefix(property, "feature@") {
			return carapace.ActionValues("enabled", "disabled").StyleF(style.ForKeyword)
		}
		switch property {
		case "altroot":
			return carapace.ActionDirectories()
		case "ashift":
			return carapace.ActionValues("0", "9", "10", "11", "12", "13", "14", "15", "16")
		case "autoexpand", "autoreplace", "autotrim", "delegation", "listsnapshots", "multihost", "readonly":
			return carapace.ActionValues("on", "off").StyleF(style.ForKeyword)
		case "bootfs":
			return ActionFilesystems()
		case "dedup_table_quota":
			return carapace.ActionValues("none", "auto")
		case "cachefile":
			return carapace.Batch(
				carapace.ActionValues("none"),
				carapace.ActionFiles(),
			).ToA()
		case "compatibility":
			return carapace.Batch(
				carapace.ActionValues("off", "legacy"),
				carapace.ActionFiles(),
			).ToA()
		case "failmode":
			return carapace.ActionValues("wait", "continue", "panic")
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionPoolPropertyAssignments completes pool property=value assignments
//
//	autoexpand=on
//	failmode=continue
func ActionPoolPropertyAssignments() carapace.Action {
	return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionPoolProperties().Suffix("=")
		default:
			return ActionPoolPropertyValues(c.Parts[0])
		}
	})
}

// ActionPermissions completes ZFS delegated permission names
//
//	create (create datasets)
//	destroy (destroy datasets)
func ActionPermissions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"allow", "grant permissions",
		"bookmark", "create bookmarks",
		"change-key", "change encryption key",
		"clone", "clone datasets",
		"create", "create datasets",
		"destroy", "destroy datasets",
		"diff", "show differences",
		"groupobjquota", "access group object quotas",
		"groupobjused", "read group object usage",
		"groupquota", "access group quotas",
		"groupused", "read group usage",
		"hold", "hold snapshots",
		"load-key", "load/unload encryption key",
		"mount", "mount/unmount",
		"promote", "promote clones",
		"projectobjquota", "access project object quotas",
		"projectobjused", "read project object usage",
		"projectquota", "access project quotas",
		"projectused", "read project usage",
		"receive", "receive streams",
		"receive:append", "receive without force",
		"release", "release holds",
		"rename", "rename datasets",
		"rollback", "rollback snapshots",
		"send", "send streams",
		"send:raw", "send raw streams only",
		"share", "share/unshare",
		"snapshot", "create snapshots",
		"userobjquota", "access user object quotas",
		"userobjused", "read user object usage",
		"userprop", "change user properties",
		"userquota", "access user quotas",
		"userused", "read user usage",
	).Tag("permissions")
}
