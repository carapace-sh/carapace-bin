package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/devenv"
	"github.com/spf13/cobra"
)

var container_buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_buildCmd).Standalone()

	containerCmd.AddCommand(container_buildCmd)

	carapace.Gen(container_buildCmd).PositionalCompletion(
		devenv.ActionContainers(),
	)
}
