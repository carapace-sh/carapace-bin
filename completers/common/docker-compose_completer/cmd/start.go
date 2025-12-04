package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [SERVICE...]",
	Short: "Start services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startCmd).Standalone()

	startCmd.Flags().Bool("wait", false, "Wait for services to be running|healthy. Implies detached mode.")
	startCmd.Flags().String("wait-timeout", "", "Maximum duration in seconds to wait for the project to be running|healthy")
	rootCmd.AddCommand(startCmd)

	carapace.Gen(startCmd).PositionalAnyCompletion(
		action.ActionServices(startCmd).FilterArgs(),
	)
}
