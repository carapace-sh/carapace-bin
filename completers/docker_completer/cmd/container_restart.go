package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_restartCmd = &cobra.Command{
	Use:   "restart [OPTIONS] CONTAINER [CONTAINER...]",
	Short: "Restart one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_restartCmd).Standalone()

	container_restartCmd.Flags().StringP("signal", "s", "", "Signal to send to the container")
	container_restartCmd.Flags().IntP("time", "t", 0, "Seconds to wait before killing the container")
	containerCmd.AddCommand(container_restartCmd)

	carapace.Gen(container_restartCmd).FlagCompletion(carapace.ActionMap{
		"signal": ps.ActionKillSignals(),
	})

	carapace.Gen(container_restartCmd).PositionalAnyCompletion(docker.ActionContainers())
}
