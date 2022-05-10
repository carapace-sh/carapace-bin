package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker/compose"
	"github.com/spf13/cobra"
)

func ActionContainers(cmd *cobra.Command, opts compose.ContainerOpts) carapace.Action {
	files, err := cmd.Root().Flags().GetStringArray("file")
	if err != nil {
		return carapace.ActionMessage(err.Error())
	}
	opts.Files = files
	return compose.ActionContainers(opts)
}
