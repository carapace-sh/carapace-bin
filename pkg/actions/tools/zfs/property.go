package zfs

import "github.com/carapace-sh/carapace"

// ActionDatasetProperties completes common ZFS dataset properties.
func ActionDatasetProperties() carapace.Action {
	return carapace.ActionValuesDescribed(
		"all", "all properties",
		"available", "space available",
		"clones", "clone datasets",
		"compressratio", "compression ratio",
		"compression", "compression algorithm",
		"copies", "number of data copies",
		"creation", "creation time",
		"dedup", "deduplication setting",
		"filesystem_count", "descendant filesystem count",
		"filesystem_limit", "descendant filesystem limit",
		"guid", "dataset GUID",
		"logicalreferenced", "logical space referenced",
		"logicalused", "logical space used",
		"mountpoint", "mount point",
		"mounted", "mounted state",
		"name", "dataset name",
		"origin", "clone origin",
		"quota", "space quota",
		"readonly", "read-only flag",
		"recordsize", "record size",
		"referenced", "referenced space",
		"refquota", "referenced quota",
		"refreservation", "referenced reservation",
		"reservation", "minimum reserved space",
		"setuid", "setuid behavior",
		"snapdir", "snapshot directory visibility",
		"snapshot_count", "snapshot count",
		"snapshot_limit", "snapshot limit",
		"type", "dataset type",
		"used", "used space",
		"usedbychildren", "space used by children",
		"usedbydataset", "space used by dataset",
		"usedbyrefreservation", "space used by refreservation",
		"usedbysnapshots", "space used by snapshots",
		"volblocksize", "volume block size",
		"volsize", "volume size",
		"written", "space written",
		"xattr", "extended attribute behavior",
		"zoned", "zone visibility",
	).Tag("zfs dataset properties")
}

// ActionGetFields completes zfs get output fields.
func ActionGetFields() carapace.Action {
	return carapace.ActionValues("name", "property", "value", "source", "received").Tag("zfs get fields")
}

// ActionPropertySources completes zfs get property source filters.
func ActionPropertySources() carapace.Action {
	return carapace.ActionValuesDescribed(
		"local", "locally set properties",
		"default", "default property values",
		"inherited", "inherited property values",
		"temporary", "temporary property values",
		"received", "received property values",
		"none", "properties with no source",
	).Tag("zfs property sources")
}

// ActionUserSpaceTypes completes zfs userspace and groupspace type filters.
func ActionUserSpaceTypes() carapace.Action {
	return carapace.ActionValues(
		"all",
		"posixuser",
		"smbuser",
		"posixgroup",
		"smbgroup",
		"project",
	).Tag("zfs userspace types")
}

// ActionUserSpaceFields completes zfs userspace and groupspace output fields.
func ActionUserSpaceFields() carapace.Action {
	return carapace.ActionValues(
		"type",
		"name",
		"used",
		"quota",
		"objused",
		"objquota",
	).Tag("zfs userspace fields")
}

// ActionWaitActivities completes zfs wait activity types.
func ActionWaitActivities() carapace.Action {
	return carapace.ActionValues("deleteq", "freeing", "initializing", "resilver", "scrub", "trim").Tag("zfs wait activities")
}

// ActionPermissions completes zfs delegated permissions.
func ActionPermissions() carapace.Action {
	return carapace.ActionValues(
		"allow",
		"bookmark",
		"change-key",
		"clone",
		"create",
		"destroy",
		"diff",
		"groupquota",
		"groupused",
		"hold",
		"load-key",
		"mount",
		"promote",
		"projectquota",
		"projectused",
		"receive",
		"release",
		"rename",
		"rollback",
		"send",
		"share",
		"snapshot",
		"userprop",
		"userquota",
		"userused",
	).Tag("zfs permissions")
}
