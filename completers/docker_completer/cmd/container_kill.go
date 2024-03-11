package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_killCmd = &cobra.Command{
	Use:   "kill [OPTIONS] CONTAINER [CONTAINER...]",
	Short: "Kill one or more running containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_killCmd).Standalone()

	container_killCmd.Flags().StringP("signal", "s", "", "Signal to send to the container")
	containerCmd.AddCommand(container_killCmd)

	carapace.Gen(container_killCmd).FlagCompletion(carapace.ActionMap{
		"signal": ps.ActionKillSignals(),
	})

	carapace.Gen(container_killCmd).PositionalAnyCompletion(docker.ActionContainers())
}
