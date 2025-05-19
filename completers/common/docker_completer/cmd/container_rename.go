package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_renameCmd = &cobra.Command{
	Use:   "rename CONTAINER NEW_NAME",
	Short: "Rename a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_renameCmd).Standalone()

	containerCmd.AddCommand(container_renameCmd)

	carapace.Gen(container_renameCmd).PositionalCompletion(
		docker.ActionContainers(),
	)
}
