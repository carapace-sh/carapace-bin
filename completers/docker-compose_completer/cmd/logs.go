package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs [OPTIONS] [SERVICE...]",
	Short: "View output from containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logsCmd).Standalone()

	logsCmd.Flags().BoolP("follow", "f", false, "Follow log output")
	logsCmd.Flags().String("index", "", "index of the container if service has multiple replicas")
	logsCmd.Flags().Bool("no-color", false, "Produce monochrome output")
	logsCmd.Flags().Bool("no-log-prefix", false, "Don't print prefix in logs")
	logsCmd.Flags().String("since", "", "Show logs since timestamp (e.g. 2013-01-02T13:23:37Z) or relative (e.g. 42m for 42 minutes)")
	logsCmd.Flags().StringP("tail", "n", "", "Number of lines to show from the end of the logs for each container")
	logsCmd.Flags().BoolP("timestamps", "t", false, "Show timestamps")
	logsCmd.Flags().String("until", "", "Show logs before a timestamp (e.g. 2013-01-02T13:23:37Z) or relative (e.g. 42m for 42 minutes)")
	rootCmd.AddCommand(logsCmd)

	// TODO index

	carapace.Gen(logsCmd).PositionalAnyCompletion(
		action.ActionServices(logsCmd).FilterArgs(),
	)
}
