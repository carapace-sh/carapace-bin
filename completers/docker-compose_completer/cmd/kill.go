package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Force stop service containers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killCmd).Standalone()
	killCmd.Flags().StringP("signal", "s", "SIGKILL", "SIGNAL to send to the container.")
	rootCmd.AddCommand(killCmd)

	carapace.Gen(killCmd).FlagCompletion(carapace.ActionMap{
		"signal": os.ActionKillSignals(),
	})

	carapace.Gen(killCmd).PositionalAnyCompletion(
		action.ActionServices(killCmd),
	)
}
