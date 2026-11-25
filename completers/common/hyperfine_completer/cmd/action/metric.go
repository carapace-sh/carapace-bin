package action

import (
	"slices"
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionMetrics() carapace.Action {
	return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
		selected := make([]string, 0, len(c.Parts))
		for _, part := range c.Parts {
			id, _, _ := strings.Cut(part, ":")
			switch id {
			case "default", "time", "all":
				return carapace.ActionValues()
			}
			selected = append(selected, id)
		}
		standalone := len(c.Parts) == 0

		return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			if len(c.Parts) == 0 {
				metrics := carapace.ActionValuesDescribed(
					"branch_misses", "Branch misses",
					"cache_misses", "Cache misses",
					"cache_references", "Cache references",
					"cpu_cycles", "CPU cycles",
					"instructions", "Instructions executed",
					"memory_peak_resident", "Peak resident memory",
					"time_cpu", "CPU time",
					"time_system", "System CPU time",
					"time_user", "User CPU time",
					"time_wall_clock", "Wall-clock time",
				)
				if !standalone {
					return metrics.Filter(selected...)
				}
				return carapace.Batch(metrics, carapace.ActionValuesDescribed(
					"all", "Preset: all available metrics",
					"default", "Preset: wall-clock time and peak memory",
					"time", "Preset: wall-clock, CPU, user, and system time",
				)).ToA()
			}

			if slices.Contains(selected, c.Parts[0]) {
				return carapace.ActionValues()
			}
			switch c.Parts[0] {
			case "time_wall_clock", "time_cpu", "time_user", "time_system":
				return carapace.ActionValues( // carapace:curated
					"ns", "us", "ms", "s", "min", "h",
				)
			case "memory_peak_resident":
				return carapace.ActionValues( // carapace:curated
					"B", "kB", "MB", "GB", "TB",
					"KiB", "MiB", "GiB", "TiB",
				)
			case "cpu_cycles", "instructions", "cache_references", "cache_misses", "branch_misses":
				return carapace.ActionValues( // carapace:curated
					"count", "k", "M", "B",
				)
			default:
				return carapace.ActionValues()
			}
		})
	})
}
