package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_topCmd = &cobra.Command{
	Use:   "top CONTAINER [ps OPTIONS]",
	Short: "Display the running processes of a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_topCmd).Standalone()
	container_topCmd.Flags().SetInterspersed(false)

	containerCmd.AddCommand(container_topCmd)

	carapace.Gen(container_topCmd).PositionalAnyCompletion(docker.ActionContainers())
}
