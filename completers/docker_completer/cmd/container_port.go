package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_portCmd = &cobra.Command{
	Use:   "port CONTAINER [PRIVATE_PORT[/PROTO]]",
	Short: "List port mappings or a specific mapping for the container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_portCmd).Standalone()

	containerCmd.AddCommand(container_portCmd)

	carapace.Gen(container_portCmd).PositionalCompletion(
		docker.ActionContainers(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return docker.ActionContainerPorts(c.Args[0])
		}),
	)
}
