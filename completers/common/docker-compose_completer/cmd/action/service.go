package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker/compose"
	"github.com/spf13/cobra"
)

func ActionServices(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return compose.ActionServices(files(cmd)...)
	})
}
