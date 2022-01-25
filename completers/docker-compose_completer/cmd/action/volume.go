package action

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func ActionVolumes(cmd *cobra.Command) carapace.Action {
	return actionConfig(func(c config) carapace.Action {
		vals := make([]string, 0)
		for name, volume := range c.Volumes {
			vals = append(vals, name, volume.Name)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
