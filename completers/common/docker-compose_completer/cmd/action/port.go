package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker/compose"
	"github.com/spf13/cobra"
)

func ActionPorts(cmd *cobra.Command, service string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return compose.ActionPorts(compose.ServicePathOpts{Files: files(cmd), Service: service})
	})
}
