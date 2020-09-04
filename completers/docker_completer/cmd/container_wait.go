package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var container_waitCmd = &cobra.Command{
	Use:   "wait",
	Short: "Block until one or more containers stop, then print their exit codes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_waitCmd).Standalone()

	containerCmd.AddCommand(container_waitCmd)

	carapace.Gen(container_waitCmd).PositionalAnyCompletion(action.ActionContainers())
}
