package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_waitCmd = &cobra.Command{
	Use:   "wait CONTAINER [CONTAINER...]",
	Short: "Block until one or more containers stop, then print their exit codes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_waitCmd).Standalone()

	containerCmd.AddCommand(container_waitCmd)

	carapace.Gen(container_waitCmd).PositionalAnyCompletion(docker.ActionContainers())
}
