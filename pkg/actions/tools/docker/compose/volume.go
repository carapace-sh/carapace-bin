package compose

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
)

// ActionVolumes completes volumes
//
//	go (carapace-bin_go)
//	example (carapace-bin_example)
func ActionVolumes(files ...string) carapace.Action {
	return actionConfig(files, func(c config) carapace.Action {
		vals := make([]string, 0)
		for name, volume := range c.Volumes {
			vals = append(vals, name, volume.Name)
		}
		return carapace.ActionValuesDescribed(vals...).StyleR(&styles.Docker.Volume)
	})
}
