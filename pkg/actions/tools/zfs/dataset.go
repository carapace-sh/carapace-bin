package zfs

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionDatasets completes ZFS datasets for the given dataset types.
func ActionDatasets(types ...string) carapace.Action {
	if len(types) == 0 {
		types = []string{"all"}
	}

	return carapace.ActionExecCommand("zfs", "list", "-H", "-o", "name", "-t", strings.Join(types, ","))(func(output []byte) carapace.Action {
		return actionValuesFromLines(output).Tag("zfs datasets")
	})
}

// ActionSnapshots completes ZFS snapshots.
func ActionSnapshots() carapace.Action {
	return ActionDatasets("snapshot").Tag("zfs snapshots")
}

// ActionBookmarks completes ZFS bookmarks.
func ActionBookmarks() carapace.Action {
	return ActionDatasets("bookmark").Tag("zfs bookmarks")
}

// ActionDatasetTypes completes ZFS dataset types.
func ActionDatasetTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"filesystem", "file system datasets",
		"snapshot", "read-only point-in-time datasets",
		"volume", "block device datasets",
		"bookmark", "snapshot bookmarks",
		"all", "all dataset types",
	).Tag("zfs dataset types")
}

func actionValuesFromLines(output []byte) carapace.Action {
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")

	values := make([]string, 0, len(lines))
	for _, line := range lines {
		if line != "" {
			values = append(values, line)
		}
	}
	return carapace.ActionValues(values...)
}
