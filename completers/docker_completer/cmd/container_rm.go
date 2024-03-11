package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_rmCmd = &cobra.Command{
	Use:     "rm [OPTIONS] CONTAINER [CONTAINER...]",
	Short:   "Remove one or more containers",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_rmCmd).Standalone()

	container_rmCmd.Flags().BoolP("force", "f", false, "Force the removal of a running container (uses SIGKILL)")
	container_rmCmd.Flags().BoolP("link", "l", false, "Remove the specified link")
	container_rmCmd.Flags().BoolP("volumes", "v", false, "Remove anonymous volumes associated with the container")
	containerCmd.AddCommand(container_rmCmd)

	carapace.Gen(container_rmCmd).PositionalAnyCompletion(docker.ActionContainers())
}
