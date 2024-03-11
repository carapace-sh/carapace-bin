package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_stopCmd = &cobra.Command{
	Use:   "stop [OPTIONS] CONTAINER [CONTAINER...]",
	Short: "Stop one or more running containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_stopCmd).Standalone()

	container_stopCmd.Flags().StringP("signal", "s", "", "Signal to send to the container")
	container_stopCmd.Flags().IntP("time", "t", 0, "Seconds to wait before killing the container")
	containerCmd.AddCommand(container_stopCmd)

	carapace.Gen(container_stopCmd).PositionalAnyCompletion(docker.ActionContainers())
}
