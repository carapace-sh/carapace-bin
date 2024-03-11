package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_unpauseCmd = &cobra.Command{
	Use:   "unpause CONTAINER [CONTAINER...]",
	Short: "Unpause all processes within one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_unpauseCmd).Standalone()

	containerCmd.AddCommand(container_unpauseCmd)

	carapace.Gen(container_unpauseCmd).PositionalAnyCompletion(docker.ActionContainers())
}
