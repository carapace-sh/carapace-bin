package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill [OPTIONS] [SERVICE...]",
	Short: "Force stop service containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killCmd).Standalone()

	killCmd.Flags().Bool("remove-orphans", false, "Remove containers for services not defined in the Compose file")
	killCmd.Flags().StringP("signal", "s", "", "SIGNAL to send to the container")
	rootCmd.AddCommand(killCmd)

	carapace.Gen(killCmd).FlagCompletion(carapace.ActionMap{
		"signal": ps.ActionKillSignals(),
	})

	carapace.Gen(killCmd).PositionalAnyCompletion(
		action.ActionServices(killCmd).FilterArgs(),
	)
}
