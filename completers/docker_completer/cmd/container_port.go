package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var container_portCmd = &cobra.Command{
	Use:   "port",
	Short: "List port mappings or a specific mapping for the container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_portCmd).Standalone()

	containerCmd.AddCommand(container_portCmd)

	carapace.Gen(container_portCmd).PositionalCompletion(
		action.ActionContainers(),
		// TODO completion port
	)
}
