package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rex"
	"github.com/spf13/cobra"
)

// ActionBlockFlavors completes block flavors
func ActionBlockFlavors(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return rex.ActionBlockFlavors(flagValue(cmd, "server"))
	})
}
