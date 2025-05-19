package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var unpauseCmd = &cobra.Command{
	Use:   "unpause [SERVICE...]",
	Short: "Unpause services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unpauseCmd).Standalone()

	rootCmd.AddCommand(unpauseCmd)

	carapace.Gen(unpauseCmd).PositionalAnyCompletion(
		action.ActionServices(unpauseCmd).FilterArgs(),
	)
}
