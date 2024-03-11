package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_pauseCmd = &cobra.Command{
	Use:   "pause CONTAINER [CONTAINER...]",
	Short: "Pause all processes within one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_pauseCmd).Standalone()

	containerCmd.AddCommand(container_pauseCmd)

	carapace.Gen(container_pauseCmd).PositionalAnyCompletion(docker.ActionContainers())
}
