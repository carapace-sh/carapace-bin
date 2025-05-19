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

	rootCmd.AddCommand(startCmd)

	carapace.Gen(startCmd).PositionalAnyCompletion(
		action.ActionServices(startCmd).FilterArgs(),
	)
}
