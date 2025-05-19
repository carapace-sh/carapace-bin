package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:   "pause [SERVICE...]",
	Short: "Pause services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pauseCmd).Standalone()

	rootCmd.AddCommand(pauseCmd)

	carapace.Gen(pauseCmd).PositionalAnyCompletion(
		action.ActionServices(pauseCmd).FilterArgs(),
	)
}
