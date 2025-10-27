package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
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
	container_stopCmd.Flags().String("time", "", "Seconds to wait before killing the container (deprecated: use --timeout)")
	container_stopCmd.Flags().StringP("timeout", "t", "", "Seconds to wait before killing the container")
	container_stopCmd.Flag("time").Hidden = true
	containerCmd.AddCommand(container_stopCmd)

	carapace.Gen(container_stopCmd).FlagCompletion(carapace.ActionMap{
		"signal": ps.ActionKillSignals(),
	})

	carapace.Gen(container_stopCmd).PositionalAnyCompletion(
		docker.ActionContainers(),
	)
}
