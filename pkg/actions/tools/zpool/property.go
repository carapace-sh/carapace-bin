package zpool

import "github.com/carapace-sh/carapace"

// ActionPoolProperties completes common zpool properties.
func ActionPoolProperties() carapace.Action {
	return carapace.ActionValuesDescribed(
		"all", "all properties",
		"allocated", "allocated space",
		"altroot", "alternate root directory",
		"ashift", "sector size exponent",
		"autoexpand", "automatic pool expansion",
		"autoreplace", "automatic device replacement",
		"autotrim", "automatic TRIM behavior",
		"bootfs", "default bootable dataset",
		"cachefile", "pool configuration cache file",
		"capacity", "used capacity percentage",
		"checkpoint", "checkpoint space",
		"comment", "pool comment",
		"compatibility", "feature compatibility set",
		"dedupratio", "deduplication ratio",
		"delegation", "delegated administration setting",
		"expandsize", "expandable space",
		"failmode", "failure mode",
		"fragmentation", "fragmentation percentage",
		"free", "free space",
		"freeing", "space being freed",
		"guid", "pool GUID",
		"health", "pool health",
		"leaked", "leaked space",
		"listsnapshots", "snapshot listing behavior",
		"multihost", "multihost import protection",
		"name", "pool name",
		"readonly", "read-only import state",
		"size", "total pool size",
		"version", "on-disk version",
	).Tag("zpool properties")
}

// ActionListProperties completes zpool list output properties.
func ActionListProperties() carapace.Action {
	return carapace.ActionValuesDescribed(
		"name", "pool name",
		"size", "total pool size",
		"allocated", "allocated space",
		"free", "free space",
		"checkpoint", "checkpoint space",
		"expandsize", "expandable space",
		"fragmentation", "fragmentation percentage",
		"capacity", "used capacity percentage",
		"dedupratio", "deduplication ratio",
		"health", "pool health",
		"altroot", "alternate root directory",
	).Tag("zpool list properties")
}

// ActionGetFields completes zpool get output fields.
func ActionGetFields() carapace.Action {
	return carapace.ActionValues("name", "property", "value", "source").Tag("zpool get fields")
}

// ActionTimestampStyles completes zpool timestamp styles.
func ActionTimestampStyles() carapace.Action {
	return carapace.ActionValuesDescribed(
		"u", "Unix epoch timestamp",
		"d", "standard date timestamp",
	).Tag("zpool timestamp styles")
}

// ActionWaitActivities completes zpool wait activity types.
func ActionWaitActivities() carapace.Action {
	return carapace.ActionValues("discard", "free", "initialize", "replace", "remove", "resilver", "scrub", "trim").Tag("zpool wait activities")
}
