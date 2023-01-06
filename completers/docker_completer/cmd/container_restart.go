package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
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

	rootAlias(container_restartCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalAnyCompletion(docker.ActionContainers())
	})
}
