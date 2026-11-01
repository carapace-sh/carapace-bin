package zfs

import "github.com/carapace-sh/carapace"

// ActionListColumns completes column names for zfs list -o
//
//	name (dataset name)
//	used (space used)
func ActionListColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"acltype", "ACL type",
		"atime", "access time",
		"avail", "available space",
		"casesensitivity", "case sensitivity",
		"checksum", "checksum algorithm",
		"clones", "clones",
		"compression", "compression algorithm",
		"compressratio", "compression ratio",
		"context", "SELinux context",
		"copies", "copies",
		"creation", "creation time",
		"dedup", "dedup",
		"defcontext", "SELinux default context",
		"devices", "devices",
		"encryption", "encryption",
		"encryptionroot", "encryption root",
		"exec", "exec",
		"fscontext", "SELinux filesystem context",
		"guid", "GUID",
		"keyformat", "key format",
		"keylocation", "key location",
		"keystatus", "key status",
		"logbias", "log bias",
		"logicalreferenced", "logical referenced",
		"logicalused", "logical used",
		"mounted", "is mounted",
		"mountpoint", "mount point",
		"name", "dataset name",
		"origin", "clone origin",
		"primarycache", "primary cache",
		"quota", "quota",
		"ratio", "compression ratio",
		"readonly", "read-only",
		"recordsize", "record size",
		"redact_snaps", "redaction snapshots",
		"refer", "referenced space",
		"reservation", "reservation",
		"rootcontext", "SELinux root context",
		"secondarycache", "secondary cache",
		"setuid", "setuid",
		"sharenfs", "NFS share",
		"snapdir", "snapshot directory visibility",
		"space", "space usage column set",
		"type", "dataset type",
		"used", "space used",
		"usedbychildren", "used by children",
		"usedbydataset", "used by dataset",
		"usedbyrefreservation", "used by refreservation",
		"usedbysnapshots", "used by snapshots",
		"volblocksize", "volume block size",
		"volsize", "volume size",
		"written", "written since last snapshot",
		"zoned", "zone managed",
	).Tag("columns")
}

// ActionPoolListColumns completes column names for zpool list -o
//
//	name (pool name)
//	size (total pool size)
func ActionPoolListColumns() carapace.Action {
	return carapace.Batch(
		carapace.ActionValuesDescribed(
			"allocated", "allocated space",
			"altroot", "alternate root",
			"ashift", "ashift value",
			"autoexpand", "auto expand",
			"autoreplace", "auto replace",
			"autotrim", "auto trim",
			"bcloneratio", "block clone ratio",
			"bclonesaved", "block clone space saved",
			"bcloneused", "block clone space used",
			"bootfs", "boot filesystem",
			"cachefile", "cache file",
			"capacity", "capacity percentage",
			"checkpoint", "checkpoint space",
			"comment", "comment",
			"compatibility", "compatibility",
			"dedup_table_quota", "dedup table size quota",
			"dedup_table_size", "on-disk dedup table size",
			"dedupcached", "dedup table size cached in ARC",
			"dedupratio", "dedup ratio",
			"delegation", "delegation",
			"expandsize", "expandable space",
			"failmode", "failure mode",
			"fragmentation", "fragmentation percentage",
			"free", "free space",
			"freeing", "space being freed",
			"guid", "pool GUID",
			"health", "health status",
			"last_scrubbed_txg", "txg of last scrub",
			"leaked", "leaked space",
			"listsnapshots", "list snapshots",
			"load_guid", "pool load GUID",
			"multihost", "multihost protection",
			"name", "pool name",
			"readonly", "read-only",
			"size", "total pool size",
			"version", "version",
		).Tag("columns"),
		ActionPoolFeatures(),
	).ToA()
}

// ActionDatasetTypes completes ZFS dataset types for list filtering
//
//	filesystem
//	snapshot
func ActionDatasetTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"all", "all types",
		"bookmark", "bookmarks",
		"filesystem", "filesystems",
		"fs", "alias for filesystem",
		"snap", "alias for snapshot",
		"snapshot", "snapshots",
		"vol", "alias for volume",
		"volume", "volumes",
	)
}

// ActionSortProperties completes properties valid for zfs list sort options
//
//	name
//	used
func ActionSortProperties() carapace.Action {
	return carapace.ActionValues(
		"avail",
		"compressratio",
		"creation",
		"logicalreferenced",
		"logicalused",
		"mounted",
		"mountpoint",
		"name",
		"origin",
		"quota",
		"ratio",
		"refer",
		"reservation",
		"type",
		"used",
		"volsize",
		"written",
	).Tag("sort properties")
}
