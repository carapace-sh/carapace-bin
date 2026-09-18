package docker

import (
	"fmt"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace"
)

func actionLabels(args ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", args...)(func(output []byte) carapace.Action {
			values := map[string]map[string]bool{} // key -> distinct values
			for _, line := range strings.Split(string(output), "\n") {
				for _, label := range strings.Split(line, ",") {
					if key, value, found := strings.Cut(label, "="); found && key != "" {
						if values[key] == nil {
							values[key] = map[string]bool{}
						}
						values[key][value] = true
					}
				}
			}

			keys := make([]string, 0, len(values))
			for key := range values {
				keys = append(keys, key)
			}
			sort.Strings(keys)

			vals := make([]string, 0, len(keys)*2)
			for _, key := range keys {
				if len(values[key]) == 1 {
					for value := range values[key] {
						vals = append(vals, key, value)
					}
				} else {
					vals = append(vals, key, fmt.Sprintf("%v values", len(values[key])))
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		}).Suppress("This node is not a swarm manager")
	})
}
