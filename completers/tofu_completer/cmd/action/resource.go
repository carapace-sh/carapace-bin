package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tofu"
	"github.com/spf13/cobra"
)

func ActionResources(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if f := cmd.Flag("state"); f != nil {
			return tofu.ActionResources(f.Value.String())
		}
		return tofu.ActionResources("")
	})
}
