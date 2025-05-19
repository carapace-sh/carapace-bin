package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_diffCmd = &cobra.Command{
	Use:   "diff CONTAINER",
	Short: "Inspect changes to files or directories on a container's filesystem",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_diffCmd).Standalone()

	containerCmd.AddCommand(container_diffCmd)

	carapace.Gen(container_diffCmd).PositionalCompletion(docker.ActionContainers())
}
