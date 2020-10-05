package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/actions/os"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var container_killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kill one or more running containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_killCmd).Standalone()

	container_killCmd.Flags().StringP("signal", "s", "", "Signal to send to the container (default \"KILL\")")
	containerCmd.AddCommand(container_killCmd)

	carapace.Gen(container_killCmd).FlagCompletion(carapace.ActionMap{
		"signal": os.ActionKillSignals(),
	})

	carapace.Gen(container_killCmd).PositionalAnyCompletion(action.ActionContainers())
}
