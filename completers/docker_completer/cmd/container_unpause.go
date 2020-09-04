package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var container_unpauseCmd = &cobra.Command{
	Use:   "unpause",
	Short: "Unpause all processes within one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_unpauseCmd).Standalone()

	containerCmd.AddCommand(container_unpauseCmd)

	carapace.Gen(container_unpauseCmd).PositionalAnyCompletion(action.ActionContainers())
}
