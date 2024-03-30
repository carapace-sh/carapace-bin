package transmission

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// ActionFilter completes filters
func ActionFilter() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.Batch(
			actionFilterConditions(),
			actionFilterConditions().Prefix("~"),
			carapace.ActionValuesDescribed("~", "Negate the following filter"),
		).ToA()
	})
}

// Available filters for torrents
func actionFilterConditions() carapace.Action {
	return carapace.Batch(
		carapace.ActionValuesDescribed(
			"i", "currently idle",
			"u", "currently uploading",
			"d", "currently downloading",
			"w", "torrent has unwanted files",
		),
		carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			if len(c.Parts) == 0 {
				return carapace.ActionValuesDescribed(
					"n:", "torrent name includes the string following :",
					"l:", "torrent has the label following :",
					"r:", "torrent has an upload ratio greater than or equal to the ratio following :",
				)
			}
			return carapace.ActionValues()
		}),
	).ToA()
}

func getFilters(cmd *cobra.Command) (result []string) {
	result, _ = cmd.Root().Flags().GetStringArray("filter")
	return
}

// ActionIds completes torrent IDs, accounting for previously applied filters
func ActionIds(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		filters := getFilters(cmd)
		return carapace.Batch(
			actionIdsValues(),
			getIdsDescribed(filters),
		).ToA()
	})
}

func actionIdsValues() carapace.Action {
	return carapace.ActionValuesDescribed(
		"all", "select all filtered torrents",
		"active", "select active filtered torrents",
	)
}

// Call into transmission-remote to the list of ids resulting from the filters
func getIdsDescribed(rawFilters []string) carapace.Action {
	filters := make([]string, len(rawFilters)*2+1)
	for idx, f := range rawFilters {
		filters[idx*2] = "-F"
		filters[idx*2+1] = f
	}
	filters[len(filters)-1] = "-l"

	// Can't use --print-ids flag to get ids since it doesn't give descriptions and doesn't always return numeric ids
	return carapace.ActionExecCommand("transmission-remote", filters...)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		idsDescribed := make([]string, 0)
		// The first and last lines are always additional information, not torrents
		for _, line := range lines[1 : len(lines)-2] {
			fields := strings.Fields(line)
			if len(fields) <= 9 {
				idsDescribed = append(idsDescribed, fields[0], "")
			} else {
				idsDescribed = append(idsDescribed, fields[0], strings.Join(fields[9:], " "))
			}
		}

		return carapace.ActionValuesDescribed(idsDescribed...)
	})
}
