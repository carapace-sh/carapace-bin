package zpool

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPools completes ZFS storage pools.
func ActionPools() carapace.Action {
	return carapace.ActionExecCommand("zpool", "list", "-H", "-o", "name")(func(output []byte) carapace.Action {
		lines := strings.Split(strings.TrimSpace(string(output)), "\n")

		values := make([]string, 0, len(lines))
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 0 {
				values = append(values, fields[0])
			}
		}
		return carapace.ActionValues(values...).Tag("zfs pools")
	})
}

// ActionVdevTypes completes zpool virtual device type keywords.
func ActionVdevTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"mirror", "mirrored vdev",
		"raidz", "RAID-Z vdev",
		"raidz1", "single-parity RAID-Z vdev",
		"raidz2", "double-parity RAID-Z vdev",
		"raidz3", "triple-parity RAID-Z vdev",
		"draid", "distributed RAID-Z vdev",
		"spare", "hot spare vdev",
		"log", "separate intent log vdev",
		"cache", "L2ARC cache vdev",
		"special", "special allocation class vdev",
		"dedup", "deduplication table vdev",
	).Tag("zpool vdev types")
}
