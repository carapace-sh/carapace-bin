package zfs

import "github.com/carapace-sh/carapace"

// ActionListColumns completes column names for zfs list -o
//
//	name (dataset name)
//	used (space used)
func ActionListColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"name", "dataset name",
		"space", "space usage column set",
		"type", "dataset type",
		"creation", "creation time",
		"used", "space used",
		"avail", "available space",
		"refer", "referenced space",
		"ratio", "compression ratio",
		"compressratio", "compression ratio",
		"mounted", "is mounted",
		"origin", "clone origin",
		"quota", "quota",
		"reservation", "reservation",
		"volsize", "volume size",
		"volblocksize", "volume block size",
		"recordsize", "record size",
		"mountpoint", "mount point",
		"sharenfs", "NFS share",
		"checksum", "checksum algorithm",
		"compression", "compression algorithm",
		"atime", "access time",
		"devices", "devices",
		"exec", "exec",
		"setuid", "setuid",
		"readonly", "read-only",
		"zoned", "zone managed",
		"snapdir", "snapshot directory visibility",
		"acltype", "ACL type",
		"context", "SELinux context",
		"fscontext", "SELinux filesystem context",
		"defcontext", "SELinux default context",
		"rootcontext", "SELinux root context",
		"casesensitivity", "case sensitivity",
		"copies", "copies",
		"dedup", "dedup",
		"logbias", "log bias",
		"primarycache", "primary cache",
		"secondarycache", "secondary cache",
		"usedbysnapshots", "used by snapshots",
		"usedbychildren", "used by children",
		"usedbyrefreservation", "used by refreservation",
		"usedbydataset", "used by dataset",
		"logicalused", "logical used",
		"logicalreferenced", "logical referenced",
		"written", "written since last snapshot",
		"guid", "GUID",
		"encryption", "encryption",
		"keyformat", "key format",
		"keylocation", "key location",
		"keystatus", "key status",
		"encryptionroot", "encryption root",
		"clones", "clones",
		"redact_snaps", "redaction snapshots",
	).Tag("columns")
}

// ActionPoolListColumns completes column names for zpool list -o
//
//	name (pool name)
//	size (total pool size)
func ActionPoolListColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"name", "pool name",
		"size", "total pool size",
		"allocated", "allocated space",
		"free", "free space",
		"checkpoint", "checkpoint space",
		"expandsize", "expandable space",
		"fragmentation", "fragmentation percentage",
		"capacity", "capacity percentage",
		"dedup_table_quota", "dedup table size quota",
		"dedup_table_size", "on-disk dedup table size",
		"dedupcached", "dedup table size cached in ARC",
		"dedupratio", "dedup ratio",
		"health", "health status",
		"altroot", "alternate root",
		"guid", "pool GUID",
		"load_guid", "pool load GUID",
		"ashift", "ashift value",
		"autoexpand", "auto expand",
		"autoreplace", "auto replace",
		"autotrim", "auto trim",
		"bootfs", "boot filesystem",
		"cachefile", "cache file",
		"comment", "comment",
		"compatibility", "compatibility",
		"delegation", "delegation",
		"failmode", "failure mode",
		"listsnapshots", "list snapshots",
		"multihost", "multihost protection",
		"readonly", "read-only",
		"version", "version",
		"freeing", "space being freed",
		"last_scrubbed_txg", "txg of last scrub",
		"leaked", "leaked space",
		"feature@...", "pool feature",
		"bcloneratio", "block clone ratio",
		"bclonesaved", "block clone space saved",
		"bcloneused", "block clone space used",
	).Tag("columns")
}

// ActionDatasetTypes completes ZFS dataset types for list filtering
//
//	filesystem
//	snapshot
func ActionDatasetTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"filesystem", "filesystems",
		"fs", "alias for filesystem",
		"snapshot", "snapshots",
		"snap", "alias for snapshot",
		"volume", "volumes",
		"vol", "alias for volume",
		"bookmark", "bookmarks",
		"all", "all types",
	)
}

// ActionSortProperties completes properties valid for zfs list sort options
//
//	name
//	used
func ActionSortProperties() carapace.Action {
	return carapace.ActionValues(
		"name",
		"type",
		"creation",
		"used",
		"avail",
		"refer",
		"ratio",
		"compressratio",
		"mounted",
		"origin",
		"quota",
		"reservation",
		"volsize",
		"mountpoint",
		"written",
		"logicalused",
		"logicalreferenced",
	).Tag("sort properties")
}
