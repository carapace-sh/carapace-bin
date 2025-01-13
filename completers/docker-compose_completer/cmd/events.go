package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var eventsCmd = &cobra.Command{
	Use:   "events [OPTIONS] [SERVICE...]",
	Short: "Receive real time events from containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eventsCmd).Standalone()

	eventsCmd.Flags().Bool("json", false, "Output events as a stream of json objects")
	rootCmd.AddCommand(eventsCmd)

	carapace.Gen(eventsCmd).PositionalAnyCompletion(
		action.ActionServices(eventsCmd).FilterArgs(),
	)
}
