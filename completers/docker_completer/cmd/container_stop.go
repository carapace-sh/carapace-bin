package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var container_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop one or more running containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_stopCmd).Standalone()

	container_stopCmd.Flags().StringP("time", "t", "", "Seconds to wait for stop before killing it (default 10)")
	containerCmd.AddCommand(container_stopCmd)

	carapace.Gen(container_stopCmd).PositionalAnyCompletion(action.ActionContainers())
}
