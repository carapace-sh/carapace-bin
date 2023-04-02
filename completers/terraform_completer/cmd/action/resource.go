package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/terraform"
	"github.com/spf13/cobra"
)

func ActionResources(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if f := cmd.Flag("state"); f != nil {
			return terraform.ActionResources(f.Value.String())
		}
		return terraform.ActionResources("")
	})
}
